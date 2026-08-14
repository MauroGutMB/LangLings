package main

import "fmt"

// Drenar deve recolher tudo o que já estiver pronto em ch, sem nunca bloquear,
// e devolver os valores na ordem em que saíram.
//
// TODO: o esqueleto abaixo nunca lê nada do canal.
func Drenar(ch <-chan int) []int {
	return []int{}
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	fmt.Println(Drenar(ch)) // [1 2]
	fmt.Println(Drenar(ch)) // []  — e sem travar, apesar do canal estar aberto
}
