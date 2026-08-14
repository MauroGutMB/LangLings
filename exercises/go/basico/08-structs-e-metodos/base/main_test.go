package main

import "testing"

func TestPerimetro(t *testing.T) {
	r := Retangulo{Largura: 3, Altura: 2}

	if got := r.Perimetro(); got != 10 {
		t.Fatalf("Retangulo{3, 2}.Perimetro() = %v, quero 10", got)
	}
}

func TestPerimetroDoZeroValue(t *testing.T) {
	var r Retangulo

	if got := r.Perimetro(); got != 0 {
		t.Fatalf("Retangulo{}.Perimetro() = %v, quero 0", got)
	}
}

func TestPerimetroNaoAlteraORetangulo(t *testing.T) {
	r := Retangulo{Largura: 1.5, Altura: 2.5}

	r.Perimetro()

	if r.Largura != 1.5 || r.Altura != 2.5 {
		t.Fatalf("Perimetro alterou o retângulo: agora ele é %v", r)
	}
}

// Perimetro precisa continuar acessível pelo tipo que embute Retangulo.
func TestPerimetroPeloTipoEmbutido(t *testing.T) {
	c := Caixa{Retangulo: Retangulo{Largura: 4, Altura: 1}, Profundidade: 9}

	if got := c.Perimetro(); got != 10 {
		t.Fatalf("Caixa{...}.Perimetro() = %v, quero 10", got)
	}
}
