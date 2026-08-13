package catalog

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"langlings/internal/domain"
)

func TestLoad_CatalogoValido(t *testing.T) {
	c, err := Load("testdata/valido")
	require.NoError(t, err)

	require.Len(t, c.Languages, 1)
	require.Equal(t, "go", c.Languages[0].Slug)
	require.Equal(t, "golang:1.26-alpine", c.Languages[0].ImageRef())

	require.Len(t, c.Exercises, 2)
}

func TestLoad_OrdenaPelaOrdemCanonicaDasCategorias(t *testing.T) {
	// "compilador" vem antes de "sintaxe" no alfabeto, mas depois na ordem de
	// exibição. A ordem do catálogo é a da TUI, não a do alfabeto.
	c, err := Load("testdata/valido")
	require.NoError(t, err)

	require.Equal(t, "go/sintaxe/slices-append", c.Exercises[0].Path)
	require.Equal(t, "go/compilador/cross-compile", c.Exercises[1].Path)
}

func TestLoad_CamposDoExercicio(t *testing.T) {
	c, err := Load("testdata/valido")
	require.NoError(t, err)

	ex, ok := c.Exercise("go/sintaxe/slices-append")
	require.True(t, ok)

	require.Equal(t, "Append em slices", ex.Title)
	require.Equal(t, domain.CategorySintaxe, ex.Category)
	require.Equal(t, "go", ex.Language)
	require.Equal(t, []string{"main.go"}, ex.Editable)
	require.Len(t, ex.Hints, 1)
	require.Equal(t, domain.ModeTest, ex.Validation.Mode)
	require.Equal(t, []string{"go", "test", "./..."}, ex.Validation.Command)
	require.Equal(t, 30*time.Second, ex.Validation.Timeout)
	require.False(t, ex.Validation.Network, "rede é desligada por padrão")

	// O objective multi-linha chega sem o \n solto das aspas triplas.
	require.NotEmpty(t, ex.Objective)
	require.NotContains(t, ex.Objective[:1], "\n")
}

func TestLoad_CriteriosDoModoCriteria(t *testing.T) {
	c, err := Load("testdata/valido")
	require.NoError(t, err)

	ex, ok := c.Exercise("go/compilador/cross-compile")
	require.True(t, ok)

	require.Equal(t, domain.ModeCriteria, ex.Validation.Mode)
	require.Len(t, ex.Validation.Criteria, 2)
	require.Equal(t, domain.KindFileExists, ex.Validation.Criteria[0].Kind)
	require.Equal(t, "bin/hello.exe", ex.Validation.Criteria[0].Path)
	require.Equal(t, domain.KindStdoutMatches, ex.Validation.Criteria[1].Kind)
	require.Equal(t, `PE32\+.*x86-64`, ex.Validation.Criteria[1].Pattern)
}

func TestCatalog_Consultas(t *testing.T) {
	c, err := Load("testdata/valido")
	require.NoError(t, err)

	t.Run("ByLanguage", func(t *testing.T) {
		require.Len(t, c.ByLanguage("go"), 2)
		require.Empty(t, c.ByLanguage("rust"))
	})

	t.Run("Language", func(t *testing.T) {
		_, ok := c.Language("go")
		require.True(t, ok)
		_, ok = c.Language("cobol")
		require.False(t, ok)
	})

	t.Run("Exercise inexistente", func(t *testing.T) {
		_, ok := c.Exercise("go/sintaxe/nao-existe")
		require.False(t, ok)
	})
}

func TestLoad_DiretorioInexistente(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "vazio"))
	require.Error(t, err)
	require.Contains(t, err.Error(), "não encontrado")
}

// ---------- manifestos inválidos ----------
//
// Cada caso monta um catálogo mínimo num diretório temporário. Ficam inline
// em vez de em testdata/ porque o manifesto quebrado e a mensagem esperada
// se leem melhor lado a lado.

const languageOK = `
name  = "Go"
image = "golang:1.26-alpine"
`

// escreveCatalogo monta um catálogo temporário com um único exercício de
// sintaxe cujo manifesto é o texto informado, e devolve a raiz.
func escreveCatalogo(t *testing.T, exerciseTOML string, extras ...string) string {
	t.Helper()
	root := t.TempDir()

	exDir := filepath.Join(root, "exercises", "go", "sintaxe", "exemplo")
	require.NoError(t, os.MkdirAll(filepath.Join(root, "languages", "go"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(exDir, "base"), 0o755))

	write := func(path, content string) {
		require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	}
	write(filepath.Join(root, "languages", "go", LanguageFile), languageOK)
	write(filepath.Join(exDir, ExerciseFile), exerciseTOML)
	write(filepath.Join(exDir, "base", "main.go"), "package main\n")

	// extras vêm em pares path-relativo/conteúdo.
	for i := 0; i+1 < len(extras); i += 2 {
		full := filepath.Join(root, filepath.FromSlash(extras[i]))
		require.NoError(t, os.MkdirAll(filepath.Dir(full), 0o755))
		write(full, extras[i+1])
	}
	return root
}

func TestLoad_ManifestosInvalidos(t *testing.T) {
	tests := []struct {
		nome     string
		manifest string
		querErro string
	}{
		{
			nome: "chave desconhecida não passa em silêncio",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]
timeouts  = "30s"

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: "desconhecida",
		},
		{
			nome: "chave desconhecida dentro de validation",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode     = "test"
command  = ["go", "test"]
timeoout = "30s"
`,
			querErro: "desconhecida",
		},
		{
			nome: "TOML sintaticamente quebrado",
			manifest: `
title = "X
`,
			querErro: "exercise.toml",
		},
		{
			nome: "sem objective",
			manifest: `
title    = "X"
editable = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: "objective é obrigatório",
		},
		{
			nome: "sem editable",
			manifest: `
title     = "X"
objective = "Y"

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: "editable é obrigatório",
		},
		{
			nome: "editable escapando do exercício",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["../../../etc/passwd"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: "escapar",
		},
		{
			nome: "modo desconhecido",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode = "manual"
`,
			querErro: "desconhecido",
		},
		{
			nome: "modo criteria sem critério",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode = "criteria"
`,
			querErro: "pelo menos um",
		},
		{
			nome: "timeout malformado",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
timeout = "30 segundos"
`,
			querErro: "timeout",
		},
		{
			nome: "regex quebrada falha no boot, não durante a validação",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode = "criteria"

[[validation.criteria]]
kind    = "stdout_matches"
command = ["file", "bin/x"]
pattern = "PE32(+"
`,
			querErro: "pattern inválido",
		},
		{
			nome: "categoria do manifesto divergindo do diretório",
			manifest: `
title     = "X"
category  = "compilador"
objective = "Y"
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: "o diretório manda",
		},
		{
			nome: "editable declarado mas ausente em base/",
			manifest: `
title     = "X"
objective = "Y"
editable  = ["main.go", "calc.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			querErro: `"calc.go" não existe`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			root := escreveCatalogo(t, tt.manifest)

			c, err := Load(root)

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.querErro)
			require.Empty(t, c.Exercises, "exercício inválido não entra no catálogo")
		})
	}
}

func TestLoad_ErroCitaOExercicio(t *testing.T) {
	root := escreveCatalogo(t, `
title     = "X"
objective = "Y"

[validation]
mode    = "test"
command = ["go", "test"]
`)

	_, err := Load(root)

	require.Error(t, err)
	require.Contains(t, err.Error(), "go/sintaxe/exemplo", "a mensagem precisa dizer QUAL exercício")
}

func TestLoad_AcumulaProblemas(t *testing.T) {
	// Dois exercícios quebrados devem produzir dois erros, não parar no primeiro.
	root := escreveCatalogo(t, `
title     = "X"
objective = "Y"

[validation]
mode    = "test"
command = ["go", "test"]
`,
		"exercises/go/sintaxe/outro/"+ExerciseFile, `
title    = "Sem objective também"
editable = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
		"exercises/go/sintaxe/outro/base/main.go", "package main\n",
	)

	_, err := Load(root)

	require.Error(t, err)
	require.Contains(t, err.Error(), "go/sintaxe/exemplo")
	require.Contains(t, err.Error(), "go/sintaxe/outro")
}

func TestLoad_CategoriaOuLinguagemDesconhecida(t *testing.T) {
	t.Run("categoria fora das quatro", func(t *testing.T) {
		root := escreveCatalogo(t, `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			"exercises/go/algoritmos/x/"+ExerciseFile, "title = \"X\"\n")

		_, err := Load(root)
		require.Error(t, err)
		require.Contains(t, err.Error(), "categoria desconhecida")
	})

	t.Run("exercício de linguagem não declarada", func(t *testing.T) {
		root := escreveCatalogo(t, `
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`,
			"exercises/rust/sintaxe/x/"+ExerciseFile, "title = \"X\"\n")

		_, err := Load(root)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nenhuma linguagem")
	})
}

func TestLoad_ManifestoDeLinguagemInvalido(t *testing.T) {
	tests := []struct {
		nome     string
		manifest string
		querErro string
	}{
		{"sem name", `image = "golang:1.26-alpine"`, "name é obrigatório"},
		{"sem image nem dockerfile", `name = "Go"`, "declare image"},
		{"image e dockerfile juntos", "name = \"Go\"\nimage = \"golang:1.26-alpine\"\ndockerfile = \"Dockerfile\"", "não os dois"},
		{"workdir relativo", "name = \"Go\"\nimage = \"x\"\nworkdir = \"workspace\"", "absoluto"},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			root := t.TempDir()
			require.NoError(t, os.MkdirAll(filepath.Join(root, "languages", "go"), 0o755))
			require.NoError(t, os.MkdirAll(filepath.Join(root, "exercises"), 0o755))
			require.NoError(t, os.WriteFile(
				filepath.Join(root, "languages", "go", LanguageFile), []byte(tt.manifest), 0o644))

			_, err := Load(root)

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.querErro)
		})
	}
}

func TestLoad_PadroesDaLinguagem(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(root, "languages", "lua"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "exercises"), 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "languages", "lua", LanguageFile),
		[]byte("name = \"Lua\"\nimage = \"lua:5.4\"\n"), 0o644))

	c, err := Load(root)
	require.NoError(t, err)
	require.Len(t, c.Languages, 1)

	require.Equal(t, DefaultWorkdir, c.Languages[0].Workdir)
	require.Equal(t, []string{"/bin/sh"}, c.Languages[0].Shell)
}

func TestLoad_ExercicioSemBase(t *testing.T) {
	root := t.TempDir()
	exDir := filepath.Join(root, "exercises", "go", "sintaxe", "exemplo")
	require.NoError(t, os.MkdirAll(filepath.Join(root, "languages", "go"), 0o755))
	require.NoError(t, os.MkdirAll(exDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "languages", "go", LanguageFile), []byte(languageOK), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(exDir, ExerciseFile), []byte(`
title     = "X"
objective = "Y"
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test"]
`), 0o644))

	_, err := Load(root)

	require.Error(t, err)
	require.Contains(t, err.Error(), "base/ ausente")
}
