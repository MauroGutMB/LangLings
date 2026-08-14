package main

import "fmt"

// Enfileirar devolve um canal por onde os valores de xs saem, em ordem, já
// fechado no fim.
//
// A goroutine é o que permite a Enfileirar retornar imediatamente: num canal
// sem buffer o envio fica parado esperando um receptor, e o receptor só
// aparece depois que quem chamou tiver o canal em mãos — o que nunca
// aconteceria se os envios rodassem aqui mesmo. Dar buffer de len(xs) também
// resolveria, ao custo de materializar tudo na memória de uma vez; com a
// goroutine, o canal continua funcionando como fluxo.
func Enfileirar(xs []int) <-chan int {
	ch := make(chan int)

	go func() {
		// O close é responsabilidade de quem envia, e vem depois do último
		// envio: é ele que faz o range de quem recebe terminar.
		defer close(ch)

		for _, x := range xs {
			ch <- x
		}
	}()

	return ch
}

func main() {
	for v := range Enfileirar([]int{1, 2, 3}) {
		fmt.Println(v)
	}
}
