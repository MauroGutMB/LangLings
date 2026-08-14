package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func exercicioValido() Exercise {
	return Exercise{
		Path:      "go/sintaxe/slices-append",
		Language:  "go",
		Category:  CategorySintaxe,
		Title:     "Append em slices",
		Objective: "Faça Dobrar devolver um slice novo.",
		Editable:  []string{"main.go"},
		Validation: ValidationSpec{
			Mode:    ModeTest,
			Command: []string{"go", "test", "./..."},
			Timeout: DefaultTimeout,
		},
	}
}

func TestExercise_Validate(t *testing.T) {
	tests := []struct {
		nome     string
		ajuste   func(*Exercise)
		querErro string // substring; vazio = deve passar
	}{
		{
			nome:   "exercício completo passa",
			ajuste: func(e *Exercise) {},
		},
		{
			nome:     "sem title",
			ajuste:   func(e *Exercise) { e.Title = "" },
			querErro: "title é obrigatório",
		},
		{
			nome:     "sem objective",
			ajuste:   func(e *Exercise) { e.Objective = "  " },
			querErro: "objective é obrigatório",
		},
		{
			nome:     "categoria inventada",
			ajuste:   func(e *Exercise) { e.Category = "algoritmos" },
			querErro: "desconhecida",
		},
		{
			nome:     "sem editable o watcher não sabe o que observar",
			ajuste:   func(e *Exercise) { e.Editable = nil },
			querErro: "editable é obrigatório",
		},
		{
			nome:     "editable com path absoluto",
			ajuste:   func(e *Exercise) { e.Editable = []string{"/etc/passwd"} },
			querErro: "relativo",
		},
		{
			nome:     "editable escapando do exercício",
			ajuste:   func(e *Exercise) { e.Editable = []string{"../../../.ssh/id_rsa"} },
			querErro: "escapar",
		},
		{
			nome:     "editable disfarçando o escape",
			ajuste:   func(e *Exercise) { e.Editable = []string{"sub/../../fora.go"} },
			querErro: "escapar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			ex := exercicioValido()
			tt.ajuste(&ex)

			err := ex.Validate()
			if tt.querErro == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.querErro)
		})
	}
}

func TestValidationSpec_Validate(t *testing.T) {
	criterioOK := Criterion{Kind: KindFileExists, Path: "bin/hello"}

	tests := []struct {
		nome     string
		spec     ValidationSpec
		querErro string
	}{
		{
			nome: "modo test válido",
			spec: ValidationSpec{Mode: ModeTest, Command: []string{"go", "test"}, Timeout: time.Second},
		},
		{
			nome: "modo criteria válido",
			spec: ValidationSpec{Mode: ModeCriteria, Criteria: []Criterion{criterioOK}, Timeout: time.Second},
		},
		{
			nome:     "modo test sem comando",
			spec:     ValidationSpec{Mode: ModeTest, Timeout: time.Second},
			querErro: "command é obrigatório",
		},
		{
			nome:     "modo criteria sem critério",
			spec:     ValidationSpec{Mode: ModeCriteria, Timeout: time.Second},
			querErro: "pelo menos um",
		},
		{
			nome: "modo test com criteria é confusão de modo",
			spec: ValidationSpec{
				Mode: ModeTest, Command: []string{"go", "test"},
				Criteria: []Criterion{criterioOK}, Timeout: time.Second,
			},
			querErro: "não se aplica",
		},
		{
			nome: "modo criteria com command é confusão de modo",
			spec: ValidationSpec{
				Mode: ModeCriteria, Criteria: []Criterion{criterioOK},
				Command: []string{"go", "test"}, Timeout: time.Second,
			},
			querErro: "não se aplica",
		},
		{
			nome:     "modo desconhecido",
			spec:     ValidationSpec{Mode: "manual", Timeout: time.Second},
			querErro: "desconhecido",
		},
		{
			nome:     "timeout zero",
			spec:     ValidationSpec{Mode: ModeTest, Command: []string{"go"}},
			querErro: "timeout deve ser positivo",
		},
		{
			nome: "critério inválido é reportado com o índice",
			spec: ValidationSpec{
				Mode:     ModeCriteria,
				Criteria: []Criterion{criterioOK, {Kind: KindExitCode}},
				Timeout:  time.Second,
			},
			querErro: "criteria[1]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			err := tt.spec.Validate()
			if tt.querErro == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.querErro)
		})
	}
}

func TestLanguage_ImageRef(t *testing.T) {
	t.Run("imagem pronta é usada como está", func(t *testing.T) {
		l := Language{Slug: "go", Image: "golang:1.26-alpine"}
		require.Equal(t, "golang:1.26-alpine", l.ImageRef())
		require.False(t, l.BuildsOwnImage())
	})

	t.Run("com Dockerfile o nome é derivado do slug", func(t *testing.T) {
		l := Language{Slug: "rust", Dockerfile: "Dockerfile"}
		require.Equal(t, "langlings/rust:latest", l.ImageRef())
		require.True(t, l.BuildsOwnImage())
	})
}

func TestCategory(t *testing.T) {
	require.True(t, CategoryCompilador.Valid())
	require.True(t, CategoryBasico.Valid())
	require.False(t, Category("algoritmos").Valid())
	require.Equal(t, "Compilador/Interpretador", CategoryCompilador.Label())
	require.Equal(t, "Básico", CategoryBasico.Label())

	// A ordem de Categories é a ordem de exibição na TUI, não detalhe interno:
	// Básico precisa vir antes de Sintaxe, ou o aluno encontra a armadilha
	// idiomática antes da introdução à linguagem.
	require.Equal(t, CategoryBasico, Categories[0])
	require.Equal(t, CategorySintaxe, Categories[1])
}

func TestCategoriesHint(t *testing.T) {
	// A dica é derivada de Categories justamente para não virar uma terceira
	// cópia da lista, que uma categoria nova deixaria mentindo em silêncio.
	require.Equal(t, "basico, sintaxe, compilador, frameworks, exemplos", CategoriesHint())
}

func TestExercise_Layout(t *testing.T) {
	ex := Exercise{Dir: "/repo/exercises/go/sintaxe/slices-append"}
	require.Equal(t, "/repo/exercises/go/sintaxe/slices-append/base", ex.BaseDir())
	require.Equal(t, "/repo/exercises/go/sintaxe/slices-append/solution", ex.SolutionDir())
}
