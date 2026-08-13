// Package engine costura catálogo, persistência, container e watcher no loop
// que é o produto: editar → salvar → validar → refletir estado.
package engine

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sync"
	"time"

	"langlings/internal/catalog"
	"langlings/internal/domain"
	"langlings/internal/runner"
	"langlings/internal/store"
	"langlings/internal/watcher"
)

// Engine orquestra uma sessão de uso.
type Engine struct {
	Catalog *catalog.Catalog
	Store   *store.Store
	Runner  runner.Runner
	Paths   Paths
	Watch   watcher.Config

	// Now existe para os testes controlarem o relógio.
	Now func() time.Time

	mu       sync.Mutex
	sessions map[string]runner.SessionID // por slug de linguagem
}

// New devolve um engine pronto para uso.
func New(c *catalog.Catalog, s *store.Store, r runner.Runner, p Paths) *Engine {
	return &Engine{
		Catalog:  c,
		Store:    s,
		Runner:   r,
		Paths:    p,
		Now:      time.Now,
		sessions: map[string]runner.SessionID{},
	}
}

// Bootstrap prepara o ambiente: limpa containers deixados por uma execução
// anterior e alinha o banco com o conteúdo em disco.
func (e *Engine) Bootstrap(ctx context.Context) error {
	if err := e.Paths.EnsureDirs(); err != nil {
		return err
	}

	// Containers órfãos vêm de um CLI morto abruptamente. Limpar no boot é o
	// que evita acumular containers a cada crash.
	if _, err := e.Runner.CleanupOrphans(ctx); err != nil {
		return fmt.Errorf("limpando containers de execuções anteriores: %w", err)
	}

	return e.Store.Reconcile(e.Catalog.Exercises)
}

// LanguageView junta a linguagem, se ela está instalada e o resumo de progresso.
type LanguageView struct {
	Language  domain.Language
	Installed bool
	Summary   domain.Summary
}

// Languages monta a lista da tela de linguagens.
func (e *Engine) Languages(ctx context.Context) ([]LanguageView, error) {
	out := make([]LanguageView, 0, len(e.Catalog.Languages))

	for _, lang := range e.Catalog.Languages {
		// "Instalada" é sempre consultado ao vivo, nunca persistido: um
		// `docker rmi` feito por fora precisa se refletir na hora.
		installed, err := e.Runner.ImageExists(ctx, lang.ImageRef())
		if err != nil {
			return nil, err
		}

		progress, err := e.Store.ByLanguage(lang.Slug)
		if err != nil {
			return nil, err
		}

		out = append(out, LanguageView{
			Language:  lang,
			Installed: installed,
			Summary:   domain.Summarize(progress),
		})
	}
	return out, nil
}

// Install baixa ou constrói a imagem de uma linguagem, transmitindo o log.
func (e *Engine) Install(ctx context.Context, lang domain.Language, progress io.Writer) error {
	return e.Runner.EnsureImage(ctx, lang, e.Paths.ContentRoot, progress)
}

// ExerciseView junta o exercício ao seu progresso, que é o que a lista exibe.
type ExerciseView struct {
	Exercise domain.Exercise
	Progress domain.Progress
}

// Exercises monta a lista de atividades de uma linguagem.
func (e *Engine) Exercises(language string) ([]ExerciseView, error) {
	exercises := e.Catalog.ByLanguage(language)
	out := make([]ExerciseView, 0, len(exercises))

	for _, ex := range exercises {
		p, err := e.Store.Get(ex.Path)
		if err != nil {
			return nil, err
		}
		out = append(out, ExerciseView{Exercise: ex, Progress: p})
	}
	return out, nil
}

// ---------- sessão de container ----------

// ensureSession devolve a sessão viva da linguagem, subindo uma se necessário.
//
// A sessão monta a raiz do workspace, não o diretório de um exercício: assim
// um único container serve todos os exercícios da linguagem, e trocar de
// exercício não paga o custo de criar container.
func (e *Engine) ensureSession(ctx context.Context, lang domain.Language) (runner.SessionID, error) {
	e.mu.Lock()
	id, ok := e.sessions[lang.Slug]
	e.mu.Unlock()
	if ok {
		return id, nil
	}

	id, err := e.Runner.StartSession(ctx, runner.SessionSpec{
		Language:  lang,
		Workspace: e.Paths.Workspace,
	})
	if err != nil {
		return "", err
	}

	e.mu.Lock()
	e.sessions[lang.Slug] = id
	e.mu.Unlock()
	return id, nil
}

// dropSession descarta a sessão de uma linguagem.
//
// É o que fazemos depois de um timeout: o cliente do docker exec morreu, mas o
// processo dentro do container pode ter sobrevivido e continuar consumindo CPU.
// Jogar o container fora é a forma determinística de garantir que a próxima
// validação começa limpa — e o volume de cache torna o recomeço barato.
func (e *Engine) dropSession(ctx context.Context, slug string) {
	e.mu.Lock()
	id, ok := e.sessions[slug]
	delete(e.sessions, slug)
	e.mu.Unlock()

	if ok {
		_ = e.Runner.Stop(ctx, id)
	}
}

// Shutdown encerra todas as sessões. Chamado na saída do CLI.
func (e *Engine) Shutdown(ctx context.Context) {
	e.mu.Lock()
	ids := make([]runner.SessionID, 0, len(e.sessions))
	for _, id := range e.sessions {
		ids = append(ids, id)
	}
	e.sessions = map[string]runner.SessionID{}
	e.mu.Unlock()

	for _, id := range ids {
		_ = e.Runner.Stop(ctx, id)
	}
}

// ShellCommand monta o comando de shell interativo dentro do container da
// linguagem, já posicionado no diretório do exercício.
//
// Devolver o comando em vez de executá-lo é o que permite entregá-lo ao
// tea.ExecProcess, que precisa assumir o TTY.
func (e *Engine) ShellCommand(ctx context.Context, ex domain.Exercise) (*exec.Cmd, error) {
	lang, ok := e.Catalog.Language(ex.Language)
	if !ok {
		return nil, fmt.Errorf("linguagem %q não encontrada", ex.Language)
	}
	if _, err := e.Paths.Materialize(ex); err != nil {
		return nil, err
	}

	id, err := e.ensureSession(ctx, lang)
	if err != nil {
		return nil, err
	}

	// O shell abre já dentro do exercício: o usuário não deveria precisar
	// descobrir para onde navegar.
	langWithWorkdir := lang
	langWithWorkdir.Workdir = containerDir(lang, ex)

	return e.Runner.InteractiveCmd(id, langWithWorkdir), nil
}

// containerDir é o diretório do exercício visto de dentro do container.
func containerDir(lang domain.Language, ex domain.Exercise) string {
	return path.Join(lang.Workdir, ex.Path)
}

// ---------- validação ----------

// CriterionResult é o veredito sobre um critério, para exibição.
type CriterionResult struct {
	Label   string
	Outcome domain.Outcome
}

// Result é o desfecho de uma validação.
type Result struct {
	Exercise domain.Exercise
	Passed   bool

	// Output é a saída bruta do container. É o que o usuário realmente quer
	// ler quando algo falha, então não é resumida nem interpretada.
	Output string

	// Criteria é preenchido no modo criteria, um item por critério.
	Criteria []CriterionResult

	TimedOut bool
	Hash     string

	// Progress é o estado depois de aplicada a transição.
	Progress domain.Progress

	// Err sinaliza falha de infraestrutura, não reprovação do exercício.
	Err error
}

// Validate executa a validação do exercício sobre o workspace do usuário e
// persiste o resultado.
func (e *Engine) Validate(ctx context.Context, ex domain.Exercise) Result {
	dir, err := e.Paths.Materialize(ex)
	if err != nil {
		return Result{Exercise: ex, Err: err}
	}

	files, err := ReadEditable(dir, ex)
	if err != nil {
		return Result{Exercise: ex, Err: err}
	}
	hash := domain.ContentHash(files)

	res := e.runValidation(ctx, ex, e.Paths.Workspace, ex.Path)
	res.Hash = hash

	if res.Err != nil {
		return res
	}

	current, err := e.Store.Get(ex.Path)
	if err != nil {
		res.Err = err
		return res
	}
	current.Language = ex.Language
	current.Category = ex.Category

	next := domain.Apply(current, res.Passed, hash, e.Now())
	if err := e.Store.Save(next); err != nil {
		res.Err = err
		return res
	}
	res.Progress = next
	return res
}

// runValidation roda a validação sobre um workspace arbitrário.
//
// workspaceRoot é o diretório do host montado no container e subPath é o
// caminho do exercício dentro dele. Verify usa um workspace temporário; o loop
// de save usa o do usuário.
func (e *Engine) runValidation(ctx context.Context, ex domain.Exercise, workspaceRoot, subPath string) Result {
	lang, ok := e.Catalog.Language(ex.Language)
	if !ok {
		return Result{Exercise: ex, Err: fmt.Errorf("linguagem %q não encontrada", ex.Language)}
	}

	hostDir := filepath.Join(workspaceRoot, filepath.FromSlash(subPath))
	workdir := path.Join(lang.Workdir, subPath)

	run := e.executor(ctx, lang, workspaceRoot, workdir, ex.Validation)

	switch ex.Validation.Mode {
	case domain.ModeTest:
		return e.validateByTest(ex, run)
	case domain.ModeCriteria:
		return e.validateByCriteria(ex, run, hostDir)
	default:
		return Result{Exercise: ex, Err: fmt.Errorf("modo de validação %q desconhecido", ex.Validation.Mode)}
	}
}

// executorFunc roda um comando dentro do container já configurado.
type executorFunc func(cmd []string) (runner.ExecResult, error)

// executor escolhe entre a sessão viva e um container efêmero.
//
// Dois motivos levam ao caminho efêmero:
//
//   - O exercício declara rede. Limites de rede são fixados na criação do
//     container e a sessão nasce isolada, então ela não tem como servi-lo.
//   - O workspace não é o do usuário. É o caso do verify, que roda sobre um
//     diretório temporário para não encostar no trabalho de ninguém — e a
//     sessão monta um workspace só, escolhido quando subiu.
//
// Nos dois casos pagamos a criação de um container por validação; o volume de
// cache mantém isso na casa de centenas de milissegundos.
func (e *Engine) executor(ctx context.Context, lang domain.Language, workspaceRoot, workdir string, spec domain.ValidationSpec) executorFunc {
	opts := runner.ExecOpts{Workdir: workdir, Timeout: spec.Timeout}

	if spec.Network || workspaceRoot != e.Paths.Workspace {
		return func(cmd []string) (runner.ExecResult, error) {
			return e.Runner.RunEphemeral(ctx, runner.SessionSpec{
				Language:  lang,
				Workspace: workspaceRoot,
				Network:   spec.Network,
			}, cmd, opts)
		}
	}

	return func(cmd []string) (runner.ExecResult, error) {
		id, err := e.ensureSession(ctx, lang)
		if err != nil {
			return runner.ExecResult{}, err
		}

		res, err := e.Runner.Exec(ctx, id, cmd, opts)
		if err == nil && res.TimedOut {
			// O processo lá dentro pode ter sobrevivido ao cliente que o
			// invocou; a sessão deixa de ser confiável.
			e.dropSession(ctx, lang.Slug)
		}
		return res, err
	}
}

func (e *Engine) validateByTest(ex domain.Exercise, run executorFunc) Result {
	res, err := run(ex.Validation.Command)
	if err != nil {
		return Result{Exercise: ex, Err: err}
	}

	return Result{
		Exercise: ex,
		Passed:   res.ExitCode == 0 && !res.TimedOut,
		Output:   res.Output(),
		TimedOut: res.TimedOut,
	}
}

func (e *Engine) validateByCriteria(ex domain.Exercise, run executorFunc, hostDir string) Result {
	out := Result{Exercise: ex, Passed: true}

	for _, criterion := range ex.Validation.Criteria {
		probe := criterion.Probe()

		var probeResult domain.ProbeResult
		switch probe.Kind {
		case domain.ProbeStat:
			// O workspace é montado no container, então o artefato gerado lá
			// dentro já está visível aqui. Statar no host é equivalente e
			// dispensa um round-trip de container.
			_, err := os.Stat(filepath.Join(hostDir, filepath.FromSlash(probe.Path)))
			probeResult = domain.ProbeResult{Exists: err == nil}

		case domain.ProbeRun:
			res, err := run(probe.Command)
			if err != nil {
				return Result{Exercise: ex, Err: err}
			}
			probeResult = domain.ProbeResult{
				ExitCode: res.ExitCode,
				Stdout:   res.Stdout,
				Stderr:   res.Stderr,
			}
			if res.TimedOut {
				probeResult.Err = fmt.Errorf("o comando excedeu %s", ex.Validation.Timeout)
				out.TimedOut = true
			}
			out.Output = appendOutput(out.Output, res.Output())

		default:
			probeResult = domain.ProbeResult{Err: fmt.Errorf("probe %q desconhecida", probe.Kind)}
		}

		outcome := criterion.Evaluate(probeResult)
		if !outcome.Passed {
			out.Passed = false
		}
		out.Criteria = append(out.Criteria, CriterionResult{Label: criterion.Label(), Outcome: outcome})
	}

	return out
}

func appendOutput(acc, add string) string {
	switch {
	case add == "":
		return acc
	case acc == "":
		return add
	default:
		return acc + "\n" + add
	}
}

// ---------- loop de save ----------

// UpdateKind distingue as duas coisas que a TUI precisa saber.
type UpdateKind int

const (
	// UpdateValidating é emitido no instante em que um save vira trabalho.
	// Existe para a tela poder mostrar o spinner antes de haver resultado:
	// três segundos de silêncio parecem travamento, três segundos de spinner
	// parecem trabalho.
	UpdateValidating UpdateKind = iota

	// UpdateResult carrega o desfecho.
	UpdateResult
)

// Update é o que o engine empurra para quem observa um exercício. Um canal só,
// em ordem, garante que o "validando" sempre chega antes do seu resultado.
type Update struct {
	Kind   UpdateKind
	Result Result
}

// Session é a observação de um exercício aberto.
type Session struct {
	Dir     string // diretório no host que o usuário deve abrir no editor
	Updates <-chan Update
}

// Open materializa o exercício, começa a observar os arquivos editáveis e
// valida a cada mudança real de conteúdo.
//
// O canal é fechado quando ctx termina.
func (e *Engine) Open(ctx context.Context, ex domain.Exercise) (*Session, error) {
	dir, err := e.Paths.Materialize(ex)
	if err != nil {
		return nil, err
	}

	events, err := watcher.Watch(ctx, dir, ex.Editable, e.Watch)
	if err != nil {
		return nil, err
	}

	updates := make(chan Update)
	go e.loop(ctx, ex, dir, events, updates)

	return &Session{Dir: dir, Updates: updates}, nil
}

func (e *Engine) loop(ctx context.Context, ex domain.Exercise, dir string, events <-chan watcher.Event, updates chan<- Update) {
	defer close(updates)

	// cancelAtual interrompe a validação em andamento quando chega um save
	// novo: o resultado da anterior seria sobre código que o usuário já
	// substituiu. Last-write-wins.
	var cancelAtual context.CancelFunc
	defer func() {
		if cancelAtual != nil {
			cancelAtual()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case _, ok := <-events:
			if !ok {
				return
			}

			files, err := ReadEditable(dir, ex)
			if err != nil {
				emit(ctx, updates, Update{Kind: UpdateResult, Result: Result{Exercise: ex, Err: err}})
				continue
			}
			hash := domain.ContentHash(files)

			// O árbitro entre fsnotify e polling: se o conteúdo é o mesmo da
			// última validação, não há o que validar.
			current, err := e.Store.Get(ex.Path)
			if err != nil {
				emit(ctx, updates, Update{Kind: UpdateResult, Result: Result{Exercise: ex, Err: err}})
				continue
			}
			if current.LastContentHash != "" && current.LastContentHash == hash {
				continue
			}

			if cancelAtual != nil {
				cancelAtual()
			}
			validCtx, cancel := context.WithCancel(ctx)
			cancelAtual = cancel

			emit(ctx, updates, Update{Kind: UpdateValidating})

			res := e.Validate(validCtx, ex)
			if validCtx.Err() != nil {
				// Foi cancelada por um save mais novo: descartar em silêncio.
				continue
			}
			emit(ctx, updates, Update{Kind: UpdateResult, Result: res})
		}
	}
}

func emit(ctx context.Context, ch chan<- Update, u Update) {
	select {
	case ch <- u:
	case <-ctx.Done():
	}
}
