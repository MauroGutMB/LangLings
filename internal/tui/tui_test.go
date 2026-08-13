package tui

import (
	"testing"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"langlings/internal/domain"
	"langlings/internal/engine"
)

// Estes testes não abrem terminal, não sobem goroutine e não tocam em Docker.
// O Update do Bubble Tea é (Model, Msg) → (Model, Cmd): dá para construir o
// estado, mandar uma tecla e olhar o estado que volta.

func tecla(s string) tea.KeyMsg {
	switch s {
	case "enter":
		return tea.KeyMsg{Type: tea.KeyEnter}
	case "esc":
		return tea.KeyMsg{Type: tea.KeyEsc}
	case "up":
		return tea.KeyMsg{Type: tea.KeyUp}
	case "down":
		return tea.KeyMsg{Type: tea.KeyDown}
	case "ctrl+c":
		return tea.KeyMsg{Type: tea.KeyCtrlC}
	default:
		return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}
	}
}

// aplicar manda uma mensagem e devolve o Model resultante.
func aplicar(m Model, msg tea.Msg) Model {
	seguinte, _ := m.Update(msg)
	return seguinte.(Model)
}

func modeloBase() Model {
	return Model{
		spinner: spinner.New(),
		linguagens: []engine.LanguageView{
			{Language: domain.Language{Slug: "go", Name: "Go", Image: "golang:1.26-alpine"}, Installed: true,
				Summary: domain.Summary{Total: 2, Completed: 1}},
			{Language: domain.Language{Slug: "rust", Name: "Rust", Image: "rust:1-alpine"}, Installed: false},
		},
	}
}

func exerciciosDeTeste() []engine.ExerciseView {
	return []engine.ExerciseView{
		{
			Exercise: domain.Exercise{
				Path: "go/sintaxe/a", Language: "go", Category: domain.CategorySintaxe,
				Title: "Slices", Objective: "Faça X.", Editable: []string{"main.go"},
				Hints:      []string{"primeira dica", "segunda dica"},
				Validation: domain.ValidationSpec{Mode: domain.ModeTest},
			},
			Progress: domain.Progress{Status: domain.StatusCompleted},
		},
		{
			Exercise: domain.Exercise{
				Path: "go/compilador/b", Language: "go", Category: domain.CategoryCompilador,
				Title: "Cross-compile", Objective: "Gere o binário.", Editable: []string{"main.go"},
				Validation: domain.ValidationSpec{
					Mode:     domain.ModeCriteria,
					Criteria: []domain.Criterion{{Kind: domain.KindFileExists, Path: "bin/hello.exe"}},
				},
			},
			Progress: domain.Progress{Status: domain.StatusNotStarted},
		},
	}
}

// ---------- boas-vindas ----------

func TestBoasVindas_QualquerTeclaAvanca(t *testing.T) {
	m := modeloBase()

	m = aplicar(m, tecla("x"))

	require.Equal(t, telaLinguagens, m.tela)
}

func TestBoasVindas_QSai(t *testing.T) {
	m := modeloBase()

	m = aplicar(m, tecla("q"))

	require.True(t, m.Sair)
}

func TestCtrlCSaiDeQualquerTela(t *testing.T) {
	for _, tl := range []tela{telaBoasVindas, telaLinguagens, telaExercicios, telaExercicio} {
		m := modeloBase()
		m.tela = tl

		m = aplicar(m, tecla("ctrl+c"))

		require.True(t, m.Sair, "tela %d deveria sair com ctrl+c", tl)
	}
}

// ---------- linguagens ----------

func TestLinguagens_NavegacaoRespeitaOsLimites(t *testing.T) {
	m := modeloBase()
	m.tela = telaLinguagens

	// Subir no topo não passa do primeiro.
	m = aplicar(m, tecla("up"))
	require.Equal(t, 0, m.cursorIdioma)

	m = aplicar(m, tecla("down"))
	require.Equal(t, 1, m.cursorIdioma)

	// Descer no fim não passa do último.
	m = aplicar(m, tecla("down"))
	require.Equal(t, 1, m.cursorIdioma)
}

func TestLinguagens_TeclasDoVimTambemNavegam(t *testing.T) {
	m := modeloBase()
	m.tela = telaLinguagens

	m = aplicar(m, tecla("j"))
	require.Equal(t, 1, m.cursorIdioma)

	m = aplicar(m, tecla("k"))
	require.Equal(t, 0, m.cursorIdioma)
}

func TestLinguagens_InstaladaAbreAsAtividades(t *testing.T) {
	m := modeloBase()
	m.tela = telaLinguagens
	m.cursorIdioma = 0

	m = aplicar(m, tecla("enter"))

	require.Equal(t, telaExercicios, m.tela)
	require.Equal(t, "go", m.linguagemAtual.Slug)
}

func TestLinguagens_NaoInstaladaDisparaAInstalacao(t *testing.T) {
	// Selecionar é o gatilho: não existe um passo separado de "instalar".
	m := modeloBase()
	m.tela = telaLinguagens
	m.cursorIdioma = 1

	m = aplicar(m, tecla("enter"))

	require.Equal(t, telaInstalando, m.tela)
	require.True(t, m.instalando)
	require.Equal(t, "rust", m.alvoInstalacao.Slug)
}

func TestLinguagens_ListaVaziaNaoQuebra(t *testing.T) {
	m := Model{spinner: spinner.New(), tela: telaLinguagens}

	m = aplicar(m, tecla("enter"))
	m = aplicar(m, tecla("down"))

	require.Equal(t, telaLinguagens, m.tela)
	require.Equal(t, 0, m.cursorIdioma)
}

// ---------- exercícios ----------

func TestExercicios_CursorVaiParaOPrimeiroNaoConcluido(t *testing.T) {
	m := modeloBase()

	m = aplicar(m, exerciciosCarregadosMsg{views: exerciciosDeTeste()})

	// O primeiro está concluído; a sugestão é o segundo. A progressão é livre,
	// mas o cursor aponta por onde seguir.
	require.Equal(t, 1, m.cursorExercicio)
}

func TestExercicios_TudoConcluidoVoltaAoTopo(t *testing.T) {
	views := exerciciosDeTeste()
	views[1].Progress.Status = domain.StatusCompleted
	m := modeloBase()

	m = aplicar(m, exerciciosCarregadosMsg{views: views})

	require.Equal(t, 0, m.cursorExercicio)
}

func TestExercicios_EscVoltaParaLinguagens(t *testing.T) {
	m := modeloBase()
	m.tela = telaExercicios
	m.exercicios = exerciciosDeTeste()

	m = aplicar(m, tecla("esc"))

	require.Equal(t, telaLinguagens, m.tela)
}

func TestExercicios_EnterAbreOExercicio(t *testing.T) {
	m := modeloBase()
	m.tela = telaExercicios
	m.exercicios = exerciciosDeTeste()
	m.cursorExercicio = 1
	m.dicasAbertas = 3

	m = aplicar(m, tecla("enter"))

	require.Equal(t, telaExercicio, m.tela)
	require.Equal(t, "go/compilador/b", m.atual.Exercise.Path)
	require.Zero(t, m.dicasAbertas, "as dicas do exercício anterior não podem vazar")
	require.Nil(t, m.ultimo)
}

// ---------- exercício ----------

func modeloNoExercicio(t *testing.T) (Model, *bool) {
	t.Helper()
	cancelado := false

	m := modeloBase()
	m.tela = telaExercicio
	m.linguagemAtual = domain.Language{Slug: "go", Name: "Go"}
	m.exercicios = exerciciosDeTeste()
	m.atual = m.exercicios[0]
	m.sessao = &engine.Session{Dir: "/home/user/.local/share/langlings/workspace/go/sintaxe/a"}
	m.cancelarSessao = func() { cancelado = true }

	return m, &cancelado
}

func TestExercicio_EscFechaOWatcher(t *testing.T) {
	m, cancelado := modeloNoExercicio(t)

	m = aplicar(m, tecla("esc"))

	require.Equal(t, telaExercicios, m.tela)
	require.True(t, *cancelado, "sair da tela precisa parar de observar o arquivo")
	require.Nil(t, m.sessao)
}

func TestExercicio_DicasSaoReveladasUmaAUma(t *testing.T) {
	m, _ := modeloNoExercicio(t)
	require.Len(t, m.atual.Exercise.Hints, 2)

	m = aplicar(m, tecla("h"))
	require.Equal(t, 1, m.dicasAbertas)

	m = aplicar(m, tecla("h"))
	require.Equal(t, 2, m.dicasAbertas)

	// Não passa do número de dicas existentes.
	m = aplicar(m, tecla("h"))
	require.Equal(t, 2, m.dicasAbertas)
}

func TestExercicio_SpinnerAparecerNoInstanteDoSave(t *testing.T) {
	// O aviso de "validando" chega antes de qualquer resultado. Sem ele,
	// três segundos de compilação parecem um travamento.
	m, _ := modeloNoExercicio(t)
	require.False(t, m.validando)

	m = aplicar(m, atualizacaoMsg{update: engine.Update{Kind: engine.UpdateValidating}})

	require.True(t, m.validando)
	require.Contains(t, m.View(), "validando")
}

func TestExercicio_ResultadoAtualizaOEstado(t *testing.T) {
	m, _ := modeloNoExercicio(t)
	m.validando = true

	m = aplicar(m, atualizacaoMsg{update: engine.Update{
		Kind: engine.UpdateResult,
		Result: engine.Result{
			Passed:   true,
			Output:   "ok\n",
			Progress: domain.Progress{Status: domain.StatusCompleted},
		},
	}})

	require.False(t, m.validando)
	require.NotNil(t, m.ultimo)
	require.True(t, m.ultimo.Passed)
	require.Equal(t, domain.StatusCompleted, m.atual.Progress.Status)
	require.Contains(t, m.View(), "✓ passou")
}

func TestExercicio_ReprovacaoMostraASaidaBruta(t *testing.T) {
	m, _ := modeloNoExercicio(t)

	m = aplicar(m, atualizacaoMsg{update: engine.Update{
		Kind: engine.UpdateResult,
		Result: engine.Result{
			Passed:   false,
			Output:   "--- FAIL: TestDobrar\n    main_test.go:12: quero [2 4 6]",
			Progress: domain.Progress{Status: domain.StatusInProgress},
		},
	}})

	view := m.View()
	require.Contains(t, view, "✗ ainda não")
	require.Contains(t, view, "quero [2 4 6]", "a saída do test runner precisa chegar aos olhos do usuário")
}

func TestExercicio_TimeoutTemMensagemPropria(t *testing.T) {
	// Estourar o tempo é diferente de errar o exercício.
	m, _ := modeloNoExercicio(t)

	m = aplicar(m, atualizacaoMsg{update: engine.Update{
		Kind:   engine.UpdateResult,
		Result: engine.Result{Passed: false, TimedOut: true},
	}})

	require.Contains(t, m.View(), "excedeu")
}

func TestExercicio_MostraOCaminhoNoHost(t *testing.T) {
	// É a informação mais importante da tela: é o que o usuário abre no editor.
	m, _ := modeloNoExercicio(t)

	view := m.View()

	require.Contains(t, view, "/home/user/.local/share/langlings/workspace/go/sintaxe/a")
	require.Contains(t, view, "main.go")
}

func TestExercicio_ModoCriteriaExplicaOFluxoManual(t *testing.T) {
	m, _ := modeloNoExercicio(t)
	m.atual = m.exercicios[1] // o de critérios

	view := m.View()

	require.Contains(t, view, "bin/hello.exe existe")
	require.Contains(t, view, "[s]", "o usuário precisa saber como abrir o shell")
}

func TestExercicio_RegressaoEhAvisadaSemPerderOProgresso(t *testing.T) {
	m, _ := modeloNoExercicio(t)
	agora := time.Now()
	m.atual.Progress = domain.Progress{
		Status: domain.StatusCompleted, LastValidatedAt: &agora, LastPassed: false,
	}

	m = aplicar(m, atualizacaoMsg{update: engine.Update{
		Kind: engine.UpdateResult,
		Result: engine.Result{
			Passed:   false,
			Progress: m.atual.Progress,
		},
	}})

	view := m.View()
	require.Contains(t, view, "já foi concluído")
	require.Contains(t, view, "🟢", "o ícone continua verde: progresso é histórico")
}

func TestExercicio_CanalFechadoNaoQuebra(t *testing.T) {
	m, _ := modeloNoExercicio(t)

	m = aplicar(m, atualizacaoMsg{fim: true})

	require.Nil(t, m.sessao)
}

// ---------- listas e utilitários ----------

func TestViewExercicios_AgrupaPorCategoria(t *testing.T) {
	m := modeloBase()
	m.tela = telaExercicios
	m.linguagemAtual = domain.Language{Slug: "go", Name: "Go"}
	m.exercicios = exerciciosDeTeste()

	view := m.View()

	require.Contains(t, view, "Sintaxe")
	require.Contains(t, view, "Compilador/Interpretador")
	require.Contains(t, view, "🟢 ")
	require.Contains(t, view, "🔴 ")
}

func TestViewLinguagens_MostraInstalacao(t *testing.T) {
	m := modeloBase()
	m.tela = telaLinguagens

	view := m.View()

	require.Contains(t, view, "Go")
	require.Contains(t, view, "instalada")
	require.Contains(t, view, "Rust")
	require.Contains(t, view, "não instalada")
}

func TestMover(t *testing.T) {
	require.Equal(t, 0, mover(0, -1, 3))
	require.Equal(t, 1, mover(0, +1, 3))
	require.Equal(t, 2, mover(2, +1, 3))
	require.Equal(t, 0, mover(0, +1, 0), "lista vazia não move")
}

func TestUltimasLinhas(t *testing.T) {
	// A saída de um test runner põe o que interessa no fim.
	texto := "a\nb\nc\nd\ne"

	got := ultimasLinhas(texto, 2)

	require.Contains(t, got, "d")
	require.Contains(t, got, "e")
	require.NotContains(t, got, "a")
}

func TestEscritorDeLinhas(t *testing.T) {
	var recebidas []string
	w := &escritorDeLinhas{enviar: func(l string) { recebidas = append(recebidas, l) }}

	w.Write([]byte("primeira\nsegun"))
	w.Write([]byte("da\r\nterceira"))
	w.descarregar()

	require.Equal(t, []string{"primeira", "segunda", "terceira"}, recebidas,
		"linhas partidas entre writes precisam ser remontadas")
}

func TestView_NaoQuebraSemDados(t *testing.T) {
	// Uma tela pode ser renderizada antes de os dados chegarem.
	for _, tl := range []tela{telaBoasVindas, telaLinguagens, telaExercicios, telaExercicio, telaInstalando} {
		m := Model{spinner: spinner.New(), tela: tl}
		require.NotPanics(t, func() { _ = m.View() }, "tela %d entrou em pânico sem dados", tl)
	}
}

func TestView_ErroEhExibido(t *testing.T) {
	m := modeloBase()
	m.tela = telaLinguagens
	m.erro = errDeTeste{}

	require.Contains(t, m.View(), "daemon do Docker não respondeu")
}

type errDeTeste struct{}

func (errDeTeste) Error() string { return "daemon do Docker não respondeu" }
