package domain

import (
	"fmt"
	"regexp"
	"strings"
)

// CriterionKind é o tipo de efeito colateral observado na categoria
// Compilador/Interpretador.
type CriterionKind string

const (
	// KindFileExists confere que um artefato foi gerado. Ex: o binário existe.
	KindFileExists CriterionKind = "file_exists"

	// KindExitCode roda um comando e confere o código de saída.
	KindExitCode CriterionKind = "exit_code"

	// KindStdoutEquals roda um comando e compara a saída exatamente.
	KindStdoutEquals CriterionKind = "stdout_equals"

	// KindStdoutMatches roda um comando e casa a saída com uma regex.
	KindStdoutMatches CriterionKind = "stdout_matches"
)

// Criterion é uma condição verificável sobre o resultado do trabalho manual
// do usuário.
type Criterion struct {
	Kind     CriterionKind
	Describe string // texto amigável exibido na TUI; opcional

	Path     string   // KindFileExists
	Command  []string // demais kinds
	Code     int      // KindExitCode
	Expected string   // KindStdoutEquals
	Pattern  string   // KindStdoutMatches
}

// ProbeKind é o tipo de medição que o engine precisa executar.
type ProbeKind string

const (
	ProbeStat ProbeKind = "stat" // existe este caminho?
	ProbeRun  ProbeKind = "run"  // rode este comando e me diga o que aconteceu
)

// Probe descreve *o que medir*, sem medir.
//
// Esta cisão é o que mantém o domínio puro: avaliar um critério exigiria rodar
// comandos dentro de um container, o que é I/O. Em vez disso, o critério
// descreve a medição (Probe), o engine a executa, e o critério interpreta o
// resultado (Evaluate). Assim toda a lógica de decisão é testável sem Docker.
type Probe struct {
	Kind    ProbeKind
	Path    string
	Command []string
}

// ProbeResult é o que o engine devolve depois de executar uma Probe.
type ProbeResult struct {
	Exists   bool   // ProbeStat
	ExitCode int    // ProbeRun
	Stdout   string // ProbeRun
	Stderr   string // ProbeRun

	// Err sinaliza falha de infraestrutura (container morreu, timeout), que é
	// diferente de "o critério não foi satisfeito".
	Err error
}

// Outcome é o veredito sobre um critério, com o motivo para exibição.
type Outcome struct {
	Passed bool
	Detail string
}

// Probe descreve a medição necessária para avaliar este critério.
func (c Criterion) Probe() Probe {
	if c.Kind == KindFileExists {
		return Probe{Kind: ProbeStat, Path: c.Path}
	}
	return Probe{Kind: ProbeRun, Command: c.Command}
}

// Evaluate interpreta o resultado da medição. É função pura.
func (c Criterion) Evaluate(r ProbeResult) Outcome {
	if r.Err != nil {
		return Outcome{Passed: false, Detail: fmt.Sprintf("não foi possível verificar: %v", r.Err)}
	}

	switch c.Kind {
	case KindFileExists:
		if r.Exists {
			return Outcome{Passed: true, Detail: fmt.Sprintf("%s existe", c.Path)}
		}
		return Outcome{Passed: false, Detail: fmt.Sprintf("%s não foi encontrado", c.Path)}

	case KindExitCode:
		if r.ExitCode == c.Code {
			return Outcome{Passed: true, Detail: fmt.Sprintf("exit %d, como esperado", c.Code)}
		}
		return Outcome{Passed: false, Detail: fmt.Sprintf("esperava exit %d, veio %d", c.Code, r.ExitCode)}

	case KindStdoutEquals:
		if r.Stdout == c.Expected {
			return Outcome{Passed: true, Detail: "saída bate exatamente"}
		}
		return Outcome{
			Passed: false,
			Detail: fmt.Sprintf("esperava %q, veio %q", c.Expected, truncate(r.Stdout, 200)),
		}

	case KindStdoutMatches:
		re, err := regexp.Compile(c.Pattern)
		if err != nil {
			// Não deveria acontecer: Validate() compila no parsing do manifesto.
			return Outcome{Passed: false, Detail: fmt.Sprintf("pattern inválido %q: %v", c.Pattern, err)}
		}
		if re.MatchString(r.Stdout) {
			return Outcome{Passed: true, Detail: fmt.Sprintf("saída casa com /%s/", c.Pattern)}
		}
		return Outcome{
			Passed: false,
			Detail: fmt.Sprintf("saída não casa com /%s/: %q", c.Pattern, truncate(r.Stdout, 200)),
		}

	default:
		return Outcome{Passed: false, Detail: fmt.Sprintf("kind %q desconhecido", c.Kind)}
	}
}

// Label é o texto exibido na TUI para este critério.
func (c Criterion) Label() string {
	if c.Describe != "" {
		return c.Describe
	}
	switch c.Kind {
	case KindFileExists:
		return fmt.Sprintf("%s existe", c.Path)
	case KindExitCode:
		return fmt.Sprintf("`%s` sai com %d", strings.Join(c.Command, " "), c.Code)
	case KindStdoutEquals:
		return fmt.Sprintf("`%s` imprime a saída esperada", strings.Join(c.Command, " "))
	case KindStdoutMatches:
		return fmt.Sprintf("`%s` casa com /%s/", strings.Join(c.Command, " "), c.Pattern)
	default:
		return string(c.Kind)
	}
}

// Validate confere que o critério tem os campos que seu kind exige. A regex é
// compilada aqui para que um pattern quebrado falhe no boot, e não no meio de
// uma validação.
func (c Criterion) Validate() error {
	switch c.Kind {
	case KindFileExists:
		if c.Path == "" {
			return fmt.Errorf("kind %q exige path", c.Kind)
		}
		return validateRelPath(c.Path)

	case KindExitCode:
		return requireCommand(c)

	case KindStdoutEquals:
		if err := requireCommand(c); err != nil {
			return err
		}
		if c.Expected == "" {
			return fmt.Errorf("kind %q exige expected", c.Kind)
		}
		return nil

	case KindStdoutMatches:
		if err := requireCommand(c); err != nil {
			return err
		}
		if c.Pattern == "" {
			return fmt.Errorf("kind %q exige pattern", c.Kind)
		}
		if _, err := regexp.Compile(c.Pattern); err != nil {
			return fmt.Errorf("pattern inválido: %w", err)
		}
		return nil

	default:
		return fmt.Errorf("kind %q desconhecido (use file_exists, exit_code, stdout_equals ou stdout_matches)", c.Kind)
	}
}

func requireCommand(c Criterion) error {
	if len(c.Command) == 0 {
		return fmt.Errorf("kind %q exige command", c.Kind)
	}
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
