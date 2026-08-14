package main

import "fmt"

// Produto é uma linha de estoque.
type Produto struct {
	Nome string
	Qtd  int
}

// Repor deve somar n à quantidade de todos os produtos de ps.
//
// TODO: o laço abaixo roda, soma e não deixa rastro. Descubra em que ele está
// somando e corrija.
func Repor(ps []Produto, n int) {
	for _, p := range ps {
		p.Qtd += n
	}
}

func main() {
	ps := []Produto{{Nome: "caneta", Qtd: 1}, {Nome: "caderno", Qtd: 2}}

	Repor(ps, 10)
	fmt.Println(ps)
}
