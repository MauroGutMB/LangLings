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

func TestDrenarRecolheOQueEstaNoBuffer(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3

	got := Drenar(ch)

	if want := []int{1, 2, 3}; !iguais(got, want) {
		t.Fatalf("Drenar devolveu %v, quero %v", got, want)
	}
}

// O canal continua ABERTO aqui: uma leitura que espera trava para sempre.
func TestDrenarNaoBloqueiaEmCanalAbertoEVazio(t *testing.T) {
	ch := make(chan int, 2)

	if got := Drenar(ch); len(got) != 0 {
		t.Fatalf("Drenar devolveu %v, quero comprimento 0", got)
	}
}

func TestDrenarPodeSerChamadaDeNovo(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 1

	if got := Drenar(ch); !iguais(got, []int{1}) {
		t.Fatalf("primeira chamada devolveu %v, quero [1]", got)
	}
	if got := Drenar(ch); len(got) != 0 {
		t.Fatalf("segunda chamada devolveu %v, quero comprimento 0", got)
	}

	ch <- 9
	if got := Drenar(ch); !iguais(got, []int{9}) {
		t.Fatalf("terceira chamada devolveu %v, quero [9]", got)
	}
}

// Um canal fechado está sempre pronto para receber: sem distinguir isso de um
// valor de verdade, o laço nunca termina.
func TestDrenarComCanalFechadoComValores(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 4
	ch <- 5
	close(ch)

	got := Drenar(ch)

	if want := []int{4, 5}; !iguais(got, want) {
		t.Fatalf("Drenar devolveu %v, quero %v", got, want)
	}
}

func TestDrenarComCanalFechadoEVazio(t *testing.T) {
	ch := make(chan int)
	close(ch)

	if got := Drenar(ch); len(got) != 0 {
		t.Fatalf("Drenar devolveu %v, quero comprimento 0", got)
	}
}

// Receber de um canal nil bloqueia para sempre — a não ser dentro de um select
// que tenha para onde ir.
func TestDrenarComCanalNil(t *testing.T) {
	var ch chan int

	if got := Drenar(ch); len(got) != 0 {
		t.Fatalf("Drenar(nil) devolveu %v, quero comprimento 0", got)
	}
}

func TestDrenarNuncaDevolveNil(t *testing.T) {
	ch := make(chan int, 1)

	if got := Drenar(ch); got == nil {
		t.Fatal("Drenar devolveu nil, quero um slice de comprimento 0")
	}
}
