// Package tui é a interface em Bubble Tea.
//
// O Update é uma função pura (Model, Msg) → (Model, Cmd): toda a navegação e
// toda a transição de tela são testáveis construindo um Model, mandando uma
// tea.KeyMsg e olhando o Model que volta. Sem TTY, sem goroutine, sem espera.
package tui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/domain"
	"langlings/internal/engine"
)

type tela int

const (
	telaBoasVindas tela = iota
	telaLinguagens
	telaExercicios
	telaExercicio
	telaInstalando
)

// Model é o estado inteiro da interface.
type Model struct {
	motor *engine.Engine
	ctx   context.Context

	// program permite que a instalação empurre linhas de log a partir de outra
	// goroutine. É preenchido depois da criação, por Attach.
	program *tea.Program

	tela    tela
	largura int
	altura  int
	Sair    bool

	// Tela de linguagens.
	linguagens   []engine.LanguageView
	cursorIdioma int

	// Tela de exercícios.
	linguagemAtual  domain.Language
	exercicios      []engine.ExerciseView
	cursorExercicio int

	// Tela do exercício.
	atual          engine.ExerciseView
	sessao         *engine.Session
	cancelarSessao context.CancelFunc
	validando      bool
	ultimo         *engine.Result
	dicasAbertas   int

	// Instalação.
	instalando     bool
	logInstalacao  []string
	alvoInstalacao domain.Language

	spinner spinner.Model
	aviso   string
	erro    error
}

// New monta o modelo inicial.
func New(ctx context.Context, motor *engine.Engine) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return Model{
		motor:   motor,
		ctx:     ctx,
		tela:    telaBoasVindas,
		spinner: s,
	}
}

// Attach registra o programa, necessário para o log de instalação em streaming.
func (m *Model) Attach(p *tea.Program) { m.program = p }

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, carregarLinguagens(m.ctx, m.motor))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.largura, m.altura = msg.Width, msg.Height
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		return m.tratarTecla(msg)

	case linguagensCarregadasMsg:
		if msg.err != nil {
			m.erro = msg.err
			return m, nil
		}
		m.linguagens = msg.views
		m.erro = nil
		return m, nil

	case exerciciosCarregadosMsg:
		if msg.err != nil {
			m.erro = msg.err
			return m, nil
		}
		m.exercicios = msg.views
		m.cursorExercicio = primeiroNaoConcluido(msg.views)
		return m, nil

	case exercicioAbertoMsg:
		if msg.err != nil {
			m.erro = msg.err
			m.tela = telaExercicios
			return m, nil
		}
		m.sessao = msg.session
		m.cancelarSessao = msg.cancel
		return m, esperarAtualizacao(msg.session)

	case atualizacaoMsg:
		return m.tratarAtualizacao(msg)

	case validacaoManualMsg:
		m.validando = false
		res := msg.resultado
		m.ultimo = &res
		if res.Err != nil {
			m.erro = res.Err
			return m, nil
		}
		m.atual.Progress = res.Progress
		return m, carregarExercicios(m.motor, m.linguagemAtual.Slug)

	case instalacaoLinhaMsg:
		m.logInstalacao = append(m.logInstalacao, msg.linha)
		// Só as últimas linhas interessam; o resto rolaria para fora da tela.
		if len(m.logInstalacao) > 200 {
			m.logInstalacao = m.logInstalacao[len(m.logInstalacao)-200:]
		}
		return m, nil

	case instalacaoPromptaMsg:
		m.instalando = false
		m.tela = telaLinguagens
		if msg.err != nil {
			m.erro = msg.err
			return m, nil
		}
		m.aviso = fmt.Sprintf("%s instalada.", m.alvoInstalacao.Name)
		return m, carregarLinguagens(m.ctx, m.motor)

	case shellEncerradoMsg:
		if msg.err != nil {
			m.erro = msg.err
			return m, nil
		}
		// Sair do shell é o gatilho da categoria Compilador: os critérios
		// olham artefatos que nenhum arquivo editável denunciou.
		m.validando = true
		return m, validarAgora(m.ctx, m.motor, m.atual.Exercise)

	case resetFeitoMsg:
		if msg.err != nil {
			m.erro = msg.err
			return m, nil
		}
		m.aviso = msg.mensagem
		m.ultimo = nil
		return m, carregarExercicios(m.motor, m.linguagemAtual.Slug)

	case erroMsg:
		m.erro = msg.err
		return m, nil
	}

	return m, nil
}

func (m Model) tratarAtualizacao(msg atualizacaoMsg) (tea.Model, tea.Cmd) {
	if msg.fim {
		m.sessao = nil
		return m, nil
	}

	switch msg.update.Kind {
	case engine.UpdateValidating:
		// O spinner aparece no instante do save, antes de existir resultado.
		m.validando = true
		m.aviso = ""

	case engine.UpdateResult:
		m.validando = false
		res := msg.update.Result
		m.ultimo = &res
		if res.Err != nil {
			m.erro = res.Err
		} else {
			m.erro = nil
			m.atual.Progress = res.Progress
		}
	}

	cmds := []tea.Cmd{esperarAtualizacao(m.sessao)}
	if msg.update.Kind == engine.UpdateResult && msg.update.Result.Err == nil {
		cmds = append(cmds, carregarExercicios(m.motor, m.linguagemAtual.Slug))
	}
	return m, tea.Batch(cmds...)
}

func (m Model) tratarTecla(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Ctrl+C sempre sai, em qualquer tela.
	if msg.Type == tea.KeyCtrlC {
		return m.encerrar()
	}

	switch m.tela {
	case telaBoasVindas:
		return m.teclaBoasVindas(msg)
	case telaLinguagens:
		return m.teclaLinguagens(msg)
	case telaExercicios:
		return m.teclaExercicios(msg)
	case telaExercicio:
		return m.teclaExercicio(msg)
	case telaInstalando:
		return m, nil
	}
	return m, nil
}

func (m Model) encerrar() (tea.Model, tea.Cmd) {
	m.fecharSessao()
	m.Sair = true
	return m, tea.Quit
}

// fecharSessao cancela o watcher do exercício aberto.
func (m *Model) fecharSessao() {
	if m.cancelarSessao != nil {
		m.cancelarSessao()
		m.cancelarSessao = nil
	}
	m.sessao = nil
	m.validando = false
}

func (m Model) View() string {
	switch m.tela {
	case telaBoasVindas:
		return m.viewBoasVindas()
	case telaLinguagens:
		return m.viewLinguagens()
	case telaExercicios:
		return m.viewExercicios()
	case telaExercicio:
		return m.viewExercicio()
	case telaInstalando:
		return m.viewInstalando()
	}
	return ""
}

// primeiroNaoConcluido posiciona o cursor no próximo exercício a fazer.
//
// A progressão é livre: nada trava, mas o cursor sugere por onde seguir.
func primeiroNaoConcluido(views []engine.ExerciseView) int {
	for i, v := range views {
		if v.Progress.Status != domain.StatusCompleted {
			return i
		}
	}
	return 0
}

// mover desloca um cursor mantendo-o dentro dos limites.
func mover(cursor, delta, total int) int {
	if total == 0 {
		return 0
	}
	cursor += delta
	if cursor < 0 {
		return 0
	}
	if cursor >= total {
		return total - 1
	}
	return cursor
}

// ehParaCima e ehParaBaixo aceitam setas e as teclas do vim.
func ehParaCima(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyUp || msg.String() == "k"
}

func ehParaBaixo(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyDown || msg.String() == "j"
}

func ehVoltar(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyEsc || msg.String() == "q"
}

func ehConfirmar(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyEnter || msg.String() == " "
}
