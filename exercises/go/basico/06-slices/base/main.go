package main

import "fmt"

// Um slice não é um array: é uma vista sobre um array, feita de ponteiro,
// comprimento (len) e capacidade (cap).
func exemplos() {
	xs := []int{1, 2, 3}
	fmt.Println(len(xs), cap(xs)) // 3 3

	// make reserva capacidade sem criar elementos: len 0, cap 4. Isso evita
	// que os appends seguintes realoquem o array por baixo.
	ys := make([]int, 0, 4)
	ys = append(ys, 1)
	fmt.Println(len(ys), cap(ys)) // 1 4

	// append devolve um slice — que pode ser outro, se precisou crescer. Por
	// isso o resultado é SEMPRE reatribuído.
	ys = append(ys, 2, 3)
	fmt.Println(ys) // [1 2 3]

	// Fatiar: [início:fim), com o fim exclusivo. Omitir um lado significa
	// "do começo" ou "até o fim".
	fmt.Println(xs[1:], xs[:2], xs[1:2]) // [2 3] [1 2] [2]

	// O zero value de um slice é nil — e nil já funciona com len, range e
	// append. Não é preciso inicializar antes de usar.
	var zs []int
	fmt.Println(zs == nil, len(zs)) // true 0
	zs = append(zs, 9)
	fmt.Println(zs) // [9]

	// ... espalha um slice dentro de um append.
	fmt.Println(append([]int{0}, xs...)) // [0 1 2 3]
}

// SUA VEZ
//
// Devolva um slice novo com os elementos pares de xs, na mesma ordem.
func Pares(xs []int) []int {
	return xs // <- troque isto
}

func main() {
	exemplos()
}
