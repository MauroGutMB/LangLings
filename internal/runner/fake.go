package runner

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	"langlings/internal/domain"
)

// Fake é um Runner em memória para os testes rápidos.
//
// Existe para que o engine possa ser testado em milissegundos, sem Docker. O
// caminho real continua coberto pelos testes de integração de docker_test.go —
// o fake nunca é a única prova de que algo funciona.
type Fake struct {
	mu sync.Mutex

	// Images são as imagens consideradas presentes.
	Images map[string]bool

	// Results programa a resposta de um comando específico, pela chave
	// strings.Join(cmd, " ").
	Results map[string]ExecResult

	// ExecFunc, quando definido, tem prioridade sobre Results e permite
	// respostas dinâmicas (ex: ler o arquivo do workspace e decidir).
	ExecFunc func(cmd []string, opts ExecOpts) ExecResult

	// EphemeralFunc é o equivalente para o caminho efêmero. Recebe a spec
	// porque é o workspace dela que distingue uma execução da outra — é assim
	// que um teste do verify sabe se está olhando para base/ ou solution/.
	EphemeralFunc func(spec SessionSpec, cmd []string, opts ExecOpts) ExecResult

	// Default é a resposta quando nada mais casa.
	Default ExecResult

	// AvailableErr simula um daemon fora do ar.
	AvailableErr error

	// StartErr simula falha ao subir a sessão.
	StartErr error

	// Registros de uso, para asserções.
	Execs       [][]string
	Ephemerals  [][]string
	Started     []SessionSpec
	Stopped     []SessionID
	Pulled      []string
	CleanupHits int

	nextID int
}

var _ Runner = (*Fake)(nil)

// NewFake devolve um fake que aprova tudo por padrão.
func NewFake() *Fake {
	return &Fake{
		Images:  map[string]bool{},
		Results: map[string]ExecResult{},
	}
}

func (f *Fake) Available(context.Context) error { return f.AvailableErr }

func (f *Fake) ImageExists(_ context.Context, ref string) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.Images[ref], nil
}

func (f *Fake) EnsureImage(_ context.Context, lang domain.Language, _ string, progress io.Writer) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	ref := lang.ImageRef()
	if f.Images[ref] {
		return nil
	}
	f.Images[ref] = true
	f.Pulled = append(f.Pulled, ref)
	if progress != nil {
		fmt.Fprintf(progress, "fake: imagem %s pronta\n", ref)
	}
	return nil
}

func (f *Fake) StartSession(_ context.Context, spec SessionSpec) (SessionID, error) {
	if f.StartErr != nil {
		return "", f.StartErr
	}
	f.mu.Lock()
	defer f.mu.Unlock()

	f.nextID++
	f.Started = append(f.Started, spec)
	return SessionID(fmt.Sprintf("fake-session-%d", f.nextID)), nil
}

func (f *Fake) Exec(_ context.Context, _ SessionID, cmd []string, opts ExecOpts) (ExecResult, error) {
	f.mu.Lock()
	f.Execs = append(f.Execs, cmd)
	execFunc := f.ExecFunc
	result, programmed := f.Results[strings.Join(cmd, " ")]
	def := f.Default
	f.mu.Unlock()

	if execFunc != nil {
		return execFunc(cmd, opts), nil
	}
	if programmed {
		return result, nil
	}
	return def, nil
}

func (f *Fake) RunEphemeral(ctx context.Context, spec SessionSpec, cmd []string, opts ExecOpts) (ExecResult, error) {
	f.mu.Lock()
	f.Ephemerals = append(f.Ephemerals, cmd)
	ephemeralFunc := f.EphemeralFunc
	f.mu.Unlock()

	if ephemeralFunc != nil {
		return ephemeralFunc(spec, cmd, opts), nil
	}
	return f.Exec(ctx, "ephemeral", cmd, opts)
}

func (f *Fake) InteractiveCmd(id SessionID, lang domain.Language) *exec.Cmd {
	// `true` existe em qualquer sistema e sai com 0 imediatamente, o que
	// permite exercitar o fluxo de shell sem abrir shell nenhum.
	return exec.Command("true", string(id), lang.Slug)
}

func (f *Fake) Stop(_ context.Context, id SessionID) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Stopped = append(f.Stopped, id)
	return nil
}

func (f *Fake) CleanupOrphans(context.Context) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.CleanupHits++
	return 0, nil
}

// Program registra a resposta de um comando.
func (f *Fake) Program(cmd []string, res ExecResult) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Results[strings.Join(cmd, " ")] = res
}

// ExecCount devolve quantas execuções aconteceram.
func (f *Fake) ExecCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.Execs)
}
