package tui

import "github.com/charmbracelet/lipgloss"

// Cores adaptativas: o mesmo binário roda bem em terminal claro e escuro, sem
// configuração. lipgloss escolhe a variante conforme o fundo detectado.
var (
	corTitulo   = lipgloss.AdaptiveColor{Light: "#5A31F4", Dark: "#A78BFA"}
	corDestaque = lipgloss.AdaptiveColor{Light: "#0B7285", Dark: "#67E8F9"}
	corApagada  = lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"}
	corSucesso  = lipgloss.AdaptiveColor{Light: "#15803D", Dark: "#4ADE80"}
	corAtencao  = lipgloss.AdaptiveColor{Light: "#B45309", Dark: "#FBBF24"}
	corErro     = lipgloss.AdaptiveColor{Light: "#B91C1C", Dark: "#F87171"}
	corBorda    = lipgloss.AdaptiveColor{Light: "#D1D5DB", Dark: "#4B5563"}
)

var (
	estiloTitulo = lipgloss.NewStyle().
			Foreground(corTitulo).
			Bold(true)

	estiloSubtitulo = lipgloss.NewStyle().
			Foreground(corApagada)

	estiloSelecionado = lipgloss.NewStyle().
				Foreground(corDestaque).
				Bold(true)

	estiloApagado = lipgloss.NewStyle().
			Foreground(corApagada)

	estiloSucesso = lipgloss.NewStyle().Foreground(corSucesso).Bold(true)
	estiloAtencao = lipgloss.NewStyle().Foreground(corAtencao)
	estiloErro    = lipgloss.NewStyle().Foreground(corErro).Bold(true)

	estiloCaminho = lipgloss.NewStyle().
			Foreground(corDestaque).
			Bold(true)

	estiloCaixa = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(corBorda).
			Padding(0, 1)

	estiloRodape = lipgloss.NewStyle().
			Foreground(corApagada).
			MarginTop(1)

	estiloCategoria = lipgloss.NewStyle().
			Foreground(corTitulo).
			Bold(true).
			MarginTop(1)
)

// cursor devolve o marcador da linha, para que a seleção fique legível mesmo
// em terminais sem cor.
func cursor(selecionado bool) string {
	if selecionado {
		return "❯ "
	}
	return "  "
}
