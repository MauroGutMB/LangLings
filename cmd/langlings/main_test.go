package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubirAte_EncontraOConteudoDeUmaSubpasta(t *testing.T) {
	// Rodar o comando de dentro de exercises/go/… precisa funcionar.
	raiz, ok := subirAte(".")
	require.True(t, ok, "o teste roda dentro do próprio projeto")
	require.DirExists(t, filepath.Join(raiz, "exercises"))
	require.DirExists(t, filepath.Join(raiz, "languages"))
}

func TestSubirAte_DiretorioSemConteudo(t *testing.T) {
	_, ok := subirAte(t.TempDir())
	require.False(t, ok)
}

func TestTemConteudo(t *testing.T) {
	vazio := t.TempDir()
	require.False(t, temConteudo(vazio))

	require.NoError(t, os.MkdirAll(filepath.Join(vazio, "exercises"), 0o755))
	require.False(t, temConteudo(vazio), "só metade do layout não conta")

	require.NoError(t, os.MkdirAll(filepath.Join(vazio, "languages"), 0o755))
	require.True(t, temConteudo(vazio))
}

// TestCatalogo é o gate de conteúdo dentro da suíte de testes.
//
// Com ele, `go test ./...` reprova se alguém — você agora, ou o Claude Code
// escrevendo em lote na Fase 2 — criar um exercício cujo código-base já passa,
// cuja solução não passa, ou cuja solução exige mexer em arquivo que o usuário
// não pode editar. É o que separa confiar no conteúdo de verificá-lo.
func TestCatalogo(t *testing.T) {
	if testing.Short() {
		t.Skip("requer Docker; rode sem -short")
	}

	raiz, ok := subirAte(".")
	require.True(t, ok)

	ctx := context.Background()
	a, err := montar(ctx, raiz)
	if err != nil {
		if strings.Contains(err.Error(), "Docker") {
			t.Skipf("Docker indisponível: %v", err)
		}
		require.NoError(t, err)
	}
	defer a.Close()
	defer a.engine.Shutdown(context.Background())

	relatorios := a.engine.VerifyAll(ctx)
	require.NotEmpty(t, relatorios, "o catálogo não pode estar vazio")

	for _, r := range relatorios {
		t.Run(r.Exercise.Path, func(t *testing.T) {
			require.True(t, r.OK(), "problemas:\n  %s", strings.Join(r.Problems, "\n  "))
		})
	}
}
