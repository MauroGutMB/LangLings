package main

import (
	"fmt"
	"strconv"
)

// Somar devolve o total das entradas numéricas e o erro da primeira inválida.
//
// primeiroErro é declarado fora do laço e atribuído com = lá dentro. Escrever
// n, primeiroErro := strconv.Atoi(e) compilaria e não faria o que parece: o
// := declararia um primeiroErro novo, válido só naquela volta.
func Somar(entradas []string) (int, error) {
	total := 0
	var primeiroErro error

	for _, e := range entradas {
		n, err := strconv.Atoi(e)
		if err != nil {
			if primeiroErro == nil {
				primeiroErro = err
			}
			continue
		}
		total += n
	}

	return total, primeiroErro
}

func main() {
	fmt.Println(Somar([]string{"1", "2", "3"}))
	fmt.Println(Somar([]string{"1", "dois", "3", "quatro"}))
}
