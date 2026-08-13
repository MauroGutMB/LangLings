package watcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Tempos curtos para os testes rodarem em milissegundos. As janelas de espera
// são generosas de propósito: um teste de watcher que pisca por causa de
// escalonamento é pior que teste nenhum.
var testCfg = Config{Debounce: 60 * time.Millisecond, Poll: 25 * time.Millisecond}

const (
	janelaEvento   = 3 * time.Second
	janelaSilencio = 500 * time.Millisecond
)

type harness struct {
	t      *testing.T
	root   string
	events <-chan Event
}

func novoHarness(t *testing.T, allowlist []string, arquivosIniciais map[string]string) *harness {
	t.Helper()
	root := t.TempDir()

	for nome, conteudo := range arquivosIniciais {
		full := filepath.Join(root, nome)
		require.NoError(t, os.MkdirAll(filepath.Dir(full), 0o755))
		require.NoError(t, os.WriteFile(full, []byte(conteudo), 0o644))
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	events, err := Watch(ctx, root, allowlist, testCfg)
	require.NoError(t, err)

	// Dá tempo para o inotify registrar os diretórios antes do primeiro save.
	time.Sleep(50 * time.Millisecond)

	return &harness{t: t, root: root, events: events}
}

func (h *harness) escrever(nome, conteudo string) {
	h.t.Helper()
	full := filepath.Join(h.root, nome)
	require.NoError(h.t, os.MkdirAll(filepath.Dir(full), 0o755))
	require.NoError(h.t, os.WriteFile(full, []byte(conteudo), 0o644))
}

// salvarAtomicamente reproduz o que VSCode, Neovim e gofmt-on-save realmente
// fazem: escrevem um temporário e o renomeiam por cima do original. O inode
// muda, e um watch preso ao arquivo antigo morre em silêncio.
func (h *harness) salvarAtomicamente(nome, conteudo string) {
	h.t.Helper()
	full := filepath.Join(h.root, nome)
	tmp := full + ".tmp~"
	require.NoError(h.t, os.WriteFile(tmp, []byte(conteudo), 0o644))
	require.NoError(h.t, os.Rename(tmp, full))
}

func (h *harness) esperarEvento() Event {
	h.t.Helper()
	select {
	case ev, ok := <-h.events:
		require.True(h.t, ok, "canal de eventos fechou antes do esperado")
		return ev
	case <-time.After(janelaEvento):
		h.t.Fatal("nenhum evento em " + janelaEvento.String() + ": o watcher parou de reagir")
		return Event{}
	}
}

func (h *harness) exigirSilencio(motivo string) {
	h.t.Helper()
	select {
	case ev := <-h.events:
		h.t.Fatalf("evento inesperado (%s): %s vindo de %s", motivo, ev.At, ev.Source)
	case <-time.After(janelaSilencio):
	}
}

func TestWatch_SaveComumEmiteUmEvento(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	h.escrever("main.go", "package main // editado\n")

	h.esperarEvento()
	// Duas fontes observando o mesmo arquivo poderiam emitir duas vezes. O
	// retrato compartilhado é o que garante uma emissão por mudança.
	h.exigirSilencio("fsnotify e polling não podem emitir em dobro")
}

func TestWatch_RajadaDeSavesViraUmEvento(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	// Um único Ctrl+S costuma gerar de dois a quatro eventos de inotify.
	for i := 0; i < 4; i++ {
		h.escrever("main.go", "package main // versão "+string(rune('a'+i))+"\n")
		time.Sleep(10 * time.Millisecond)
	}

	h.esperarEvento()
	h.exigirSilencio("a rajada inteira precisa colapsar em uma validação")
}

func TestWatch_SalvamentoAtomicoContinuaSendoDetectado(t *testing.T) {
	// Este é o caso que quebra um watcher ingênuo: depois do primeiro save o
	// arquivo tem outro inode, e um watch no arquivo (em vez de no diretório)
	// deixaria de receber eventos para sempre.
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	for i := 0; i < 3; i++ {
		h.salvarAtomicamente("main.go", "package main // atômico "+string(rune('a'+i))+"\n")
		ev := h.esperarEvento()
		require.NotEmpty(t, ev.Source)
	}
}

func TestWatch_ArtefatoDeBuildEhIgnorado(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	// É o que um `go build` faz dentro do workspace. Sem a allowlist, isto
	// dispararia revalidação em cascata.
	h.escrever("bin/hello", "binário falso")
	h.escrever("go.sum", "hash falso")

	h.exigirSilencio("apenas os arquivos editáveis contam como save do usuário")
}

func TestWatch_ArquivoApagadoEmiteEvento(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	require.NoError(t, os.Remove(filepath.Join(h.root, "main.go")))

	// Apagar é uma mudança real: a validação vai falhar com mensagem clara e a
	// TUI pode oferecer restaurar do código-base.
	h.esperarEvento()
}

func TestWatch_ArquivoCriadoDepoisEmiteEvento(t *testing.T) {
	// O arquivo pode nem existir quando o watcher começa (exercício apagado e
	// recriado). Observar o diretório cobre isso.
	h := novoHarness(t, []string{"main.go"}, nil)

	h.escrever("main.go", "package main\n")

	h.esperarEvento()
}

func TestWatch_SaveImediatoAposIniciarNaoSePerde(t *testing.T) {
	// Regressão: o retrato inicial precisa ser tirado dentro de Watch, antes
	// de qualquer goroutine subir. Tirá-lo de forma assíncrona abre uma janela
	// em que um save chegado logo depois já entra no primeiro retrato — e,
	// como nada mais muda, o watcher fica mudo para sempre. Abrir o exercício
	// e salvar em seguida é o caminho mais natural do usuário.
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "main.go"), []byte("package main\n"), 0o644))

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	events, err := Watch(ctx, root, []string{"main.go"}, testCfg)
	require.NoError(t, err)

	// Sem pausa nenhuma: escreve no instante seguinte ao retorno de Watch.
	require.NoError(t, os.WriteFile(filepath.Join(root, "main.go"), []byte("package main // editado\n"), 0o644))

	select {
	case _, ok := <-events:
		require.True(t, ok)
	case <-time.After(janelaEvento):
		t.Fatal("o save logo após a abertura se perdeu: o watcher nasceu cego")
	}
}

func TestWatch_SemMudancaNaoEmiteNada(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	h.exigirSilencio("watcher recém-iniciado não pode inventar um save")
}

func TestWatch_VariosEditaveis(t *testing.T) {
	h := novoHarness(t, []string{"main.go", "sub/calc.go"}, map[string]string{
		"main.go":     "package main\n",
		"sub/calc.go": "package sub\n",
	})

	h.escrever("sub/calc.go", "package sub // editado\n")

	// Exercícios de "Exemplos reais" são multi-arquivo: editar qualquer um dos
	// declarados precisa disparar a validação.
	h.esperarEvento()
}

func TestWatch_SavesSeparadosEmitemSeparadamente(t *testing.T) {
	h := novoHarness(t, []string{"main.go"}, map[string]string{"main.go": "package main\n"})

	h.escrever("main.go", "package main // um\n")
	h.esperarEvento()

	h.escrever("main.go", "package main // dois\n")
	h.esperarEvento()
}

func TestWatch_AllowlistVazia(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := Watch(ctx, t.TempDir(), nil, testCfg)

	require.Error(t, err)
	require.Contains(t, err.Error(), "allowlist vazia")
}

func TestWatch_CancelamentoFechaOCanal(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "main.go"), []byte("x"), 0o644))

	ctx, cancel := context.WithCancel(context.Background())
	events, err := Watch(ctx, root, []string{"main.go"}, testCfg)
	require.NoError(t, err)

	cancel()

	select {
	case _, ok := <-events:
		require.False(t, ok, "o canal precisa fechar quando o contexto termina")
	case <-time.After(janelaEvento):
		t.Fatal("o canal não fechou depois do cancelamento")
	}
}

func TestConfig_Padroes(t *testing.T) {
	c := Config{}.withDefaults()
	require.Equal(t, DefaultDebounce, c.Debounce)
	require.Equal(t, DefaultPoll, c.Poll)
}

func TestSameSnapshot(t *testing.T) {
	agora := time.Now()
	base := map[string]fileState{"a": {exists: true, size: 10, mod: agora}}

	require.True(t, sameSnapshot(base, map[string]fileState{"a": {exists: true, size: 10, mod: agora}}))
	require.False(t, sameSnapshot(base, map[string]fileState{"a": {exists: true, size: 11, mod: agora}}))
	require.False(t, sameSnapshot(base, map[string]fileState{"a": {exists: false}}))
	require.False(t, sameSnapshot(base, map[string]fileState{"a": {exists: true, size: 10, mod: agora.Add(time.Second)}}))
	require.False(t, sameSnapshot(base, map[string]fileState{}))
}
