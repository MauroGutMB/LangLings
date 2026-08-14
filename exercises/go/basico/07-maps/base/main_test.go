package main

import "testing"

func TestContagemAgrupaRepeticoes(t *testing.T) {
	got := Contagem([]string{"go", "rust", "go", "go"})

	if len(got) != 2 {
		t.Fatalf("Contagem devolveu %v, quero 2 chaves distintas", got)
	}
	if got["go"] != 3 {
		t.Errorf(`got["go"] = %d, quero 3`, got["go"])
	}
	if got["rust"] != 1 {
		t.Errorf(`got["rust"] = %d, quero 1`, got["rust"])
	}
}

func TestContagemDistingueMaiusculas(t *testing.T) {
	got := Contagem([]string{"Go", "go"})

	if got["Go"] != 1 || got["go"] != 1 {
		t.Fatalf("Contagem devolveu %v, quero Go:1 e go:1", got)
	}
}

func TestContagemNaoInventaChaves(t *testing.T) {
	got := Contagem([]string{"a"})

	if _, ok := got["b"]; ok {
		t.Fatalf("Contagem devolveu %v, mas a chave %q não deveria existir", got, "b")
	}
}

func TestContagemComSliceVazio(t *testing.T) {
	if got := Contagem(nil); len(got) != 0 {
		t.Fatalf("Contagem(nil) = %v, quero tamanho 0", got)
	}
}
