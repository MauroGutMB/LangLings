package main

import (
	"errors"
	"fmt"
)

// ErrVazio é devolvido (embrulhado) quando não há nada para somar.
var ErrVazio = errors.New("nenhum elemento")

// Processar soma xs, registra o total e devolve-o.
//
// O defer recebe uma função anônima em vez de registrar(total) direto: os
// argumentos de um defer são avaliados na linha do defer, e ali total ainda
// vale 0. Dentro do closure a leitura acontece na hora do retorno — e o
// retorno nomeado garante que a variável lida seja a que está sendo devolvida.
func Processar(xs []int, registrar func(int)) (total int, err error) {
	defer func() { registrar(total) }()

	if len(xs) == 0 {
		return 0, fmt.Errorf("processar: %w", ErrVazio)
	}

	for _, x := range xs {
		total += x
	}
	return total, nil
}

func main() {
	log := func(n int) { fmt.Println("registrado:", n) }

	fmt.Println(Processar([]int{1, 2, 3}, log))
	fmt.Println(Processar(nil, log))
}
