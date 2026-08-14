package main

import "fmt"

// Acumular deve somar n ao valor de k em m e devolver o map resultante.
// m pode chegar nil, e ainda assim o retorno precisa ser um map utilizável.
//
// TODO: a implementação abaixo devolve o map sem tocar nele. Corrija.
func Acumular(m map[string]int, k string, n int) map[string]int {
	return m
}

func main() {
	// Repare no contraste: o slice nil aceita append e vira um slice de
	// verdade; o map nil só aceita leitura.
	var xs []int
	xs = append(xs, 1)
	fmt.Println("slice nil + append:", xs)

	var m map[string]int
	fmt.Println("map nil, leitura:", m["a"], len(m))

	fmt.Println("acumulado:", Acumular(nil, "a", 3))
}
