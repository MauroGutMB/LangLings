package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentHash_EstavelSobOrdemDoMap(t *testing.T) {
	// Maps em Go iteram em ordem aleatória. Se o hash dependesse dessa ordem,
	// o mesmo conteúdo produziria hashes diferentes e o engine revalidaria
	// sem necessidade — ou pior, deixaria de revalidar.
	files := map[string][]byte{
		"main.go":  []byte("package main"),
		"calc.go":  []byte("func Dobrar() {}"),
		"go.mod":   []byte("module x"),
		"extra.go": []byte("// nada"),
	}

	primeiro := ContentHash(files)
	for i := 0; i < 50; i++ {
		require.Equal(t, primeiro, ContentHash(files))
	}
}

func TestContentHash_DetectaMudancas(t *testing.T) {
	base := map[string][]byte{"main.go": []byte("package main")}

	tests := []struct {
		nome  string
		files map[string][]byte
	}{
		{"conteúdo diferente", map[string][]byte{"main.go": []byte("package other")}},
		{"nome diferente", map[string][]byte{"outro.go": []byte("package main")}},
		{"arquivo a mais", map[string][]byte{"main.go": []byte("package main"), "b.go": []byte("")}},
		{"vazio", map[string][]byte{}},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			require.NotEqual(t, ContentHash(base), ContentHash(tt.files))
		})
	}
}

func TestContentHash_NaoConfundeFronteiraEntreArquivos(t *testing.T) {
	// Sem incluir o tamanho de cada parte no resumo, estes dois conjuntos
	// produziriam a mesma concatenação e portanto o mesmo hash.
	a := map[string][]byte{"f": []byte("xy"), "g": []byte("")}
	b := map[string][]byte{"f": []byte("x"), "g": []byte("y")}

	require.NotEqual(t, ContentHash(a), ContentHash(b))
}

func TestContentHash_MesmoConteudoMesmoHash(t *testing.T) {
	a := map[string][]byte{"main.go": []byte("package main\n"), "go.mod": []byte("module x\n")}
	b := map[string][]byte{"go.mod": []byte("module x\n"), "main.go": []byte("package main\n")}

	require.Equal(t, ContentHash(a), ContentHash(b))
}
