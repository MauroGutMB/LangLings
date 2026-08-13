package engine

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"langlings/internal/domain"
)

// VerifyReport é o resultado do gate de qualidade de um exercício.
type VerifyReport struct {
	Exercise domain.Exercise
	Problems []string
}

// OK informa se o exercício passou no gate.
func (r VerifyReport) OK() bool { return len(r.Problems) == 0 }

func (r *VerifyReport) fail(format string, args ...any) {
	r.Problems = append(r.Problems, fmt.Sprintf(format, args...))
}

// Verify confere que um exercício é, de fato, um exercício.
//
// Este é o gate que transforma "auditar o conteúdo gerado" de leitura manual
// de dezenas de diretórios num comando com exit code. Sem ele, conteúdo escrito
// em lote é indistinguível de conteúdo quebrado — e os três defeitos típicos
// (código-base que já passa, solução que exige editar o que o usuário não pode
// editar, e solução que não passa) são todos invisíveis a olho nu.
//
// A validação roda sobre diretórios temporários: verificar conteúdo nunca pode
// encostar no trabalho do usuário.
func (e *Engine) Verify(ctx context.Context, ex domain.Exercise) VerifyReport {
	report := VerifyReport{Exercise: ex}

	if len(ex.Hints) == 0 {
		report.fail("nenhum hint: um exercício sem dica é um beco sem saída")
	}

	if _, err := os.Stat(ex.SolutionDir()); err != nil {
		report.fail("diretório solution/ ausente: sem ele o exercício não é verificável")
		return report
	}

	solucao, err := listFiles(ex.SolutionDir())
	if err != nil {
		report.fail("lendo solution/: %v", err)
		return report
	}

	if ex.Validation.Mode == domain.ModeCriteria {
		// Num exercício de Compilador a solução pode não alterar arquivo
		// nenhum: o que se resolve é uma linha de comando. O que ela precisa
		// ter é o roteiro desses comandos.
		if _, err := os.Stat(filepath.Join(ex.SolutionDir(), SolutionStepsFile)); err != nil {
			report.fail("modo criteria exige solution/%s com os comandos que satisfazem os critérios", SolutionStepsFile)
			return report
		}
	} else if len(solucao) == 0 {
		report.fail("solution/ está vazio")
		return report
	}

	// Toda diferença entre base e solução precisa caber na allowlist. Se a
	// solução mexeu num arquivo que o usuário não pode editar, o exercício é
	// insolúvel — e ninguém descobre isso até tentar resolvê-lo.
	editable := map[string]bool{}
	for _, f := range ex.Editable {
		editable[f] = true
	}
	for _, f := range solucao {
		if !editable[f] {
			report.fail("solution/%s não está em editable: o usuário não teria como alterá-lo", f)
		}
	}

	// O código-base precisa REPROVAR. Um exercício cujo código inicial já
	// passa é um exercício sem exercício, e é o defeito mais provável em
	// conteúdo escrito em lote.
	baseRes, err := e.validateInTemp(ctx, ex, false)
	switch {
	case err != nil:
		report.fail("validando base/: %v", err)
	case baseRes.Err != nil:
		report.fail("validando base/: %v", baseRes.Err)
	case baseRes.Passed:
		report.fail("base/ já passa na validação: não há nada para o usuário resolver")
	}

	// A solução precisa APROVAR, ou o exercício é impossível.
	solRes, err := e.validateInTemp(ctx, ex, true)
	switch {
	case err != nil:
		report.fail("validando solution/: %v", err)
	case solRes.Err != nil:
		report.fail("validando solution/: %v", solRes.Err)
	case !solRes.Passed:
		report.fail("solution/ não passa na validação: o exercício é insolúvel\n%s",
			indent(strings.TrimSpace(solRes.Output)))
	}

	return report
}

// VerifyAll roda o gate sobre o catálogo inteiro.
func (e *Engine) VerifyAll(ctx context.Context) []VerifyReport {
	out := make([]VerifyReport, 0, len(e.Catalog.Exercises))
	for _, ex := range e.Catalog.Exercises {
		out = append(out, e.Verify(ctx, ex))
	}
	return out
}

// validateInTemp monta o exercício num workspace descartável e o valida.
//
// A solução é aplicada como sobreposição do código-base, não como substituta:
// solution/ guarda só os arquivos que mudam, enquanto go.mod, arquivos de
// teste e o resto continuam vindo de base/. Isso mantém a solução legível como
// um diff do que o usuário precisa escrever.
func (e *Engine) validateInTemp(ctx context.Context, ex domain.Exercise, comSolucao bool) (Result, error) {
	root, err := os.MkdirTemp("", "langlings-verify-")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(root)

	dest := filepath.Join(root, filepath.FromSlash(ex.Path))
	if err := copyDir(ex.BaseDir(), dest); err != nil {
		return Result{}, fmt.Errorf("copiando base/: %w", err)
	}
	if comSolucao {
		if err := copyDir(ex.SolutionDir(), dest); err != nil {
			return Result{}, fmt.Errorf("aplicando solution/: %w", err)
		}
	}

	// Modo criteria não é verificável assim: os artefatos que ele observa vêm
	// de comandos que o usuário digita no shell, não de um arquivo editado.
	// A solução de referência precisa dizer como produzi-los.
	if ex.Validation.Mode == domain.ModeCriteria {
		if err := e.runSolutionSteps(ctx, ex, root, comSolucao); err != nil {
			return Result{}, err
		}
	}

	return e.runValidation(ctx, ex, root, ex.Path), nil
}

// SolutionStepsFile é o arquivo opcional, dentro de solution/, que descreve os
// comandos que o usuário digitaria no shell para satisfazer os critérios.
//
// Existe só para a categoria Compilador/Interpretador, onde não há "arquivo
// correto" a comparar: o que se verifica é o efeito colateral de um comando.
const SolutionStepsFile = "steps.sh"

// StepsTimeout é a folga dada aos comandos da solução de referência, que
// costumam compilar de verdade.
const StepsTimeout = 3 * time.Minute

// runSolutionSteps executa os passos manuais da solução dentro do container.
// Sem solução aplicada, os passos não rodam — é isso que faz base/ reprovar.
func (e *Engine) runSolutionSteps(ctx context.Context, ex domain.Exercise, root string, comSolucao bool) error {
	if !comSolucao {
		return nil
	}

	steps := filepath.Join(ex.SolutionDir(), SolutionStepsFile)
	content, err := os.ReadFile(steps)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("exercício de modo criteria exige solution/%s com os comandos que satisfazem os critérios", SolutionStepsFile)
		}
		return err
	}

	lang, ok := e.Catalog.Language(ex.Language)
	if !ok {
		return fmt.Errorf("linguagem %q não encontrada", ex.Language)
	}

	// O timeout do exercício dimensiona os comandos de checagem (statar um
	// arquivo, rodar um binário), não uma compilação. Os passos da solução
	// ganham uma folga própria.
	stepsSpec := ex.Validation
	stepsSpec.Timeout = StepsTimeout

	run := e.executor(ctx, lang, root, containerDir(lang, ex), stepsSpec)
	res, err := run([]string{"sh", "-c", string(content)})
	if err != nil {
		return err
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("solution/%s falhou (exit %d):\n%s", SolutionStepsFile, res.ExitCode, indent(res.Output()))
	}
	return nil
}

// listFiles devolve os caminhos relativos dos arquivos regulares de um
// diretório, ignorando o arquivo de passos da solução.
func listFiles(dir string) ([]string, error) {
	var out []string

	err := filepath.Walk(dir, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() || !fi.Mode().IsRegular() {
			return nil
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == SolutionStepsFile {
			return nil
		}
		out = append(out, rel)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(out)
	return out, nil
}

func indent(s string) string {
	if s == "" {
		return ""
	}
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = "    " + l
	}
	return strings.Join(lines, "\n")
}
