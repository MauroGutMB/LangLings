package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/domain"
)

// Versão e créditos exibidos na abertura.
const (
	Versao   = "0.1.0"
	Creditos = "LangLings — exercícios de programação multi-linguagem"
)

func (m Model) teclaBoasVindas(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "q" {
		return m.encerrar()
	}
	// Qualquer outra tecla avança.
	m.tela = telaLinguagens
	return m, nil
}

func (m Model) viewBoasVindas() string {
	var b strings.Builder

	b.WriteString(estiloTitulo.Render(`
 ██╗      █████╗ ███╗   ██╗ ██████╗ ██╗     ██╗███╗   ██╗ ██████╗ ███████╗
 ██║     ██╔══██╗████╗  ██║██╔════╝ ██║     ██║████╗  ██║██╔════╝ ██╔════╝
 ██║     ███████║██╔██╗ ██║██║  ███╗██║     ██║██╔██╗ ██║██║  ███╗███████╗
 ██║     ██╔══██║██║╚██╗██║██║   ██║██║     ██║██║╚██╗██║██║   ██║╚════██║
 ███████╗██║  ██║██║ ╚████║╚██████╔╝███████╗██║██║ ╚████║╚██████╔╝███████║
 ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝`))
	b.WriteString("\n\n")

	b.WriteString("  " + estiloSubtitulo.Render(Creditos) + "\n")
	b.WriteString("  " + estiloApagado.Render("versão "+Versao) + "\n\n")

	total := m.resumoGlobal()
	if total.Total > 0 {
		b.WriteString("  " + m.renderResumo(total) + "\n\n")
	}

	b.WriteString("  " + estiloApagado.Render("Edite no seu editor. O LangLings valida sozinho ao salvar.") + "\n\n")

	if m.erro != nil {
		b.WriteString("  " + estiloErro.Render("erro: "+m.erro.Error()) + "\n\n")
	}

	b.WriteString(estiloRodape.Render("  qualquer tecla para começar · q para sair"))
	return b.String()
}

// resumoGlobal soma o progresso de todas as linguagens.
func (m Model) resumoGlobal() domain.Summary {
	var total domain.Summary
	for _, l := range m.linguagens {
		total.Total += l.Summary.Total
		total.Completed += l.Summary.Completed
		total.InProgress += l.Summary.InProgress
		total.NotStarted += l.Summary.NotStarted
	}
	return total
}

func (m Model) renderResumo(s domain.Summary) string {
	if s.Total == 0 {
		return estiloApagado.Render("nenhum exercício no catálogo")
	}
	return fmt.Sprintf("%s %s   %s %d   %s %d",
		estiloSucesso.Render(fmt.Sprintf("%d/%d", s.Completed, s.Total)),
		estiloApagado.Render("concluídos"),
		domain.StatusInProgress.Icon(), s.InProgress,
		domain.StatusNotStarted.Icon(), s.NotStarted,
	)
}
