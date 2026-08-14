package main

import "testing"

func TestFaixaClassificaCadaIntervalo(t *testing.T) {
	casos := []struct {
		n    int
		want string
	}{
		{-42, "baixo"},
		{0, "baixo"},
		{9, "baixo"},
		{10, "medio"},
		{50, "medio"},
		{99, "medio"},
		{100, "alto"},
		{1000, "alto"},
	}

	for _, c := range casos {
		if got := Faixa(c.n); got != c.want {
			t.Errorf("Faixa(%d) = %q, quero %q", c.n, got, c.want)
		}
	}
}
