package runner

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"langlings/internal/domain"
)

// ---------- rápidos: montagem de comando, sem Docker ----------

func linguagemDeTeste() domain.Language {
	return domain.Language{
		Slug:     "itest",
		Name:     "Integração",
		Image:    "golang:1.26-alpine",
		Workdir:  "/workspace",
		Shell:    []string{"/bin/sh"},
		CacheDir: "/cache",
		Env:      map[string]string{"GOCACHE": "/cache/build", "GOMODCACHE": "/cache/mod"},
	}
}

func TestContainerArgs(t *testing.T) {
	d := &Docker{Binary: "docker", UID: 1000, GID: 1000}
	args := strings.Join(d.containerArgs(SessionSpec{
		Language:  linguagemDeTeste(),
		Workspace: "/home/user/ws",
	}), " ")

	t.Run("identidade do host", func(t *testing.T) {
		// Sem isso, todo artefato gerado no container vira root no host.
		require.Contains(t, args, "--user 1000:1000")
	})

	t.Run("isolamento por padrão", func(t *testing.T) {
		require.Contains(t, args, "--network none")
		require.Contains(t, args, "--memory "+DefaultMemory)
		require.Contains(t, args, "--cpus "+DefaultCPUs)
	})

	t.Run("workspace montado no workdir da linguagem", func(t *testing.T) {
		require.Contains(t, args, "--volume /home/user/ws:/workspace")
		require.Contains(t, args, "--workdir /workspace")
	})

	t.Run("labels permitem encontrar órfãos", func(t *testing.T) {
		require.Contains(t, args, "--label "+ManagedLabelTrue)
		require.Contains(t, args, "--label "+LanguageLabel+"=itest")
	})

	t.Run("HOME gravável e env da linguagem", func(t *testing.T) {
		require.Contains(t, args, "--env HOME=/tmp")
		require.Contains(t, args, "--env GOCACHE=/cache/build")
		require.Contains(t, args, "--env GOMODCACHE=/cache/mod")
	})

	t.Run("volume nomeado de cache", func(t *testing.T) {
		require.Contains(t, args, "--volume langlings-cache-itest:/cache")
	})
}

func TestContainerArgs_RedeDeclarada(t *testing.T) {
	d := &Docker{UID: 1000, GID: 1000}

	comRede := strings.Join(d.containerArgs(SessionSpec{
		Language: linguagemDeTeste(), Workspace: "/ws", Network: true,
	}), " ")

	require.Contains(t, comRede, "--network bridge")
	require.NotContains(t, comRede, "--network none")
}

func TestContainerArgs_EnvOrdenado(t *testing.T) {
	// Ordem estável mantém o comando reprodutível: o mesmo estado sempre gera
	// a mesma linha de comando, que é o que torna o log copiável e colável.
	d := &Docker{UID: 1000, GID: 1000}
	lang := linguagemDeTeste()
	lang.Env = map[string]string{"ZZZ": "1", "AAA": "2", "MMM": "3"}

	for i := 0; i < 20; i++ {
		args := strings.Join(d.containerArgs(SessionSpec{Language: lang, Workspace: "/ws"}), " ")
		require.Contains(t, args, "--env AAA=2 --env HOME=/tmp --env MMM=3 --env ZZZ=1")
	}
}

func TestContainerArgs_SemCacheDirNaoMontaVolume(t *testing.T) {
	d := &Docker{UID: 1000, GID: 1000}
	lang := linguagemDeTeste()
	lang.CacheDir = ""

	args := strings.Join(d.containerArgs(SessionSpec{Language: lang, Workspace: "/ws"}), " ")

	require.NotContains(t, args, "langlings-cache")
}

func TestImageArgs_ImagemPronta(t *testing.T) {
	require.Equal(t,
		[]string{"pull", "golang:1.26-alpine"},
		imageArgs(linguagemDeTeste(), "/repo"))
}

func TestImageArgs_DockerfileEhResolvidoContraOConteudo(t *testing.T) {
	lang := domain.Language{Slug: "lua", Name: "Lua", Dockerfile: "languages/lua/Dockerfile"}

	args := imageArgs(lang, "/repo/LangLings")

	// O caminho do manifesto é relativo à raiz do conteúdo. Deixá-lo cru faria
	// o docker resolvê-lo contra o CWD do processo, e o build quebraria ao
	// rodar de uma subpasta ou com -content.
	require.Equal(t, []string{
		"build", "-t", "langlings/lua:latest",
		"-f", "/repo/LangLings/languages/lua/Dockerfile",
		"/repo/LangLings/languages/lua",
	}, args)
}

func TestInteractiveCmd_NaoExecuta(t *testing.T) {
	d := &Docker{Binary: "docker", UID: 1000, GID: 1000}

	cmd := d.InteractiveCmd("abc123", linguagemDeTeste())

	// Devolver o comando montado em vez de executá-lo é o que permite entregar
	// o TTY ao tea.ExecProcess.
	require.Equal(t, "docker", filepath.Base(cmd.Path))
	require.Equal(t,
		[]string{"docker", "exec", "--interactive", "--tty", "--workdir", "/workspace", "abc123", "/bin/sh"},
		cmd.Args)
	require.Nil(t, cmd.Process, "o comando não pode ter sido iniciado")
}

func TestExecResult_Output(t *testing.T) {
	require.Equal(t, "só stdout", ExecResult{Stdout: "só stdout"}.Output())
	require.Equal(t, "só stderr", ExecResult{Stderr: "só stderr"}.Output())
	require.Equal(t, "a\nb", ExecResult{Stdout: "a", Stderr: "b"}.Output())
}

func TestExec_ComandoVazio(t *testing.T) {
	d := NewDocker()
	_, err := d.Exec(context.Background(), "x", nil, ExecOpts{})
	require.Error(t, err)
}

// ---------- integração: Docker de verdade ----------

const imagemIntegracao = "golang:1.26-alpine"

// exigeDocker pula o teste no modo rápido e falha se o Docker estiver ausente
// quando o modo completo foi pedido.
func exigeDocker(t *testing.T) *Docker {
	t.Helper()
	if testing.Short() {
		t.Skip("requer Docker; rode sem -short")
	}

	d := NewDocker()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	require.NoError(t, d.Available(ctx), "docker precisa estar acessível")

	existe, err := d.ImageExists(ctx, imagemIntegracao)
	require.NoError(t, err)
	if !existe {
		t.Skipf("imagem %s ausente; rode `docker pull %s`", imagemIntegracao, imagemIntegracao)
	}
	return d
}

// sessaoDeTeste sobe uma sessão descartável e garante sua remoção.
func sessaoDeTeste(t *testing.T, d *Docker, ws string) (SessionID, domain.Language) {
	t.Helper()
	lang := linguagemDeTeste()
	ctx := context.Background()

	id, err := d.StartSession(ctx, SessionSpec{Language: lang, Workspace: ws})
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, d.Stop(context.Background(), id))
		// O volume de cache é criado pela sessão; some junto com o teste.
		exec.Command("docker", "volume", "rm", lang.CacheVolume()).Run()
	})
	return id, lang
}

func TestIntegracao_ExecNaSessao(t *testing.T) {
	d := exigeDocker(t)
	ws := t.TempDir()
	id, _ := sessaoDeTeste(t, d, ws)

	res, err := d.Exec(context.Background(), id, []string{"echo", "olá"}, ExecOpts{})

	require.NoError(t, err)
	require.Equal(t, 0, res.ExitCode)
	require.Equal(t, "olá\n", res.Stdout)
	require.False(t, res.TimedOut)
}

func TestIntegracao_ExitCodeEStderrSaoDistintos(t *testing.T) {
	d := exigeDocker(t)
	id, _ := sessaoDeTeste(t, d, t.TempDir())

	res, err := d.Exec(context.Background(), id,
		[]string{"sh", "-c", "echo saida; echo erro >&2; exit 3"}, ExecOpts{})

	require.NoError(t, err, "comando que falha não é erro de infraestrutura")
	require.Equal(t, 3, res.ExitCode)
	require.Equal(t, "saida\n", res.Stdout)
	require.Equal(t, "erro\n", res.Stderr)
}

func TestIntegracao_ArtefatoPertenceAoUsuarioDoHost(t *testing.T) {
	// É o portão da Fase 0, agora como teste permanente: se alguém mexer nas
	// flags e derrubar o --user, isto reprova antes de o usuário descobrir que
	// não consegue apagar o próprio bin/.
	d := exigeDocker(t)
	ws := t.TempDir()
	id, _ := sessaoDeTeste(t, d, ws)

	res, err := d.Exec(context.Background(), id, []string{"sh", "-c", "echo oi > artefato.txt"}, ExecOpts{})
	require.NoError(t, err)
	require.Equal(t, 0, res.ExitCode, res.Output())

	info, err := os.Stat(filepath.Join(ws, "artefato.txt"))
	require.NoError(t, err, "o arquivo criado no container precisa aparecer no host")

	st := info.Sys().(*syscall.Stat_t)
	require.Equal(t, os.Getuid(), int(st.Uid), "artefato pertence ao root: o usuário não conseguiria editá-lo")
}

func TestIntegracao_Timeout(t *testing.T) {
	d := exigeDocker(t)
	id, _ := sessaoDeTeste(t, d, t.TempDir())

	inicio := time.Now()
	res, err := d.Exec(context.Background(), id,
		[]string{"sleep", "30"}, ExecOpts{Timeout: 500 * time.Millisecond})

	require.NoError(t, err)
	require.True(t, res.TimedOut, "o estouro de tempo precisa ser distinguível de uma reprovação")
	require.Less(t, time.Since(inicio), 10*time.Second, "não pode esperar os 30s")
}

func TestIntegracao_SemRedePorPadrao(t *testing.T) {
	d := exigeDocker(t)
	id, _ := sessaoDeTeste(t, d, t.TempDir())

	res, err := d.Exec(context.Background(), id,
		[]string{"sh", "-c", "wget -q -T 3 -O- https://proxy.golang.org/ 2>&1 || echo SEM_REDE"},
		ExecOpts{Timeout: 20 * time.Second})

	require.NoError(t, err)
	require.Contains(t, res.Stdout, "SEM_REDE", "a sessão precisa nascer isolada")
}

func TestIntegracao_CacheVolumeEhGravavel(t *testing.T) {
	// Um volume nomeado nasce pertencendo ao root. Sem o chown, o container
	// rodando como o usuário do host não conseguiria escrever nele — e todo o
	// ganho de 12s para 400ms por validação iria embora em silêncio.
	d := exigeDocker(t)
	id, lang := sessaoDeTeste(t, d, t.TempDir())

	res, err := d.Exec(context.Background(), id,
		[]string{"sh", "-c", "echo teste > " + lang.CacheDir + "/escrita.txt && cat " + lang.CacheDir + "/escrita.txt"},
		ExecOpts{Timeout: 10 * time.Second})

	require.NoError(t, err)
	require.Equal(t, 0, res.ExitCode, "não conseguiu escrever no volume de cache: %s", res.Output())
	require.Equal(t, "teste\n", res.Stdout)
}

func TestIntegracao_RunEphemeral(t *testing.T) {
	d := exigeDocker(t)
	ws := t.TempDir()
	lang := linguagemDeTeste()
	t.Cleanup(func() { exec.Command("docker", "volume", "rm", lang.CacheVolume()).Run() })

	res, err := d.RunEphemeral(context.Background(),
		SessionSpec{Language: lang, Workspace: ws},
		[]string{"sh", "-c", "pwd"}, ExecOpts{Timeout: 30 * time.Second})

	require.NoError(t, err)
	require.Equal(t, 0, res.ExitCode, res.Output())
	require.Equal(t, "/workspace\n", res.Stdout)
}

func TestIntegracao_CleanupOrphans(t *testing.T) {
	d := exigeDocker(t)
	ctx := context.Background()
	lang := linguagemDeTeste()
	t.Cleanup(func() { exec.Command("docker", "volume", "rm", lang.CacheVolume()).Run() })

	// Simula um CLI que morreu deixando a sessão viva.
	id, err := d.StartSession(ctx, SessionSpec{Language: lang, Workspace: t.TempDir()})
	require.NoError(t, err)

	n, err := d.CleanupOrphans(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, n, 1)

	res, err := d.run(ctx, "ps", "--all", "--quiet", "--filter", "id="+string(id))
	require.NoError(t, err)
	require.Empty(t, strings.TrimSpace(res.Stdout), "o container órfão precisa ter sumido")
}

func TestIntegracao_ImageExists(t *testing.T) {
	d := exigeDocker(t)
	ctx := context.Background()

	existe, err := d.ImageExists(ctx, imagemIntegracao)
	require.NoError(t, err)
	require.True(t, existe)

	existe, err = d.ImageExists(ctx, "langlings/imagem-que-nao-existe:0.0.0")
	require.NoError(t, err)
	require.False(t, existe)
}

func TestIntegracao_StopEhIdempotenteComIDVazio(t *testing.T) {
	d := exigeDocker(t)
	require.NoError(t, d.Stop(context.Background(), ""))
}

func TestJaSumiu(t *testing.T) {
	// Remover um container que já não existe é sucesso: o estado desejado é
	// "não está lá", não "fui eu quem removeu".
	require.True(t, jaSumiu("Error response from daemon: No such container: abc"))
	require.True(t, jaSumiu("removal of container abc is already in progress"))
	require.False(t, jaSumiu("Error response from daemon: permission denied"))
}
