package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/domain"
	"langlings/internal/engine"
)

// linhasDeSaida limita quanto da saída bruta cabe na tela.
const linhasDeSaida = 14

func (m Model) teclaExercicio(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case ehVoltar(msg):
		m.fecharSessao()
		m.tela = telaExercicios
		m.ultimo = nil
		return m, carregarExercicios(m.motor, m.linguagemAtual.Slug)

	case msg.String() == "h":
		if m.dicasAbertas < len(m.atual.Exercise.Hints) {
			m.dicasAbertas++
		}
		return m, nil

	case msg.String() == "s":
		// Suspende a TUI e entrega o TTY ao shell dentro do container. Ao
		// sair, shellEncerradoMsg dispara a checagem dos critérios.
		cmd, err := m.motor.ShellCommand(m.ctx, m.atual.Exercise)
		if err != nil {
			m.erro = err
			return m, nil
		}
		return m, tea.ExecProcess(cmd, func(err error) tea.Msg {
			return shellEncerradoMsg{err: err}
		})

	case msg.String() == "v":
		// Revalidação manual: útil quando o critério observa algo que nenhum
		// arquivo editável denuncia.
		m.validando = true
		return m, validarAgora(m.ctx, m.motor, m.atual.Exercise)

	case msg.String() == "R":
		ex := m.atual.Exercise
		return m, func() tea.Msg {
			err := m.motor.ResetExercise(ex, engine.ResetCompleto)
			return resetFeitoMsg{mensagem: "Exercício restaurado ao código-base.", err: err}
		}
	}

	return m, nil
}

func (m Model) viewExercicio() string {
	ex := m.atual.Exercise
	var b strings.Builder

	b.WriteString(estiloTitulo.Render(ex.Title) + "\n")
	b.WriteString(estiloSubtitulo.Render(fmt.Sprintf("%s · %s · %s",
		m.linguagemAtual.Name, ex.Category.Label(), m.atual.Progress.Status.Icon())) + "\n\n")

	b.WriteString(estiloCaixa.Render(ex.Objective) + "\n\n")

	// O caminho no host é a informação mais importante da tela: é o que o
	// usuário precisa abrir no editor dele.
	if m.sessao != nil {
		b.WriteString(estiloApagado.Render("  edite este arquivo:") + "\n")
		b.WriteString("  " + estiloCaminho.Render(m.sessao.Dir) + "\n")
		for _, f := range ex.Editable {
			b.WriteString("  " + estiloApagado.Render("└ "+f) + "\n")
		}
		b.WriteString("\n")
	} else {
		b.WriteString("  " + m.spinner.View() + estiloApagado.Render(" preparando o exercício…") + "\n\n")
	}

	if ex.Validation.Mode == domain.ModeCriteria {
		b.WriteString(estiloApagado.Render("  esta atividade é verificada por critérios:") + "\n")
		for _, c := range ex.Validation.Criteria {
			b.WriteString(estiloApagado.Render("  · "+c.Label()) + "\n")
		}
		b.WriteString(estiloApagado.Render("  abra o shell com [s], compile à mão e saia com exit") + "\n\n")
	}

	b.WriteString(m.viewEstadoValidacao())

	for i := 0; i < m.dicasAbertas && i < len(ex.Hints); i++ {
		b.WriteString("\n" + estiloAtencao.Render(fmt.Sprintf("  dica %d: ", i+1)) + ex.Hints[i] + "\n")
	}
	if restantes := len(ex.Hints) - m.dicasAbertas; restantes > 0 {
		b.WriteString("\n" + estiloApagado.Render(fmt.Sprintf("  %d dica(s) disponível(is) — [h]", restantes)) + "\n")
	}

	if m.aviso != "" {
		b.WriteString("\n" + estiloSucesso.Render("  "+m.aviso) + "\n")
	}
	if m.erro != nil {
		b.WriteString("\n" + estiloErro.Render("  erro: "+m.erro.Error()) + "\n")
	}

	b.WriteString(estiloRodape.Render("\nsalve o arquivo para validar · [s] shell · [v] validar · [h] dica · [R] restaurar · esc voltar"))
	return b.String()
}

func (m Model) viewEstadoValidacao() string {
	if m.validando {
		return "  " + m.spinner.View() + estiloApagado.Render(" validando…") + "\n"
	}

	if m.ultimo == nil {
		return estiloApagado.Render("  aguardando você salvar o arquivo…") + "\n"
	}

	var b strings.Builder
	res := *m.ultimo

	switch {
	case res.TimedOut:
		b.WriteString(estiloErro.Render(fmt.Sprintf("  ✗ excedeu %s", m.atual.Exercise.Validation.Timeout)) + "\n")
	case res.Passed:
		b.WriteString(estiloSucesso.Render("  ✓ passou") + "\n")
	default:
		b.WriteString(estiloErro.Render("  ✗ ainda não") + "\n")
	}

	for _, c := range res.Criteria {
		marca := estiloErro.Render("✗")
		if c.Outcome.Passed {
			marca = estiloSucesso.Render("✓")
		}
		b.WriteString(fmt.Sprintf("    %s %s — %s\n", marca, c.Label, estiloApagado.Render(c.Outcome.Detail)))
	}

	if m.atual.Progress.Regressed() {
		b.WriteString(estiloAtencao.Render("  este exercício já foi concluído; o progresso continua salvo") + "\n")
	}

	if saida := strings.TrimSpace(res.Output); saida != "" && !res.Passed {
		b.WriteString("\n" + estiloApagado.Render(ultimasLinhas(saida, linhasDeSaida)) + "\n")
	}

	return b.String()
}

// ultimasLinhas devolve o final de um texto. A saída de um test runner põe o
// que interessa no fim, e o começo costuma ser ruído de compilação.
func ultimasLinhas(s string, n int) string {
	linhas := strings.Split(s, "\n")
	if len(linhas) > n {
		linhas = linhas[len(linhas)-n:]
	}
	for i, l := range linhas {
		linhas[i] = "  " + l
	}
	return strings.Join(linhas, "\n")
}
