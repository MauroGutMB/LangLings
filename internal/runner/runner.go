// Package runner executa comandos dentro de containers Docker.
//
// A interface Runner é deliberadamente estreita. É ela que permite ao engine
// ser testado em milissegundos com um fake, enquanto a implementação real é
// exercitada por testes de integração contra o Docker de verdade.
package runner

import (
	"context"
	"io"
	"os/exec"
	"time"

	"langlings/internal/domain"
)

// SessionID identifica um container de sessão vivo.
type SessionID string

// Label aplicado a todo container criado pelo LangLings. É o que permite
// encontrar e remover órfãos deixados por um CLI que morreu de forma abrupta.
const (
	ManagedLabel     = "langlings.managed"
	ManagedLabelTrue = ManagedLabel + "=true"
	LanguageLabel    = "langlings.language"
)

// Limites aplicados a todo container. São fixados na criação, o que é
// justamente por que um exercício com rede não pode ser servido pela sessão.
const (
	DefaultMemory = "512m"
	DefaultCPUs   = "2"
)

// SessionSpec descreve o container de uma linguagem.
type SessionSpec struct {
	Language  domain.Language
	Workspace string // diretório no host montado como Workdir

	// Network libera a rede. Sessões sobem sempre isoladas; este campo existe
	// para o caminho efêmero, usado por exercícios que declaram network=true.
	Network bool
}

// ExecOpts ajusta uma execução.
type ExecOpts struct {
	Workdir string        // padrão: o workdir da linguagem
	Timeout time.Duration // 0 = sem limite
}

// ExecResult é o desfecho de uma execução.
type ExecResult struct {
	ExitCode int
	Stdout   string
	Stderr   string

	// TimedOut informa que o comando estourou o tempo. É diferente de exit != 0:
	// o processo dentro do container pode ter sobrevivido ao cliente que o
	// invocou, então quem recebe isso deve descartar a sessão em vez de
	// reaproveitá-la.
	TimedOut bool
}

// Output junta stdout e stderr para exibição na TUI, que é onde o usuário
// realmente quer ler a saída do compilador.
func (r ExecResult) Output() string {
	switch {
	case r.Stdout == "":
		return r.Stderr
	case r.Stderr == "":
		return r.Stdout
	default:
		return r.Stdout + "\n" + r.Stderr
	}
}

// Runner executa trabalho dentro de containers.
type Runner interface {
	// Available informa se o Docker está acessível. Um daemon fora do ar deve
	// virar mensagem clara, não pânico no meio de uma validação.
	Available(ctx context.Context) error

	// ImageExists é o que define uma linguagem como "instalada". É consultado
	// ao vivo, nunca persistido.
	ImageExists(ctx context.Context, ref string) (bool, error)

	// EnsureImage baixa ou constrói a imagem, transmitindo o log para progress.
	EnsureImage(ctx context.Context, lang domain.Language, contextDir string, progress io.Writer) error

	// StartSession sobe o container de longa duração de uma linguagem.
	StartSession(ctx context.Context, spec SessionSpec) (SessionID, error)

	// Exec roda um comando numa sessão viva. É o caminho do loop de save.
	Exec(ctx context.Context, id SessionID, cmd []string, opts ExecOpts) (ExecResult, error)

	// RunEphemeral roda um comando num container descartável. É o caminho dos
	// exercícios que precisam de rede, que a sessão isolada não pode servir.
	RunEphemeral(ctx context.Context, spec SessionSpec, cmd []string, opts ExecOpts) (ExecResult, error)

	// InteractiveCmd monta (sem executar) o comando de shell interativo.
	// Devolver um *exec.Cmd em vez de executar é o que permite entregá-lo ao
	// tea.ExecProcess, que precisa assumir o controle do TTY.
	InteractiveCmd(id SessionID, lang domain.Language) *exec.Cmd

	// Stop remove um container de sessão.
	Stop(ctx context.Context, id SessionID) error

	// CleanupOrphans remove containers rotulados que sobraram de execuções
	// anteriores. Chamado no boot.
	CleanupOrphans(ctx context.Context) (int, error)
}
