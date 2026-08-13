package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) teclaLinguagens(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case ehVoltar(msg):
		return m.encerrar()

	case ehParaCima(msg):
		m.cursorIdioma = mover(m.cursorIdioma, -1, len(m.linguagens))
		return m, nil

	case ehParaBaixo(msg):
		m.cursorIdioma = mover(m.cursorIdioma, +1, len(m.linguagens))
		return m, nil

	case msg.String() == "r":
		return m, carregarLinguagens(m.ctx, m.motor)

	case ehConfirmar(msg):
		if len(m.linguagens) == 0 {
			return m, nil
		}
		escolhida := m.linguagens[m.cursorIdioma]
		m.aviso = ""

		// Selecionar uma linguagem não instalada é o gatilho da instalação:
		// não há um passo separado de "instalar" a ser lembrado.
		if !escolhida.Installed {
			m.tela = telaInstalando
			m.instalando = true
			m.alvoInstalacao = escolhida.Language
			m.logInstalacao = nil
			return m, instalar(m.ctx, m.motor, escolhida.Language, m.program)
		}

		m.linguagemAtual = escolhida.Language
		m.tela = telaExercicios
		return m, carregarExercicios(m.motor, escolhida.Language.Slug)
	}

	return m, nil
}

func (m Model) viewLinguagens() string {
	var b strings.Builder

	b.WriteString(estiloTitulo.Render("Linguagens") + "\n")
	b.WriteString(estiloSubtitulo.Render("escolha uma para ver as atividades") + "\n\n")

	if len(m.linguagens) == 0 {
		b.WriteString(estiloApagado.Render("  nenhuma linguagem no catálogo") + "\n")
	}

	for i, l := range m.linguagens {
		selecionada := i == m.cursorIdioma

		nome := l.Language.Name
		if selecionada {
			nome = estiloSelecionado.Render(nome)
		}

		var marca string
		if l.Installed {
			marca = estiloSucesso.Render("instalada")
		} else {
			marca = estiloAtencao.Render("não instalada")
		}

		linha := fmt.Sprintf("%s%-12s %s", cursor(selecionada), nome, marca)
		if l.Summary.Total > 0 {
			linha += estiloApagado.Render(fmt.Sprintf("  %d/%d", l.Summary.Completed, l.Summary.Total))
		}
		b.WriteString(linha + "\n")
	}

	if m.aviso != "" {
		b.WriteString("\n" + estiloSucesso.Render("  "+m.aviso) + "\n")
	}
	if m.erro != nil {
		b.WriteString("\n" + estiloErro.Render("  erro: "+m.erro.Error()) + "\n")
	}

	b.WriteString(estiloRodape.Render("\n↑/↓ navegar · enter abrir ou instalar · r recarregar · q sair"))
	return b.String()
}

func (m Model) viewInstalando() string {
	var b strings.Builder

	b.WriteString(estiloTitulo.Render("Instalando "+m.alvoInstalacao.Name) + "\n")
	b.WriteString(estiloSubtitulo.Render(m.alvoInstalacao.ImageRef()) + "\n\n")

	if m.instalando {
		b.WriteString("  " + m.spinner.View() + estiloApagado.Render(" preparando a imagem…") + "\n\n")
	}

	// O log é exibido em streaming: baixar uma imagem leva minutos, e uma tela
	// parada nesse tempo é indistinguível de um travamento.
	inicio := 0
	if len(m.logInstalacao) > 15 {
		inicio = len(m.logInstalacao) - 15
	}
	for _, linha := range m.logInstalacao[inicio:] {
		b.WriteString(estiloApagado.Render("  "+linha) + "\n")
	}

	if m.erro != nil {
		b.WriteString("\n" + estiloErro.Render("  erro: "+m.erro.Error()) + "\n")
	}

	return b.String()
}
