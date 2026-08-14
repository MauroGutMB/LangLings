package main

import (
	"errors"
	"testing"
)

func TestResumirAgrupaPorCategoria(t *testing.T) {
	itens := []Item{
		{Nome: "caneta", Categoria: "papelaria", Preco: 250, Qtd: 4},
		{Nome: "caderno", Categoria: "papelaria", Preco: 1200, Qtd: 2},
		{Nome: "café", Categoria: "cozinha", Preco: 1850, Qtd: 3},
	}

	got, err := Resumir(itens)
	if err != nil {
		t.Fatalf("Resumir devolveu erro %v, não queria nenhum", err)
	}
	if len(got) != 2 {
		t.Fatalf("Resumir devolveu %v, quero 2 categorias", got)
	}
	if got["papelaria"] != 3400 {
		t.Errorf(`got["papelaria"] = %d, quero 3400`, got["papelaria"])
	}
	if got["cozinha"] != 5550 {
		t.Errorf(`got["cozinha"] = %d, quero 5550`, got["cozinha"])
	}
}

func TestResumirComUmItemSo(t *testing.T) {
	got, err := Resumir([]Item{{Nome: "x", Categoria: "a", Preco: 100, Qtd: 0}})
	if err != nil {
		t.Fatalf("Resumir devolveu erro %v", err)
	}
	if len(got) != 1 || got["a"] != 0 {
		t.Fatalf("Resumir devolveu %v, quero map[a:0]", got)
	}
}

func TestResumirSemItens(t *testing.T) {
	got, err := Resumir(nil)
	if !errors.Is(err, ErrSemItens) {
		t.Fatalf("Resumir(nil) devolveu erro %v, quero um que embrulhe ErrSemItens", err)
	}
	if got != nil {
		t.Errorf("Resumir(nil) devolveu %v, quero nil junto com o erro", got)
	}
}

func TestResumirComCategoriaVazia(t *testing.T) {
	itens := []Item{
		{Nome: "ok", Categoria: "a", Preco: 1, Qtd: 1},
		{Nome: "solto", Categoria: "", Preco: 1, Qtd: 1},
	}

	got, err := Resumir(itens)
	if !errors.Is(err, ErrCategoriaVazia) {
		t.Fatalf("Resumir devolveu erro %v, quero um que embrulhe ErrCategoriaVazia", err)
	}
	if got != nil {
		t.Errorf("Resumir devolveu %v, quero nil junto com o erro", got)
	}
}

func TestResumirNaoModificaAEntrada(t *testing.T) {
	itens := []Item{{Nome: "x", Categoria: "a", Preco: 100, Qtd: 2}}

	if _, err := Resumir(itens); err != nil {
		t.Fatalf("Resumir devolveu erro %v", err)
	}
	if itens[0].Qtd != 2 || itens[0].Preco != 100 {
		t.Fatalf("Resumir alterou a entrada: agora ela é %v", itens)
	}
}
