package main

import "testing"

func TestDobrarMultiplicaCadaElemento(t *testing.T) {
	got := Dobrar([]int{1, 2, 3})
	want := []int{2, 4, 6}

	if len(got) != len(want) {
		t.Fatalf("Dobrar([1 2 3]) devolveu %v, quero %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Dobrar([1 2 3]) devolveu %v, quero %v", got, want)
		}
	}
}

func TestDobrarNaoModificaAEntrada(t *testing.T) {
	original := []int{1, 2, 3}

	Dobrar(original)

	for i, want := range []int{1, 2, 3} {
		if original[i] != want {
			t.Fatalf("Dobrar modificou a entrada: agora ela é %v", original)
		}
	}
}

// Este é o teste que pega a armadilha: se o resultado reaproveitar o array da
// entrada, escrever num deles altera o outro.
func TestDobrarNaoCompartilhaOArrayComAEntrada(t *testing.T) {
	original := []int{1, 2, 3}

	got := Dobrar(original)
	if len(got) == 0 {
		t.Fatal("Dobrar devolveu um slice vazio")
	}
	got[0] = 999

	if original[0] != 1 {
		t.Fatalf("mexer no resultado alterou a entrada (%v): os dois compartilham o mesmo array", original)
	}
}

func TestDobrarComSliceVazio(t *testing.T) {
	if got := Dobrar([]int{}); len(got) != 0 {
		t.Fatalf("Dobrar([]) devolveu %v, quero um slice vazio", got)
	}
}

func TestDobrarComNil(t *testing.T) {
	if got := Dobrar(nil); len(got) != 0 {
		t.Fatalf("Dobrar(nil) devolveu %v, quero um slice vazio", got)
	}
}
