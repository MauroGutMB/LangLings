package main

import "fmt"

// Dobrar deve devolver um slice NOVO com cada elemento de xs multiplicado
// por 2, sem modificar xs e sem compartilhar o array por baixo.
//
// TODO: a implementação abaixo devolve a própria entrada. Corrija.
func Dobrar(xs []int) []int {
	return xs
}

func main() {
	original := []int{1, 2, 3}
	fmt.Println("original:", original)
	fmt.Println("dobrado: ", Dobrar(original))
}
