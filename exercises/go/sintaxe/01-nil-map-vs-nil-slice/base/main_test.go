package main

import "testing"

func TestAcumularSomaNumaChaveExistente(t *testing.T) {
	m := map[string]int{"a": 1}

	got := Acumular(m, "a", 2)

	if got == nil {
		t.Fatal("Acumular devolveu nil")
	}
	if got["a"] != 3 {
		t.Fatalf(`got["a"] = %d, quero 3`, got["a"])
	}
}

func TestAcumularCriaAChaveAusente(t *testing.T) {
	m := map[string]int{"a": 1}

	got := Acumular(m, "b", 5)

	if got == nil {
		t.Fatal("Acumular devolveu nil")
	}
	if got["b"] != 5 {
		t.Fatalf(`got["b"] = %d, quero 5`, got["b"])
	}
	if got["a"] != 1 {
		t.Fatalf(`got["a"] = %d, quero 1: as outras chaves não deveriam mudar`, got["a"])
	}
}

func TestAcumularNaoMexeNasOutrasChaves(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	got := Acumular(m, "a", 10)

	if len(got) != 2 {
		t.Fatalf("Acumular devolveu %v, quero 2 chaves", got)
	}
}

// Aqui a diferença entre o nil de map e o nil de slice aparece: escrever num
// map nil é pânico, não um mapa criado por conveniência.
func TestAcumularComMapNil(t *testing.T) {
	got := Acumular(nil, "a", 7)

	if got == nil {
		t.Fatal("Acumular(nil, ...) devolveu nil: quero um map utilizável")
	}
	if got["a"] != 7 {
		t.Fatalf(`got["a"] = %d, quero 7`, got["a"])
	}
	if len(got) != 1 {
		t.Fatalf("Acumular(nil, ...) devolveu %v, quero exatamente uma chave", got)
	}
}

// Um map nil declarado com var é o mesmo nil que o literal — e o resultado
// devolvido precisa aceitar escrita de quem chamou.
func TestResultadoAceitaEscritaDepois(t *testing.T) {
	var m map[string]int

	got := Acumular(m, "a", 1)
	if got == nil {
		t.Fatal("Acumular devolveu nil")
	}

	got["b"] = 2 // entraria em pânico se o retorno ainda fosse nil

	if len(got) != 2 {
		t.Fatalf("depois de escrever, o map é %v, quero 2 chaves", got)
	}
}
