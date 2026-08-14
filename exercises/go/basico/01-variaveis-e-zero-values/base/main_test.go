package main

import "testing"

func TestMediaDeInteirosPares(t *testing.T) {
	if got := Media(2, 4); got != 3 {
		t.Fatalf("Media(2, 4) = %v, quero 3", got)
	}
}

// A fração é o ponto do exercício: uma média que devolve 1 em vez de 1.5
// dividiu antes de converter.
func TestMediaMantemAFracao(t *testing.T) {
	if got := Media(1, 2); got != 1.5 {
		t.Fatalf("Media(1, 2) = %v, quero 1.5 (a parte fracionária foi perdida?)", got)
	}
}

func TestMediaComNegativos(t *testing.T) {
	if got := Media(-3, -4); got != -3.5 {
		t.Fatalf("Media(-3, -4) = %v, quero -3.5", got)
	}
}

func TestMediaDeZeros(t *testing.T) {
	if got := Media(0, 0); got != 0 {
		t.Fatalf("Media(0, 0) = %v, quero 0", got)
	}
}
