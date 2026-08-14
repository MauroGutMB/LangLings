package main

import "fmt"

// Somar deve devolver a soma das entradas que são números válidos e o erro da
// primeira entrada inválida (nil se não houver nenhuma).
//
// TODO: implemente. O esqueleto abaixo não soma nada e nunca falha.
func Somar(entradas []string) (int, error) {
	return 0, nil
}

func main() {
	fmt.Println(Somar([]string{"1", "2", "3"}))
	fmt.Println(Somar([]string{"1", "dois", "3", "quatro"}))
}
