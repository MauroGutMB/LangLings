package tui

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/domain"
	"langlings/internal/engine"
)

// Mensagens que chegam de fora do laço de teclado.
type (
	linguagensCarregadasMsg struct {
		views []engine.LanguageView
		err   error
	}

	exerciciosCarregadosMsg struct {
		views []engine.ExerciseView
		err   error
	}

	exercicioAbertoMsg struct {
		session *engine.Session
		cancel  context.CancelFunc
		err     error
	}

	// atualizacaoMsg carrega um Update do engine. É o que traz tanto o aviso
	// de "validando" quanto o resultado.
	atualizacaoMsg struct {
		update engine.Update
		fim    bool // o canal fechou
	}

	instalacaoLinhaMsg   struct{ linha string }
	instalacaoPromptaMsg struct{ err error }

	shellEncerradoMsg struct{ err error }

	validacaoManualMsg struct{ resultado engine.Result }

	resetFeitoMsg struct {
		mensagem string
		err      error
	}

	erroMsg struct{ err error }
)

// carregarLinguagens consulta o estado das linguagens, inclusive se a imagem
// existe no Docker agora.
func carregarLinguagens(ctx context.Context, e *engine.Engine) tea.Cmd {
	return func() tea.Msg {
		views, err := e.Languages(ctx)
		return linguagensCarregadasMsg{views: views, err: err}
	}
}

func carregarExercicios(e *engine.Engine, linguagem string) tea.Cmd {
	return func() tea.Msg {
		views, err := e.Exercises(linguagem)
		return exerciciosCarregadosMsg{views: views, err: err}
	}
}

// abrirExercicio materializa o exercício e começa a observar os arquivos.
//
// O contexto é filho do contexto do programa e vive enquanto a tela do
// exercício estiver aberta; sair dela cancela o watcher.
func abrirExercicio(pai context.Context, e *engine.Engine, ex domain.Exercise) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithCancel(pai)
		sess, err := e.Open(ctx, ex)
		if err != nil {
			cancel()
			return exercicioAbertoMsg{err: err}
		}
		return exercicioAbertoMsg{session: sess, cancel: cancel}
	}
}

// esperarAtualizacao bloqueia numa leitura do canal do engine e se reagenda a
// cada mensagem. É o padrão do Bubble Tea para consumir um canal sem precisar
// de program.Send a partir de outra goroutine.
func esperarAtualizacao(s *engine.Session) tea.Cmd {
	return func() tea.Msg {
		up, ok := <-s.Updates
		if !ok {
			return atualizacaoMsg{fim: true}
		}
		return atualizacaoMsg{update: up}
	}
}

// instalar baixa ou constrói a imagem, transmitindo o log linha a linha.
func instalar(ctx context.Context, e *engine.Engine, lang domain.Language, p *tea.Program) tea.Cmd {
	return func() tea.Msg {
		w := &escritorDeLinhas{enviar: func(linha string) {
			if p != nil {
				p.Send(instalacaoLinhaMsg{linha: linha})
			}
		}}
		err := e.Install(ctx, lang, w)
		w.descarregar()
		return instalacaoPromptaMsg{err: err}
	}
}

// validarAgora dispara uma validação fora do watcher. É o que roda quando o
// usuário sai do shell, na categoria Compilador: o critério observa artefatos
// que nenhum arquivo editável denunciou.
func validarAgora(ctx context.Context, e *engine.Engine, ex domain.Exercise) tea.Cmd {
	return func() tea.Msg {
		return validacaoManualMsg{resultado: e.Validate(ctx, ex)}
	}
}

// escritorDeLinhas transforma a saída do docker em mensagens de uma linha.
type escritorDeLinhas struct {
	enviar func(string)
	buffer strings.Builder
}

func (w *escritorDeLinhas) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			w.enviar(w.buffer.String())
			w.buffer.Reset()
			continue
		}
		if b != '\r' {
			w.buffer.WriteByte(b)
		}
	}
	return len(p), nil
}

func (w *escritorDeLinhas) descarregar() {
	if w.buffer.Len() > 0 {
		w.enviar(w.buffer.String())
		w.buffer.Reset()
	}
}
