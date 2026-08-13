package engine

import (
	"context"
	"os"
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

// ---------- montagem de um ambiente de teste ----------

type ambiente struct {
	t       *testing.T
	engine  *Engine
	fake    *runner.Fake
	content string
	paths   Paths
}

const manifestoTeste = `
title     = "Append em slices"
objective = "Faça Dobrar devolver um slice novo."
editable  = ["main.go"]
hints     = ["append reaproveita o array subjacente."]

[validation]
mode    = "test"
command = ["go", "test", "./..."]
timeout = "30s"
`

const manifestoCriteria = `
title     = "Cross-compile"
objective = "Gere bin/hello.exe para windows/amd64."
editable  = ["main.go"]
hints     = ["GOOS e GOARCH."]

[validation]
mode    = "criteria"
timeout = "10s"

[[validation.criteria]]
kind = "file_exists"
path = "bin/hello.exe"
`

// novoAmbiente monta conteúdo, banco e engine com runner falso.
func novoAmbiente(t *testing.T, manifestos map[string]string) *ambiente {
	t.Helper()

	content := t.TempDir()
	data := t.TempDir()

	escrever(t, filepath.Join(content, "languages", "go", "language.toml"), `
name  = "Go"
image = "golang:1.26-alpine"
`)

	for caminho, manifesto := range manifestos {
		dir := filepath.Join(content, "exercises", filepath.FromSlash(caminho))
		escrever(t, filepath.Join(dir, "exercise.toml"), manifesto)
		escrever(t, filepath.Join(dir, "base", "main.go"), "package main // base\n")
		escrever(t, filepath.Join(dir, "solution", "main.go"), "package main // solução\n")
	}

	cat, err := catalog.Load(content)
	require.NoError(t, err)

	st, err := store.Open(filepath.Join(data, "langlings.db"))
	require.NoError(t, err)
	t.Cleanup(func() { st.Close() })

	fake := runner.NewFake()
	fake.Images["golang:1.26-alpine"] = true

	paths := Paths{
		ContentRoot: content,
		DataDir:     data,
		Workspace:   filepath.Join(data, "workspace"),
		DBPath:      filepath.Join(data, "langlings.db"),
		LogPath:     filepath.Join(data, "state", "langlings.log"),
	}

	e := New(cat, st, fake, paths)
	e.Watch = watcher.Config{Debounce: 40 * time.Millisecond, Poll: 20 * time.Millisecond}
	require.NoError(t, e.Bootstrap(context.Background()))

	return &ambiente{t: t, engine: e, fake: fake, content: content, paths: paths}
}

func escrever(t *testing.T, path, conteudo string) {
	t.Helper()
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(conteudo), 0o644))
}

func (a *ambiente) exercicio(path string) domain.Exercise {
	a.t.Helper()
	ex, ok := a.engine.Catalog.Exercise(path)
	require.True(a.t, ok, "exercício %q não está no catálogo", path)
	return ex
}

// aprovar e reprovar programam o veredito do runner falso.
func (a *ambiente) aprovar() { a.fake.Default = runner.ExecResult{ExitCode: 0, Stdout: "ok\n"} }
func (a *ambiente) reprovar() {
	a.fake.Default = runner.ExecResult{ExitCode: 1, Stdout: "FAIL\n", Stderr: "esperava 2, veio 1\n"}
}

// ---------- bootstrap e listagens ----------

func TestBootstrap_LimpaOrfaosEReconcilia(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})

	require.Equal(t, 1, a.fake.CleanupHits, "containers de um CLI morto precisam sumir no boot")

	p, err := a.engine.Store.Get("go/sintaxe/slices")
	require.NoError(t, err)
	require.Equal(t, domain.StatusNotStarted, p.Status)
}

func TestLanguages_InstaladaEhConsultadaAoVivo(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ctx := context.Background()

	views, err := a.engine.Languages(ctx)
	require.NoError(t, err)
	require.Len(t, views, 1)
	require.True(t, views[0].Installed)
	require.Equal(t, 1, views[0].Summary.Total)

	// Um `docker rmi` feito por fora precisa se refletir imediatamente.
	delete(a.fake.Images, "golang:1.26-alpine")

	views, err = a.engine.Languages(ctx)
	require.NoError(t, err)
	require.False(t, views[0].Installed)
}

func TestExercises_JuntaProgresso(t *testing.T) {
	a := novoAmbiente(t, map[string]string{
		"go/sintaxe/slices":   manifestoTeste,
		"go/compilador/cross": manifestoCriteria,
	})

	views, err := a.engine.Exercises("go")
	require.NoError(t, err)
	require.Len(t, views, 2)
	// Ordem de exibição: sintaxe antes de compilador.
	require.Equal(t, domain.CategorySintaxe, views[0].Exercise.Category)
	require.Equal(t, domain.StatusNotStarted, views[0].Progress.Status)
}

// ---------- workspace ----------

func TestMaterialize_NuncaSobrescreveOTrabalhoDoUsuario(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "main.go"), []byte("minha solução"), 0o644))

	// Voltar à tela do exercício não pode apagar o que a pessoa escreveu.
	dir2, err := a.paths.Materialize(ex)
	require.NoError(t, err)
	require.Equal(t, dir, dir2)

	conteudo, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Equal(t, "minha solução", string(conteudo))
}

func TestRestore_DescartaOTrabalhoDeProposito(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "main.go"), []byte("minha solução"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "lixo.txt"), []byte("artefato"), 0o644))

	_, err = a.paths.Restore(ex)
	require.NoError(t, err)

	conteudo, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Equal(t, "package main // base\n", string(conteudo))

	_, err = os.Stat(filepath.Join(dir, "lixo.txt"))
	require.True(t, os.IsNotExist(err), "restore precisa limpar artefatos também")
}

func TestReadEditable_ArquivoAusenteMudaOHash(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)

	comArquivo, err := ReadEditable(dir, ex)
	require.NoError(t, err)

	require.NoError(t, os.Remove(filepath.Join(dir, "main.go")))
	semArquivo, err := ReadEditable(dir, ex)
	require.NoError(t, err)

	require.NotEqual(t, domain.ContentHash(comArquivo), domain.ContentHash(semArquivo))
}

// ---------- validação ----------

func TestValidate_AprovacaoPersisteComoCompleto(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	a.aprovar()

	res := a.engine.Validate(context.Background(), ex)

	require.NoError(t, res.Err)
	require.True(t, res.Passed)
	require.Equal(t, domain.StatusCompleted, res.Progress.Status)

	// Precisa ter chegado ao banco, não só ao resultado em memória.
	salvo, err := a.engine.Store.Get(ex.Path)
	require.NoError(t, err)
	require.Equal(t, domain.StatusCompleted, salvo.Status)
	require.NotEmpty(t, salvo.LastContentHash)
}

func TestValidate_ReprovacaoViraEmProgressoComSaidaBruta(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	a.reprovar()

	res := a.engine.Validate(context.Background(), ex)

	require.NoError(t, res.Err)
	require.False(t, res.Passed)
	require.Equal(t, domain.StatusInProgress, res.Progress.Status)
	require.Contains(t, res.Output, "FAIL")
	require.Contains(t, res.Output, "esperava 2, veio 1", "a saída do compilador não pode ser resumida")
}

func TestValidate_UsaOComandoDoManifesto(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.aprovar()

	a.engine.Validate(context.Background(), a.exercicio("go/sintaxe/slices"))

	require.Equal(t, [][]string{{"go", "test", "./..."}}, a.fake.Execs)
}

func TestValidate_ModoCriteriaOlhaOArtefatoNoHost(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/compilador/cross": manifestoCriteria})
	ex := a.exercicio("go/compilador/cross")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)

	// Sem o artefato, reprova.
	res := a.engine.Validate(context.Background(), ex)
	require.NoError(t, res.Err)
	require.False(t, res.Passed)
	require.Len(t, res.Criteria, 1)
	require.Contains(t, res.Criteria[0].Outcome.Detail, "não foi encontrado")
	require.Empty(t, a.fake.Execs, "file_exists não precisa de container: o volume já está montado no host")

	// O usuário compila à mão no shell e o artefato aparece.
	escrever(t, filepath.Join(dir, "bin", "hello.exe"), "MZ...")

	res = a.engine.Validate(context.Background(), ex)
	require.True(t, res.Passed)
	require.Equal(t, domain.StatusCompleted, res.Progress.Status)
}

func TestValidate_TimeoutDescartaASessao(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	a.fake.Default = runner.ExecResult{ExitCode: -1, TimedOut: true}

	res := a.engine.Validate(context.Background(), ex)

	require.True(t, res.TimedOut)
	require.False(t, res.Passed)
	// O processo lá dentro pode ter sobrevivido ao cliente: a sessão deixa de
	// ser confiável e precisa ser jogada fora.
	require.Len(t, a.fake.Stopped, 1)

	a.aprovar()
	a.engine.Validate(context.Background(), ex)
	require.Len(t, a.fake.Started, 2, "a validação seguinte precisa subir uma sessão nova")
}

func TestValidate_ExercicioComRedeUsaContainerEfemero(t *testing.T) {
	manifesto := `
title     = "Com dependência externa"
objective = "Use uma biblioteca de terceiros."
editable  = ["main.go"]
hints     = ["go get"]

[validation]
mode    = "test"
command = ["go", "test", "./..."]
network = true
`
	a := novoAmbiente(t, map[string]string{"go/frameworks/chi": manifesto})
	a.aprovar()

	a.engine.Validate(context.Background(), a.exercicio("go/frameworks/chi"))

	// A sessão nasce isolada e limites de rede são fixados na criação, então
	// este exercício não pode ser servido por ela.
	require.Len(t, a.fake.Ephemerals, 1)
	require.Empty(t, a.fake.Started, "nenhuma sessão deveria ter subido")
}

func TestSessao_EhReaproveitadaEntreValidacoes(t *testing.T) {
	a := novoAmbiente(t, map[string]string{
		"go/sintaxe/slices": manifestoTeste,
		"go/sintaxe/maps":   manifestoTeste,
	})
	a.aprovar()

	a.engine.Validate(context.Background(), a.exercicio("go/sintaxe/slices"))
	a.engine.Validate(context.Background(), a.exercicio("go/sintaxe/maps"))

	// Um container por linguagem, não por exercício: é o que mantém o loop de
	// save em centenas de milissegundos.
	require.Len(t, a.fake.Started, 1)
	require.Equal(t, 2, a.fake.ExecCount())
}

func TestShutdown_ParaAsSessoes(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.aprovar()
	a.engine.Validate(context.Background(), a.exercicio("go/sintaxe/slices"))

	a.engine.Shutdown(context.Background())

	require.Len(t, a.fake.Stopped, 1)
}

// ---------- loop de save ----------

func TestOpen_SalvarDisparaValidacao(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	a.reprovar()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess, err := a.engine.Open(ctx, ex)
	require.NoError(t, err)
	require.Equal(t, a.paths.ExerciseDir(ex), sess.Dir, "o path exibido é o do host, não o do container")

	time.Sleep(50 * time.Millisecond)
	escrever(t, filepath.Join(sess.Dir, "main.go"), "package main // tentativa 1\n")

	res := esperarResultado(t, sess)
	require.False(t, res.Passed)
	require.Equal(t, domain.StatusInProgress, res.Progress.Status)

	// Agora o usuário acerta.
	a.aprovar()
	escrever(t, filepath.Join(sess.Dir, "main.go"), "package main // tentativa 2\n")

	res = esperarResultado(t, sess)
	require.True(t, res.Passed)
	require.Equal(t, domain.StatusCompleted, res.Progress.Status)
}

func TestOpen_ConteudoIdenticoNaoRevalida(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	a.aprovar()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess, err := a.engine.Open(ctx, ex)
	require.NoError(t, err)

	time.Sleep(50 * time.Millisecond)
	escrever(t, filepath.Join(sess.Dir, "main.go"), "package main // igual\n")
	esperarResultado(t, sess)

	execsDepoisDoPrimeiro := a.fake.ExecCount()

	// Salvar o mesmo conteúdo de novo (Ctrl+S sem editar nada) muda o mtime e
	// acorda o watcher, mas não há o que validar.
	for i := 0; i < 3; i++ {
		escrever(t, filepath.Join(sess.Dir, "main.go"), "package main // igual\n")
		time.Sleep(60 * time.Millisecond)
	}

	semResultado(t, sess)
	require.Equal(t, execsDepoisDoPrimeiro, a.fake.ExecCount(),
		"o hash é o árbitro: mesmo conteúdo não gasta um container")
}

func TestOpen_CancelamentoFechaOCanal(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ctx, cancel := context.WithCancel(context.Background())

	sess, err := a.engine.Open(ctx, a.exercicio("go/sintaxe/slices"))
	require.NoError(t, err)

	cancel()

	select {
	case _, ok := <-sess.Updates:
		require.False(t, ok)
	case <-time.After(3 * time.Second):
		t.Fatal("o canal de resultados não fechou")
	}
}

// esperarResultado consome updates até o desfecho, exigindo que o aviso de
// "validando" venha antes — é dele que depende o spinner da TUI.
func esperarResultado(t *testing.T, s *Session) Result {
	t.Helper()
	return esperarResultadoAte(t, s, 5*time.Second)
}

func esperarResultadoAte(t *testing.T, s *Session, limite time.Duration) Result {
	t.Helper()
	prazo := time.After(limite)
	viuValidando := false

	for {
		select {
		case up, ok := <-s.Updates:
			require.True(t, ok, "canal de updates fechou antes do esperado")
			if up.Kind == UpdateValidating {
				viuValidando = true
				continue
			}
			require.True(t, viuValidando, "o aviso de validando precisa preceder o resultado")
			require.NoError(t, up.Result.Err)
			return up.Result
		case <-prazo:
			t.Fatal("nenhuma validação disparou: o loop de save parou de reagir")
			return Result{}
		}
	}
}

func semResultado(t *testing.T, s *Session) {
	t.Helper()
	select {
	case up := <-s.Updates:
		t.Fatalf("update inesperado: kind=%v passed=%v", up.Kind, up.Result.Passed)
	case <-time.After(400 * time.Millisecond):
	}
}
