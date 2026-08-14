package main

import "testing"

func TestReporAlteraTodosOsProdutos(t *testing.T) {
	ps := []Produto{{Nome: "caneta", Qtd: 1}, {Nome: "caderno", Qtd: 2}}

	Repor(ps, 10)

	if ps[0].Qtd != 11 || ps[1].Qtd != 12 {
		t.Fatalf("depois de Repor, ps = %v, quero [{caneta 11} {caderno 12}]", ps)
	}
}

func TestReporNaoMexeNosOutrosCampos(t *testing.T) {
	ps := []Produto{{Nome: "caneta", Qtd: 1}}

	Repor(ps, 5)

	if ps[0].Nome != "caneta" {
		t.Fatalf("Repor alterou o nome: %q", ps[0].Nome)
	}
}

func TestReporComZeroNaoMuda(t *testing.T) {
	ps := []Produto{{Nome: "caneta", Qtd: 7}}

	Repor(ps, 0)

	if ps[0].Qtd != 7 {
		t.Fatalf("Repor(ps, 0) deixou Qtd = %d, quero 7", ps[0].Qtd)
	}
}

func TestReporAceitaNegativo(t *testing.T) {
	ps := []Produto{{Nome: "caneta", Qtd: 7}}

	Repor(ps, -3)

	if ps[0].Qtd != 4 {
		t.Fatalf("Repor(ps, -3) deixou Qtd = %d, quero 4", ps[0].Qtd)
	}
}

func TestReporComSliceVazioENil(t *testing.T) {
	Repor([]Produto{}, 1)
	Repor(nil, 1)
}

// O slice é uma vista sobre o array de quem chamou: alterar através dele tem
// que ser visível na variável original, sem retorno nenhum.
func TestReporEnxergaAlemDoSliceRecebido(t *testing.T) {
	todos := []Produto{{Nome: "a", Qtd: 0}, {Nome: "b", Qtd: 0}, {Nome: "c", Qtd: 0}}

	Repor(todos[1:], 4)

	if todos[0].Qtd != 0 {
		t.Fatalf("Repor mexeu fora do slice recebido: %v", todos)
	}
	if todos[1].Qtd != 4 || todos[2].Qtd != 4 {
		t.Fatalf("depois de Repor(todos[1:], 4), todos = %v", todos)
	}
}
