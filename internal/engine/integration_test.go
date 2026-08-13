package engine

import (
	"context"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"langlings/internal/catalog"
	"langlings/internal/domain"
	"langlings/internal/runner"
	"langlings/internal/store"
	"langlings/internal/watcher"
)

// Estes testes exercitam o caminho inteiro com Docker de verdade: container,
// volume de cache, compilação, arquivos no host. São a prova de que o fake não
// está mentindo.

const imagemIntegracao = "golang:1.26-alpine"

const linguagemGoTOML = `
name      = "Go"
image     = "golang:1.26-alpine"
workdir   = "/workspace"
cache_dir = "/cache"

[env]
GOCACHE    = "/cache/build"
GOMODCACHE = "/cache/mod"
GOFLAGS    = "-mod=mod"
`

const exercicioDobrarTOML = `
title     = "Append em slices"
objective = "Faça Dobrar devolver um slice novo com cada elemento multiplicado por 2."
editable  = ["main.go"]
hints     = ["make([]int, 0, len(xs)) te dá um destino novo."]

[validation]
mode    = "test"
command = ["go", "test", "./..."]
timeout = "120s"
`

const baseDobrar = `package main

func Dobrar(xs []int) []int {
	return xs
}

func main() {}
`

const solucaoDobrar = `package main

func Dobrar(xs []int) []int {
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		out = append(out, x*2)
	}
	return out
}

func main() {}
`

const testeDobrar = `package main

import "testing"

func TestDobrar(t *testing.T) {
	original := []int{1, 2, 3}
	got := Dobrar(original)

	want := []int{2, 4, 6}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Dobrar(%v) = %v, quero %v", original, got, want)
		}
	}
	if original[0] != 1 {
		t.Fatalf("Dobrar não pode modificar o slice original: %v", original)
	}
}
`

// ambienteReal monta o mesmo cenário do fake, mas com o runner Docker.
func ambienteReal(t *testing.T) *ambiente {
	t.Helper()
	if testing.Short() {
		t.Skip("requer Docker; rode sem -short")
	}

	d := runner.NewDocker()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	require.NoError(t, d.Available(ctx))

	existe, err := d.ImageExists(ctx, imagemIntegracao)
	require.NoError(t, err)
	if !existe {
		t.Skipf("imagem %s ausente; rode `docker pull %s`", imagemIntegracao, imagemIntegracao)
	}

	content := t.TempDir()
	data := t.TempDir()

	escrever(t, filepath.Join(content, "languages", "go", "language.toml"), linguagemGoTOML)

	exDir := filepath.Join(content, "exercises", "go", "sintaxe", "dobrar")
	escrever(t, filepath.Join(exDir, "exercise.toml"), exercicioDobrarTOML)
	escrever(t, filepath.Join(exDir, "base", "go.mod"), "module dobrar\n\ngo 1.24\n")
	escrever(t, filepath.Join(exDir, "base", "main.go"), baseDobrar)
	escrever(t, filepath.Join(exDir, "base", "main_test.go"), testeDobrar)
	escrever(t, filepath.Join(exDir, "solution", "main.go"), solucaoDobrar)

	cat, err := catalog.Load(content)
	require.NoError(t, err)

	st, err := store.Open(filepath.Join(data, "langlings.db"))
	require.NoError(t, err)
	t.Cleanup(func() { st.Close() })

	paths := Paths{
		ContentRoot: content,
		DataDir:     data,
		Workspace:   filepath.Join(data, "workspace"),
		DBPath:      filepath.Join(data, "langlings.db"),
		LogPath:     filepath.Join(data, "state", "langlings.log"),
	}

	e := New(cat, st, d, paths)
	e.Watch = watcher.Config{Debounce: 100 * time.Millisecond, Poll: 100 * time.Millisecond}
	require.NoError(t, e.Bootstrap(context.Background()))

	t.Cleanup(func() {
		e.Shutdown(context.Background())
		// O volume de cache é compartilhado entre os testes desta suíte; só
		// some no fim, para não pagar 12s de recompilação da stdlib em cada um.
		exec.Command("docker", "volume", "rm", "langlings-cache-go").Run()
	})

	return &ambiente{t: t, engine: e, content: content, paths: paths}
}

func TestIntegracao_ValidacaoRealReprovaOCodigoBase(t *testing.T) {
	a := ambienteReal(t)
	ex := a.exercicio("go/sintaxe/dobrar")

	res := a.engine.Validate(context.Background(), ex)

	require.NoError(t, res.Err)
	require.False(t, res.Passed, "o código-base precisa reprovar: %s", res.Output)
	require.Contains(t, res.Output, "quero [2 4 6]", "a saída bruta do go test precisa chegar ao usuário")
	require.Equal(t, domain.StatusInProgress, res.Progress.Status)
}

func TestIntegracao_ValidacaoRealAprovaASolucao(t *testing.T) {
	a := ambienteReal(t)
	ex := a.exercicio("go/sintaxe/dobrar")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)
	escrever(t, filepath.Join(dir, "main.go"), solucaoDobrar)

	res := a.engine.Validate(context.Background(), ex)

	require.NoError(t, res.Err)
	require.True(t, res.Passed, "saída: %s", res.Output)
	require.Equal(t, domain.StatusCompleted, res.Progress.Status)

	salvo, err := a.engine.Store.Get(ex.Path)
	require.NoError(t, err)
	require.Equal(t, domain.StatusCompleted, salvo.Status)
	require.NotNil(t, salvo.FirstCompletedAt)
}

// TestIntegracao_LoopDeSave é o produto inteiro num teste: o usuário edita o
// arquivo no host, não roda comando nenhum, e o estado muda sozinho.
func TestIntegracao_LoopDeSave(t *testing.T) {
	a := ambienteReal(t)
	ex := a.exercicio("go/sintaxe/dobrar")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess, err := a.engine.Open(ctx, ex)
	require.NoError(t, err)
	require.DirExists(t, sess.Dir)

	// Primeira tentativa: errada.
	escrever(t, filepath.Join(sess.Dir, "main.go"), `package main

func Dobrar(xs []int) []int {
	for i := range xs {
		xs[i] *= 2
	}
	return xs
}

func main() {}
`)

	res := esperarResultadoLento(t, sess)
	require.False(t, res.Passed, "modificar o slice original tem que reprovar")
	require.Equal(t, domain.StatusInProgress, res.Progress.Status)

	// Segunda tentativa: correta.
	escrever(t, filepath.Join(sess.Dir, "main.go"), solucaoDobrar)

	res = esperarResultadoLento(t, sess)
	require.True(t, res.Passed, "saída: %s", res.Output)
	require.Equal(t, domain.StatusCompleted, res.Progress.Status)
}

func TestIntegracao_LatenciaDoLoop(t *testing.T) {
	a := ambienteReal(t)
	ex := a.exercicio("go/sintaxe/dobrar")
	ctx := context.Background()

	// Primeira validação aquece o cache do toolchain.
	a.engine.Validate(ctx, ex)

	dir := a.paths.ExerciseDir(ex)
	escrever(t, filepath.Join(dir, "main.go"), solucaoDobrar)

	inicio := time.Now()
	res := a.engine.Validate(ctx, ex)
	decorrido := time.Since(inicio)

	require.NoError(t, res.Err)
	t.Logf("validação com cache quente: %v", decorrido.Round(time.Millisecond))
	require.Less(t, decorrido, 10*time.Second,
		"o loop de save precisa ser rápido o bastante para não parecer travado")
}

func TestIntegracao_VerifyComDockerReal(t *testing.T) {
	// O gate de conteúdo exercitado de verdade: base/ compila e reprova,
	// solution/ compila e aprova, tudo em containers descartáveis.
	a := ambienteReal(t)

	report := a.engine.Verify(context.Background(), a.exercicio("go/sintaxe/dobrar"))

	require.True(t, report.OK(), "problemas: %v", report.Problems)
}

func TestIntegracao_ShellCommandApontaParaOExercicio(t *testing.T) {
	a := ambienteReal(t)
	ex := a.exercicio("go/sintaxe/dobrar")

	cmd, err := a.engine.ShellCommand(context.Background(), ex)
	require.NoError(t, err)

	require.Contains(t, cmd.Args, "--tty")
	require.Contains(t, cmd.Args, "/workspace/go/sintaxe/dobrar")
	require.Nil(t, cmd.Process, "o shell só é executado pela TUI, via tea.ExecProcess")
}

func esperarResultadoLento(t *testing.T, s *Session) Result {
	t.Helper()
	return esperarResultadoAte(t, s, 3*time.Minute)
}
