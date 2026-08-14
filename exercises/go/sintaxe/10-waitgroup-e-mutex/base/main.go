package main

import "fmt"

// Processar deve chamar fn(x) para cada x de xs, cada chamada em sua própria
// goroutine, e devolver os resultados na ordem de xs junto com quantos deles
// são negativos.
//
// TODO: implemente. O esqueleto abaixo não chama fn nenhuma vez.
func Processar(xs []int, fn func(int) int) ([]int, int) {
	return nil, 0
}

func main() {
	res, negativos := Processar([]int{1, -2, 3}, func(x int) int { return x * 10 })
	fmt.Println(res, negativos)
}
