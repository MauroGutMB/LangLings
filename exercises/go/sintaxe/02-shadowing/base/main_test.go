package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestSomarTodasValidas(t *testing.T) {
	total, err := Somar([]string{"1", "2", "3"})
	if err != nil {
		t.Fatalf("Somar devolveu erro %v, não queria nenhum", err)
	}
	if total != 6 {
		t.Fatalf("total = %d, quero 6", total)
	}
}

func TestSomarIgnoraAsInvalidasNaSoma(t *testing.T) {
	total, _ := Somar([]string{"1", "dois", "3"})

	if total != 4 {
		t.Fatalf("total = %d, quero 4: as entradas inválidas não entram na soma", total)
	}
}

// Este é o teste que pega a armadilha: um erro guardado num := dentro do laço
// morre no fim da volta e o retorno sai nil.
func TestSomarLembraDoPrimeiroErro(t *testing.T) {
	_, err := Somar([]string{"1", "dois", "3"})

	if err == nil {
		t.Fatal("Somar devolveu erro nil, mas \"dois\" não é um número")
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Fatalf("Somar devolveu %v, quero o erro do strconv (ErrSyntax)", err)
	}
}

func TestSomarDevolveOPrimeiroErroENaoOUltimo(t *testing.T) {
	_, err := Somar([]string{"dois", "3", "1e999999999"})

	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Fatalf("Somar devolveu %v, quero um *strconv.NumError", err)
	}
	if numErr.Num != "dois" {
		t.Fatalf("o erro devolvido é o de %q, quero o da primeira entrada inválida (%q)", numErr.Num, "dois")
	}
}

func TestSomarTodasInvalidas(t *testing.T) {
	total, err := Somar([]string{"a", "b"})
	if err == nil {
		t.Fatal("Somar devolveu erro nil com todas as entradas inválidas")
	}
	if total != 0 {
		t.Fatalf("total = %d, quero 0", total)
	}
}

func TestSomarSemEntradas(t *testing.T) {
	total, err := Somar(nil)
	if err != nil {
		t.Fatalf("Somar(nil) devolveu erro %v, não queria nenhum", err)
	}
	if total != 0 {
		t.Fatalf("total = %d, quero 0", total)
	}
}

func TestSomarAceitaNegativos(t *testing.T) {
	total, err := Somar([]string{"-5", "10"})
	if err != nil {
		t.Fatalf("Somar devolveu erro %v", err)
	}
	if total != 5 {
		t.Fatalf("total = %d, quero 5", total)
	}
}
