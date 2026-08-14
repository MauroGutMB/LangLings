package main

import "testing"

func iguais(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParesFiltraMantendoAOrdem(t *testing.T) {
	got := Pares([]int{1, 2, 3, 4, 6, 7})
	if want := []int{2, 4, 6}; !iguais(got, want) {
		t.Fatalf("Pares([1 2 3 4 6 7]) = %v, quero %v", got, want)
	}
}

func TestParesIncluiZeroENegativos(t *testing.T) {
	got := Pares([]int{0, -2, -3})
	if want := []int{0, -2}; !iguais(got, want) {
		t.Fatalf("Pares([0 -2 -3]) = %v, quero %v", got, want)
	}
}

func TestParesSemNenhumPar(t *testing.T) {
	if got := Pares([]int{1, 3, 5}); len(got) != 0 {
		t.Fatalf("Pares([1 3 5]) = %v, quero comprimento 0", got)
	}
}

func TestParesComSliceVazio(t *testing.T) {
	if got := Pares([]int{}); len(got) != 0 {
		t.Fatalf("Pares([]) = %v, quero comprimento 0", got)
	}
}

func TestParesComNil(t *testing.T) {
	if got := Pares(nil); len(got) != 0 {
		t.Fatalf("Pares(nil) = %v, quero comprimento 0", got)
	}
}

func TestParesNaoModificaAEntrada(t *testing.T) {
	entrada := []int{1, 2, 3, 4}

	Pares(entrada)

	if want := []int{1, 2, 3, 4}; !iguais(entrada, want) {
		t.Fatalf("Pares modificou a entrada: agora ela é %v", entrada)
	}
}
