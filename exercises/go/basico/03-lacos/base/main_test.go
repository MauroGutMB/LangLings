package main

import "testing"

func TestSomaDeVariosElementos(t *testing.T) {
	if got := Soma([]int{1, 2, 3, 4}); got != 10 {
		t.Fatalf("Soma([1 2 3 4]) = %d, quero 10", got)
	}
}

func TestSomaComNegativos(t *testing.T) {
	if got := Soma([]int{5, -3, -2}); got != 0 {
		t.Fatalf("Soma([5 -3 -2]) = %d, quero 0", got)
	}
}

func TestSomaDeUmElemento(t *testing.T) {
	if got := Soma([]int{7}); got != 7 {
		t.Fatalf("Soma([7]) = %d, quero 7", got)
	}
}

func TestSomaDeSliceVazio(t *testing.T) {
	if got := Soma([]int{}); got != 0 {
		t.Fatalf("Soma([]) = %d, quero 0", got)
	}
}

func TestSomaDeNil(t *testing.T) {
	if got := Soma(nil); got != 0 {
		t.Fatalf("Soma(nil) = %d, quero 0", got)
	}
}
