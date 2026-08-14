package main

import "fmt"

// Acumular soma n ao valor de k em m e devolve o map resultante.
//
// O map só é criado quando m é nil: alocar sempre e copiar seria mais caro e
// mudaria a semântica, já que quem chamou com um map de verdade espera ver a
// própria estrutura alterada.
func Acumular(m map[string]int, k string, n int) map[string]int {
	if m == nil {
		m = make(map[string]int)
	}
	m[k] += n
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
