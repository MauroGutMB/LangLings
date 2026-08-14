package main

import "fmt"

// Drenar recolhe tudo o que já estiver pronto em ch, sem bloquear.
//
// O default é o que torna o select uma tentativa em vez de uma espera: se
// nenhum case estiver pronto na hora, ele roda em vez de o select dormir. É
// também o que salva o caso do canal nil, cujo case nunca fica pronto.
//
// O ok existe porque um canal fechado está SEMPRE pronto para receber: sem
// ele, o case entregaria o zero value indefinidamente e o laço não pararia.
func Drenar(ch <-chan int) []int {
	out := []int{}

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		default:
			return out
		}
	}
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	fmt.Println(Drenar(ch)) // [1 2]
	fmt.Println(Drenar(ch)) // []  — e sem travar, apesar do canal estar aberto
}
