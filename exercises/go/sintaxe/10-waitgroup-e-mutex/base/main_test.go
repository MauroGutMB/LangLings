package main

import (
	"sync/atomic"
	"testing"
	"time"
)

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

// lento devolve o dobro depois de uma pausa: uma função instantânea esconderia
// tanto a falta de espera quanto a corrida no contador.
func lento(x int) int {
	time.Sleep(3 * time.Millisecond)
	return x * 2
}

func TestProcessarMantemAOrdemDaEntrada(t *testing.T) {
	got, _ := Processar([]int{1, 2, 3, 4}, lento)

	if want := []int{2, 4, 6, 8}; !iguais(got, want) {
		t.Fatalf("Processar devolveu %v, quero %v", got, want)
	}
}

// Uma goroutine por elemento que ninguém espera deixa o slice zerado: este
// teste é o que cobra a espera.
func TestProcessarEsperaTodasAsGoroutines(t *testing.T) {
	xs := make([]int, 200)
	for i := range xs {
		xs[i] = i + 1
	}

	got, _ := Processar(xs, lento)

	if len(got) != len(xs) {
		t.Fatalf("Processar devolveu %d resultados, quero %d", len(got), len(xs))
	}
	for i, x := range xs {
		if got[i] != x*2 {
			t.Fatalf("resultado %d = %d, quero %d: alguma goroutine não tinha terminado", i, got[i], x*2)
		}
	}
}

func TestProcessarContaOsNegativos(t *testing.T) {
	_, negativos := Processar([]int{1, -2, 3, -4, -5}, lento)

	if negativos != 3 {
		t.Fatalf("negativos = %d, quero 3", negativos)
	}
}

// Muitos incrementos simultâneos no mesmo contador: sem proteção, alguns se
// perdem e a conta não fecha.
func TestProcessarContaCertoComMuitasGoroutines(t *testing.T) {
	xs := make([]int, 500)
	for i := range xs {
		xs[i] = -1
	}

	_, negativos := Processar(xs, lento)

	if negativos != 500 {
		t.Fatalf("negativos = %d, quero 500: incrementos concorrentes se perderam", negativos)
	}
}

func TestProcessarChamaFnUmaVezPorElemento(t *testing.T) {
	var chamadas int64

	xs := []int{1, 2, 3, 4, 5}
	Processar(xs, func(x int) int {
		atomic.AddInt64(&chamadas, 1)
		return x
	})

	if got := atomic.LoadInt64(&chamadas); got != int64(len(xs)) {
		t.Fatalf("fn foi chamada %d vez(es), quero %d", got, len(xs))
	}
}

func TestProcessarComSliceVazioENil(t *testing.T) {
	got, negativos := Processar(nil, lento)
	if len(got) != 0 || negativos != 0 {
		t.Fatalf("Processar(nil) devolveu %v, %d, quero comprimento 0 e 0", got, negativos)
	}

	got, negativos = Processar([]int{}, lento)
	if len(got) != 0 || negativos != 0 {
		t.Fatalf("Processar([]) devolveu %v, %d, quero comprimento 0 e 0", got, negativos)
	}
}

func TestProcessarZeroNaoEhNegativo(t *testing.T) {
	_, negativos := Processar([]int{0, 0}, lento)

	if negativos != 0 {
		t.Fatalf("negativos = %d, quero 0: zero não é negativo", negativos)
	}
}
