package runner

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"langlings/internal/domain"
)

// Docker fala com o daemon executando o binário `docker`.
//
// A escolha do CLI em vez do SDK oficial é deliberada: todo comando montado
// aqui pode ser copiado e colado num terminal para reproduzir exatamente o que
// o LangLings fez. O preço é interpretar exit codes e texto em vez de structs,
// e ele é pago em troca de uma árvore de dependências que cabe na cabeça.
type Docker struct {
	// Binary é o executável do docker. Existe para os testes, não para o uso.
	Binary string

	// UID e GID são a identidade usada dentro do container. Sem isso, todo
	// artefato gerado lá dentro aparece no host pertencendo ao root e o
	// usuário não consegue nem editar nem apagar o próprio trabalho.
	UID int
	GID int
}

// NewDocker devolve um runner configurado com a identidade do usuário atual.
func NewDocker() *Docker {
	return &Docker{Binary: "docker", UID: os.Getuid(), GID: os.Getgid()}
}

var _ Runner = (*Docker)(nil)

func (d *Docker) bin() string {
	if d.Binary == "" {
		return "docker"
	}
	return d.Binary
}

// run executa o docker e captura as saídas separadamente.
func (d *Docker) run(ctx context.Context, args ...string) (ExecResult, error) {
	cmd := exec.CommandContext(ctx, d.bin(), args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	res := ExecResult{Stdout: stdout.String(), Stderr: stderr.String()}

	var exitErr *exec.ExitError
	switch {
	case err == nil:
		res.ExitCode = 0
	case errors.As(err, &exitErr):
		res.ExitCode = exitErr.ExitCode()
	default:
		// Falha em invocar o próprio docker (não encontrado, sem permissão).
		return res, fmt.Errorf("executando %s: %w", d.bin(), err)
	}

	if ctx.Err() != nil {
		res.TimedOut = errors.Is(ctx.Err(), context.DeadlineExceeded)
	}
	return res, nil
}

func (d *Docker) Available(ctx context.Context) error {
	res, err := d.run(ctx, "version", "--format", "{{.Server.Version}}")
	if err != nil {
		return fmt.Errorf("docker não encontrado no PATH: %w", err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("daemon do Docker não respondeu: %s", strings.TrimSpace(res.Output()))
	}
	return nil
}

func (d *Docker) ImageExists(ctx context.Context, ref string) (bool, error) {
	res, err := d.run(ctx, "image", "inspect", ref)
	if err != nil {
		return false, err
	}
	return res.ExitCode == 0, nil
}

// EnsureImage baixa a imagem oficial ou constrói a da linguagem.
func (d *Docker) EnsureImage(ctx context.Context, lang domain.Language, contextDir string, progress io.Writer) error {
	ref := lang.ImageRef()

	exists, err := d.ImageExists(ctx, ref)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	var args []string
	if lang.BuildsOwnImage() {
		args = []string{"build", "-t", ref, "-f", lang.Dockerfile, contextDir}
	} else {
		args = []string{"pull", ref}
	}

	if progress == nil {
		progress = io.Discard
	}
	cmd := exec.CommandContext(ctx, d.bin(), args...)
	cmd.Stdout = progress
	cmd.Stderr = progress

	if err := cmd.Run(); err != nil {
		verbo := "baixar"
		if lang.BuildsOwnImage() {
			verbo = "construir"
		}
		return fmt.Errorf("não foi possível %s a imagem %s: %w", verbo, ref, err)
	}
	return nil
}

// containerArgs monta as flags comuns a sessão e container efêmero.
//
// São as mesmas flags validadas na Fase 0; qualquer mudança aqui merece
// reexecutar aquele teste de ponta a ponta.
func (d *Docker) containerArgs(spec SessionSpec) []string {
	lang := spec.Language

	args := []string{
		"--user", fmt.Sprintf("%d:%d", d.UID, d.GID),
		"--label", ManagedLabelTrue,
		"--label", LanguageLabel + "=" + lang.Slug,
		"--memory", DefaultMemory,
		"--cpus", DefaultCPUs,
		"--volume", spec.Workspace + ":" + lang.Workdir,
		"--workdir", lang.Workdir,
	}

	if spec.Network {
		args = append(args, "--network", "bridge")
	} else {
		args = append(args, "--network", "none")
	}

	// HOME precisa ser gravável: com UID arbitrário, o $HOME padrão da imagem
	// não é, e ferramentas falham com erros que não parecem ser de permissão.
	env := map[string]string{"HOME": "/tmp"}
	for k, v := range lang.Env {
		env[k] = v
	}
	for _, k := range sortedKeys(env) {
		args = append(args, "--env", k+"="+env[k])
	}

	if lang.CacheDir != "" {
		args = append(args, "--volume", lang.CacheVolume()+":"+lang.CacheDir)
	}

	return args
}

// ensureCacheVolume cria o volume de cache da linguagem, se necessário.
//
// O chown é obrigatório: um volume nomeado nasce pertencendo ao root, e o
// container roda como o usuário do host, que então não consegue escrever nele.
func (d *Docker) ensureCacheVolume(ctx context.Context, lang domain.Language) error {
	if lang.CacheDir == "" {
		return nil
	}
	vol := lang.CacheVolume()

	if res, err := d.run(ctx, "volume", "inspect", vol); err != nil {
		return err
	} else if res.ExitCode == 0 {
		return nil
	}

	if res, err := d.run(ctx, "volume", "create", vol); err != nil {
		return err
	} else if res.ExitCode != 0 {
		return fmt.Errorf("criando volume %s: %s", vol, strings.TrimSpace(res.Output()))
	}

	res, err := d.run(ctx,
		"run", "--rm", "--user", "0:0", "--network", "none",
		"--label", ManagedLabelTrue,
		"--volume", vol+":"+lang.CacheDir,
		lang.ImageRef(),
		"chown", "-R", fmt.Sprintf("%d:%d", d.UID, d.GID), lang.CacheDir,
	)
	if err != nil {
		return err
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("ajustando dono do volume %s: %s", vol, strings.TrimSpace(res.Output()))
	}
	return nil
}

func (d *Docker) StartSession(ctx context.Context, spec SessionSpec) (SessionID, error) {
	if err := d.ensureCacheVolume(ctx, spec.Language); err != nil {
		return "", err
	}

	args := append([]string{"run", "--detach"}, d.containerArgs(spec)...)
	args = append(args, spec.Language.ImageRef())
	// `tail -f /dev/null` em vez de `sleep infinity`: nem toda build do busybox
	// aceita "infinity" como argumento, e tail existe em qualquer imagem.
	args = append(args, "tail", "-f", "/dev/null")

	res, err := d.run(ctx, args...)
	if err != nil {
		return "", err
	}
	if res.ExitCode != 0 {
		return "", fmt.Errorf("não foi possível subir o container de %s: %s",
			spec.Language.Name, strings.TrimSpace(res.Output()))
	}

	id := strings.TrimSpace(res.Stdout)
	if id == "" {
		return "", fmt.Errorf("docker run não devolveu um id de container")
	}
	return SessionID(id), nil
}

func (d *Docker) Exec(ctx context.Context, id SessionID, cmd []string, opts ExecOpts) (ExecResult, error) {
	if len(cmd) == 0 {
		return ExecResult{}, fmt.Errorf("comando vazio")
	}

	ctx, cancel := withTimeout(ctx, opts.Timeout)
	defer cancel()

	args := []string{"exec"}
	if opts.Workdir != "" {
		args = append(args, "--workdir", opts.Workdir)
	}
	args = append(args, string(id))
	args = append(args, cmd...)

	return d.run(ctx, args...)
}

func (d *Docker) RunEphemeral(ctx context.Context, spec SessionSpec, cmd []string, opts ExecOpts) (ExecResult, error) {
	if len(cmd) == 0 {
		return ExecResult{}, fmt.Errorf("comando vazio")
	}
	if err := d.ensureCacheVolume(ctx, spec.Language); err != nil {
		return ExecResult{}, err
	}

	ctx, cancel := withTimeout(ctx, opts.Timeout)
	defer cancel()

	args := append([]string{"run", "--rm"}, d.containerArgs(spec)...)
	if opts.Workdir != "" {
		args = append(args, "--workdir", opts.Workdir)
	}
	args = append(args, spec.Language.ImageRef())
	args = append(args, cmd...)

	return d.run(ctx, args...)
}

func (d *Docker) InteractiveCmd(id SessionID, lang domain.Language) *exec.Cmd {
	args := []string{"exec", "--interactive", "--tty", "--workdir", lang.Workdir, string(id)}
	args = append(args, lang.Shell...)
	return exec.Command(d.bin(), args...)
}

// Stop remove um container de sessão.
//
// É idempotente de propósito: o que interessa é o estado final (o container
// não existe), não quem chegou lá primeiro. Um container já removido, ou em
// remoção por outro caminho — a limpeza de órfãos no boot, por exemplo — não
// é falha nenhuma.
func (d *Docker) Stop(ctx context.Context, id SessionID) error {
	if id == "" {
		return nil
	}
	res, err := d.run(ctx, "rm", "--force", string(id))
	if err != nil {
		return err
	}
	if res.ExitCode != 0 && !jaSumiu(res.Output()) {
		return fmt.Errorf("removendo container %s: %s", id, strings.TrimSpace(res.Output()))
	}
	return nil
}

// jaSumiu reconhece as respostas do daemon que significam "o container não
// está mais lá", que para o LangLings é sucesso.
func jaSumiu(output string) bool {
	lower := strings.ToLower(output)
	for _, marca := range []string{"no such container", "already in progress", "is already being removed"} {
		if strings.Contains(lower, marca) {
			return true
		}
	}
	return false
}

func (d *Docker) CleanupOrphans(ctx context.Context) (int, error) {
	res, err := d.run(ctx, "ps", "--all", "--quiet", "--filter", "label="+ManagedLabelTrue)
	if err != nil {
		return 0, err
	}
	if res.ExitCode != 0 {
		return 0, fmt.Errorf("listando containers órfãos: %s", strings.TrimSpace(res.Output()))
	}

	ids := strings.Fields(res.Stdout)
	if len(ids) == 0 {
		return 0, nil
	}

	args := append([]string{"rm", "--force"}, ids...)
	if _, err := d.run(ctx, args...); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func withTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	if d <= 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, d)
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// Ordem estável mantém os comandos reprodutíveis e os testes determinísticos.
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
	return keys
}
