package main

import "testing"

func TestIncrementarAlteraOContador(t *testing.T) {
	c := Contador{N: 3}

	c.Incrementar()

	if c.N != 4 {
		t.Fatalf("depois de Incrementar, N = %d, quero 4", c.N)
	}
}

func TestIncrementarAcumula(t *testing.T) {
	var c Contador

	for range 5 {
		c.Incrementar()
	}

	if c.N != 5 {
		t.Fatalf("depois de 5 chamadas, N = %d, quero 5", c.N)
	}
}

func TestIncrementarAtravesDeUmPonteiro(t *testing.T) {
	c := &Contador{N: 10}

	c.Incrementar()

	if c.N != 11 {
		t.Fatalf("depois de Incrementar, N = %d, quero 11", c.N)
	}
}

func TestDobroDevolveUmContadorNovo(t *testing.T) {
	c := Contador{N: 5}

	d := c.Dobro()

	if d.N != 10 {
		t.Fatalf("Dobro().N = %d, quero 10", d.N)
	}
}

// O contrário do primeiro teste: aqui o método não pode encostar no original.
func TestDobroNaoAlteraOOriginal(t *testing.T) {
	c := Contador{N: 5}

	c.Dobro()

	if c.N != 5 {
		t.Fatalf("depois de Dobro, o original vale %d, quero 5", c.N)
	}
}

func TestDobroPodeSerChamadoVariasVezes(t *testing.T) {
	c := Contador{N: 2}

	a := c.Dobro()
	b := c.Dobro()

	if a.N != 4 || b.N != 4 {
		t.Fatalf("Dobro() devolveu %d e depois %d, quero 4 nas duas vezes", a.N, b.N)
	}
}
