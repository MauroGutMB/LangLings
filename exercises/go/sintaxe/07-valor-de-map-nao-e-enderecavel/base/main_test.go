package main

import "testing"

func TestMarcarSomaNumTimeExistente(t *testing.T) {
	m := map[string]Placar{"go": {Pontos: 2}}

	Marcar(m, "go", 3)

	if m["go"].Pontos != 5 {
		t.Fatalf(`m["go"].Pontos = %d, quero 5`, m["go"].Pontos)
	}
}

func TestMarcarCriaOTimeAusente(t *testing.T) {
	m := map[string]Placar{}

	Marcar(m, "rust", 4)

	if len(m) != 1 {
		t.Fatalf("m = %v, quero exatamente uma chave", m)
	}
	if m["rust"].Pontos != 4 {
		t.Fatalf(`m["rust"].Pontos = %d, quero 4`, m["rust"].Pontos)
	}
}

func TestMarcarAcumulaVariasVezes(t *testing.T) {
	m := map[string]Placar{}

	Marcar(m, "go", 1)
	Marcar(m, "go", 2)
	Marcar(m, "go", 3)

	if m["go"].Pontos != 6 {
		t.Fatalf(`m["go"].Pontos = %d, quero 6`, m["go"].Pontos)
	}
}

func TestMarcarNaoMexeNosOutrosTimes(t *testing.T) {
	m := map[string]Placar{"go": {Pontos: 2}, "zig": {Pontos: 9}}

	Marcar(m, "go", 1)

	if m["zig"].Pontos != 9 {
		t.Fatalf(`m["zig"].Pontos = %d, quero 9`, m["zig"].Pontos)
	}
}

func TestMarcarAceitaNegativo(t *testing.T) {
	m := map[string]Placar{"go": {Pontos: 2}}

	Marcar(m, "go", -5)

	if m["go"].Pontos != -3 {
		t.Fatalf(`m["go"].Pontos = %d, quero -3`, m["go"].Pontos)
	}
}

func TestMarcarComMapNilNaoEntraEmPanico(t *testing.T) {
	var m map[string]Placar

	Marcar(m, "go", 1)

	if len(m) != 0 {
		t.Fatalf("m = %v, quero um map ainda vazio", m)
	}
}
