package main

import (
	"fmt"
	"sync"
)

// Processar chama fn(x) para cada x de xs em paralelo.
//
// Duas estratégias diferentes convivem aqui de propósito. Os resultados não
// precisam de trava: o slice já nasce com o tamanho final e cada goroutine é
// dona de um índice, então não há duas escritas no mesmo endereço. O contador
// precisa, porque todas escrevem na mesma variável e negativos++ é ler, somar
// e gravar — três passos que podem se intercalar e perder incrementos.
//
// O Wait é o que transforma "disparei tudo" em "tudo terminou": sem ele,
// Processar devolveria o slice ainda zerado.
func Processar(xs []int, fn func(int) int) ([]int, int) {
	resultados := make([]int, len(xs))

	var mu sync.Mutex
	negativos := 0

	var wg sync.WaitGroup
	for i, x := range xs {
		wg.Add(1)
		go func() {
			defer wg.Done()

			r := fn(x)
			resultados[i] = r

			if r < 0 {
				mu.Lock()
				negativos++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	return resultados, negativos
}

func main() {
	res, negativos := Processar([]int{1, -2, 3}, func(x int) int { return x * 10 })
	fmt.Println(res, negativos)
}
