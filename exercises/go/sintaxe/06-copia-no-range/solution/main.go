package main

import "fmt"

// Produto é uma linha de estoque.
type Produto struct {
	Nome string
	Qtd  int
}

// Repor soma n à quantidade de todos os produtos de ps.
//
// O laço percorre índices, não valores: a variável de valor do range é uma
// cópia do elemento, e escrever nela some no fim da volta. Indexando, a escrita
// acontece no array que ps aponta — o mesmo de quem chamou.
func Repor(ps []Produto, n int) {
	for i := range ps {
		ps[i].Qtd += n
	}
}

func main() {
	ps := []Produto{{Nome: "caneta", Qtd: 1}, {Nome: "caderno", Qtd: 2}}

	Repor(ps, 10)
	fmt.Println(ps)
}
