package main

import "fmt"

// Dobrar devolve um slice novo com cada elemento de xs multiplicado por 2.
//
// O make cria um destino próprio: a capacidade já reservada evita realocações
// no append, e o array é independente do de xs, que é o que os testes cobram.
func Dobrar(xs []int) []int {
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		out = append(out, x*2)
	}
	return out
}

func main() {
	original := []int{1, 2, 3}
	fmt.Println("original:", original)
	fmt.Println("dobrado: ", Dobrar(original))
}
