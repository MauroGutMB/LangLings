package main

import "testing"

func TestMaiorEncontraOMaximo(t *testing.T) {
	got, ok := Maior(3, 9, 4)
	if !ok || got != 9 {
		t.Fatalf("Maior(3, 9, 4) = %d, %v, quero 9, true", got, ok)
	}
}

func TestMaiorComUmValorSo(t *testing.T) {
	got, ok := Maior(42)
	if !ok || got != 42 {
		t.Fatalf("Maior(42) = %d, %v, quero 42, true", got, ok)
	}
}

// Começar o candidato em 0 em vez de no primeiro elemento passa nos casos
// acima e quebra aqui.
func TestMaiorComTodosNegativos(t *testing.T) {
	got, ok := Maior(-5, -1, -9)
	if !ok || got != -1 {
		t.Fatalf("Maior(-5, -1, -9) = %d, %v, quero -1, true", got, ok)
	}
}

func TestMaiorSemArgumentos(t *testing.T) {
	got, ok := Maior()
	if ok || got != 0 {
		t.Fatalf("Maior() = %d, %v, quero 0, false", got, ok)
	}
}

func TestMaiorAceitaSliceEspalhado(t *testing.T) {
	xs := []int{1, 8, 2}
	got, ok := Maior(xs...)
	if !ok || got != 8 {
		t.Fatalf("Maior(xs...) = %d, %v, quero 8, true", got, ok)
	}
}
