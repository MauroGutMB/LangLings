package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/domain"
	"langlings/internal/engine"
)

func (m Model) teclaExercicios(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case ehVoltar(msg):
		m.tela = telaLinguagens
		m.aviso = ""
		return m, carregarLinguagens(m.ctx, m.motor)

	case ehParaCima(msg):
		m.cursorExercicio = mover(m.cursorExercicio, -1, len(m.exercicios))
		return m, nil

	case ehParaBaixo(msg):
		m.cursorExercicio = mover(m.cursorExercicio, +1, len(m.exercicios))
		return m, nil

	case msg.String() == "R":
		// Maiúsculo de propósito: zerar a linguagem inteira não pode ser um
		// deslize de dedo.
		slug := m.linguagemAtual.Slug
		return m, func() tea.Msg {
			n, err := m.motor.ResetLanguage(slug, engine.ResetEstado)
			return resetFeitoMsg{mensagem: fmt.Sprintf("%d exercícios zerados.", n), err: err}
		}

	case ehConfirmar(msg):
		if len(m.exercicios) == 0 {
			return m, nil
		}
		m.atual = m.exercicios[m.cursorExercicio]
		m.tela = telaExercicio
		m.ultimo = nil
		m.dicasAbertas = 0
		m.aviso = ""
		m.erro = nil
		return m, abrirExercicio(m.ctx, m.motor, m.atual.Exercise)
	}

	return m, nil
}

func (m Model) viewExercicios() string {
	var b strings.Builder

	b.WriteString(estiloTitulo.Render(m.linguagemAtual.Name) + "\n")
	b.WriteString(estiloSubtitulo.Render("atividades") + "\n")

	if len(m.exercicios) == 0 {
		b.WriteString("\n" + estiloApagado.Render("  nenhuma atividade para esta linguagem ainda") + "\n")
	}

	categoriaAtual := domain.Category("")
	for i, v := range m.exercicios {
		if v.Exercise.Category != categoriaAtual {
			categoriaAtual = v.Exercise.Category
			b.WriteString(estiloCategoria.Render("  "+categoriaAtual.Label()) + "\n")
		}

		selecionado := i == m.cursorExercicio
		titulo := v.Exercise.Title
		if selecionado {
			titulo = estiloSelecionado.Render(titulo)
		}

		linha := fmt.Sprintf("  %s%s %s", cursor(selecionado), v.Progress.Status.Icon(), titulo)
		if v.Progress.Regressed() {
			linha += " " + estiloAtencao.Render("(regrediu)")
		}
		if v.Progress.Orphaned {
			linha += " " + estiloApagado.Render("(órfão)")
		}
		b.WriteString(linha + "\n")
	}

	if m.aviso != "" {
		b.WriteString("\n" + estiloSucesso.Render("  "+m.aviso) + "\n")
	}
	if m.erro != nil {
		b.WriteString("\n" + estiloErro.Render("  erro: "+m.erro.Error()) + "\n")
	}

	b.WriteString(estiloRodape.Render("\n↑/↓ navegar · enter abrir · R zerar a linguagem · esc voltar"))
	return b.String()
}
