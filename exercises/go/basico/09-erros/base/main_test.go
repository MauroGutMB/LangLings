package main

import (
	"errors"
	"strings"
	"testing"
)

func TestDividirCasoNormal(t *testing.T) {
	got, err := Dividir(7, 2)
	if err != nil {
		t.Fatalf("Dividir(7, 2) devolveu erro %v, não queria nenhum", err)
	}
	if got != 3 {
		t.Fatalf("Dividir(7, 2) = %d, quero 3", got)
	}
}

func TestDividirPorZeroDevolveErro(t *testing.T) {
	got, err := Dividir(1, 0)
	if err == nil {
		t.Fatal("Dividir(1, 0) não devolveu erro")
	}
	if got != 0 {
		t.Errorf("Dividir(1, 0) devolveu %d, quero 0 junto com o erro", got)
	}
}

// É este teste que cobra o %w: com %v o texto sairia igual e o errors.Is
// devolveria false.
func TestDividirPorZeroEmbrulhaASentinela(t *testing.T) {
	_, err := Dividir(1, 0)

	if !errors.Is(err, ErrDivisaoPorZero) {
		t.Fatalf("errors.Is(%v, ErrDivisaoPorZero) é false: a causa não foi preservada", err)
	}
}

// O erro precisa dizer mais do que a sentinela sozinha diria: devolver a
// própria ErrDivisaoPorZero passaria no teste acima e não daria contexto nenhum.
func TestDividirPorZeroAcrescentaContexto(t *testing.T) {
	_, err := Dividir(1, 0)

	if strings.TrimSpace(err.Error()) == ErrDivisaoPorZero.Error() {
		t.Fatalf("a mensagem é só %q: acrescente contexto ao embrulhar", err)
	}
}

func TestDividirNegativos(t *testing.T) {
	got, err := Dividir(-9, 3)
	if err != nil {
		t.Fatalf("Dividir(-9, 3) devolveu erro %v", err)
	}
	if got != -3 {
		t.Fatalf("Dividir(-9, 3) = %d, quero -3", got)
	}
}
