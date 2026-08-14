package main

import "fmt"

// for é o único laço de Go. As outras linguagens têm while, do-while e
// foreach; aqui tudo isso é escrito com for em formas diferentes.
func exemplos() {
	// 1) completa: init; condição; pós.
	soma := 0
	for i := 1; i <= 3; i++ {
		soma += i
	}
	fmt.Println(soma) // 6

	// 2) só a condição — é o "while" das outras linguagens.
	n := 8
	passos := 0
	for n > 1 {
		n /= 2
		passos++
	}
	fmt.Println(passos) // 3

	// 3) sem nada: laço infinito, do qual se sai com break.
	i := 0
	for {
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println(i) // 5

	// range percorre slices, arrays, maps, strings e canais. Num slice ele
	// entrega índice e valor; _ descarta o que não interessa.
	for idx, v := range []string{"a", "b"} {
		fmt.Println(idx, v) // 0 a  /  1 b
	}
	for _, v := range []int{10, 20} {
		soma += v
	}
	fmt.Println(soma) // 36

	// range sobre um int percorre 0..n-1. continue pula para a volta seguinte.
	pares := 0
	for k := range 10 {
		if k%2 != 0 {
			continue
		}
		pares++
	}
	fmt.Println(pares) // 5
}

// SUA VEZ
//
// Devolva a soma dos elementos de xs. Slice vazio ou nil soma 0.
func Soma(xs []int) int {
	return -1 // <- troque isto
}

func main() {
	exemplos()
}
