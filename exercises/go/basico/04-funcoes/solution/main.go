package main

import (
	"errors"
	"fmt"
)

// Uma função pode devolver mais de um valor. É por isso que Go não precisa de
// exceções: o erro é só mais um retorno, e quem chama é obrigado a lidar com
// ele ou a descartá-lo explicitamente com _.
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// Parâmetros vizinhos do mesmo tipo compartilham a anotação (xs ...int, não
// xs ...int, ys ...int). O ... na declaração torna a função variádica: xs
// chega aqui dentro como um slice.
func somar(xs ...int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func exemplos() {
	q, err := dividir(7, 2)
	fmt.Println(q, err) // 3 <nil>  — divisão inteira trunca

	if _, err := dividir(1, 0); err != nil {
		fmt.Println("erro:", err) // erro: divisão por zero
	}

	fmt.Println(somar())        // 0  — variádica aceita zero argumentos
	fmt.Println(somar(1, 2, 3)) // 6

	// O mesmo ... espalha um slice já existente numa chamada variádica.
	xs := []int{4, 5}
	fmt.Println(somar(xs...)) // 9

	// Funções são valores: cabem numa variável e podem ser passadas adiante.
	op := somar
	fmt.Println(op(10, 10)) // 20
}

// Maior devolve o maior dos números recebidos.
//
// O bool existe porque não há resposta correta para a lista vazia: devolver 0
// seria mentir para quem passou só negativos.
func Maior(xs ...int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}

	maior := xs[0]
	for _, x := range xs[1:] {
		if x > maior {
			maior = x
		}
	}
	return maior, true
}

func main() {
	exemplos()
}
