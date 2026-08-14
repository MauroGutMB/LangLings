package main

import (
	"errors"
	"fmt"
	"sort"
)

// Item é uma linha do estoque. Preco está em centavos: dinheiro em float64
// acumula erro de arredondamento, e inteiro não.
type Item struct {
	Nome      string
	Categoria string
	Preco     int
	Qtd       int
}

// Valor é o quanto este item representa no estoque.
func (i Item) Valor() int { return i.Preco * i.Qtd }

var (
	ErrSemItens       = errors.New("nenhum item")
	ErrCategoriaVazia = errors.New("item sem categoria")
)

func exemplos() {
	itens := []Item{
		{Nome: "caneta", Categoria: "papelaria", Preco: 250, Qtd: 4},
		{Nome: "caderno", Categoria: "papelaria", Preco: 1200, Qtd: 2},
		{Nome: "café", Categoria: "cozinha", Preco: 1850, Qtd: 3},
	}

	// Struct + método: cada item sabe calcular o próprio valor.
	fmt.Println(itens[0].Valor()) // 1000

	// Slice + laço + acumulador: o total geral.
	total := 0
	for _, it := range itens {
		total += it.Valor()
	}
	fmt.Println(total) // 9000

	// Map como índice: agrupa os itens por categoria. A chave ausente lê como
	// slice nil, e append em nil já funciona.
	porCategoria := map[string][]Item{}
	for _, it := range itens {
		porCategoria[it.Categoria] = append(porCategoria[it.Categoria], it)
	}
	fmt.Println(len(porCategoria["papelaria"])) // 2

	// A ordem do range sobre um map é aleatória; para imprimir estável,
	// colete as chaves e ordene.
	chaves := make([]string, 0, len(porCategoria))
	for c := range porCategoria {
		chaves = append(chaves, c)
	}
	sort.Strings(chaves)
	fmt.Println(chaves) // [cozinha papelaria]

	// Erro como valor: a validação devolve, quem chama decide.
	if err := validar(nil); err != nil {
		fmt.Println(err, errors.Is(err, ErrSemItens)) // validar estoque: nenhum item true
	}
}

func validar(itens []Item) error {
	if len(itens) == 0 {
		return fmt.Errorf("validar estoque: %w", ErrSemItens)
	}
	return nil
}

// SUA VEZ
//
// Devolva o valor total (Preco * Qtd) por categoria. Slice vazio devolve nil e
// um erro que embrulha ErrSemItens; item com Categoria vazia devolve nil e um
// erro que embrulha ErrCategoriaVazia.
func Resumir(itens []Item) (map[string]int, error) {
	return nil, nil // <- troque isto
}

func main() {
	exemplos()
}
