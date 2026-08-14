package main

import "fmt"

// Enfileirar deve devolver um canal por onde os valores de xs saem, em ordem,
// e que é fechado depois do último.
//
// TODO: o esqueleto abaixo devolve um canal já fechado e vazio. Faça os
// valores passarem por ele.
func Enfileirar(xs []int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

func main() {
	for v := range Enfileirar([]int{1, 2, 3}) {
		fmt.Println(v)
	}
}
