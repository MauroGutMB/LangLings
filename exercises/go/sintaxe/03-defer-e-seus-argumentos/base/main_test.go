package main

import (
	"errors"
	"testing"
)

// gravador guarda tudo o que Processar registrou, para que os testes possam
// cobrar tanto o valor quanto o número de chamadas.
type gravador struct{ valores []int }

func (g *gravador) registrar(n int) { g.valores = append(g.valores, n) }

func TestProcessarSomaEDevolveOTotal(t *testing.T) {
	var g gravador

	total, err := Processar([]int{1, 2, 3}, g.registrar)
	if err != nil {
		t.Fatalf("Processar devolveu erro %v, não queria nenhum", err)
	}
	if total != 6 {
		t.Fatalf("total = %d, quero 6", total)
	}
}

// Este é o teste que pega a armadilha: um defer cujo argumento foi avaliado
// cedo demais registra 0 mesmo com o total valendo 6.
func TestProcessarRegistraOTotalJaCalculado(t *testing.T) {
	var g gravador

	Processar([]int{1, 2, 3}, g.registrar)

	if len(g.valores) != 1 {
		t.Fatalf("registrar foi chamado %d vez(es) com %v, quero exatamente 1", len(g.valores), g.valores)
	}
	if g.valores[0] != 6 {
		t.Fatalf("registrar recebeu %d, quero 6 (o valor foi lido antes da soma?)", g.valores[0])
	}
}

func TestProcessarComSliceVazio(t *testing.T) {
	var g gravador

	total, err := Processar(nil, g.registrar)
	if !errors.Is(err, ErrVazio) {
		t.Fatalf("Processar(nil) devolveu erro %v, quero um que embrulhe ErrVazio", err)
	}
	if total != 0 {
		t.Fatalf("total = %d, quero 0", total)
	}
}

// O caminho de erro sai por um return diferente e precisa registrar do mesmo
// jeito — é por isso que a chamada é diferida em vez de repetida.
func TestProcessarRegistraTambemNoCaminhoDeErro(t *testing.T) {
	var g gravador

	Processar([]int{}, g.registrar)

	if len(g.valores) != 1 {
		t.Fatalf("registrar foi chamado %d vez(es) com %v, quero exatamente 1", len(g.valores), g.valores)
	}
	if g.valores[0] != 0 {
		t.Fatalf("registrar recebeu %d, quero 0", g.valores[0])
	}
}

func TestProcessarComNegativos(t *testing.T) {
	var g gravador

	total, err := Processar([]int{10, -4, -6}, g.registrar)
	if err != nil {
		t.Fatalf("Processar devolveu erro %v", err)
	}
	if total != 0 {
		t.Fatalf("total = %d, quero 0", total)
	}
	if len(g.valores) != 1 || g.valores[0] != 0 {
		t.Fatalf("registrar recebeu %v, quero exatamente [0]", g.valores)
	}
}
