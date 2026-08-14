package main

import "fmt"

// Placar é a pontuação de um time.
type Placar struct {
	Pontos int
}

// Marcar deve somar p aos pontos de time em m. Se o time não estiver no map,
// ele entra com p pontos. Com m nil, não faz nada.
//
// TODO: implemente. Comece pela forma mais direta e leia o erro do compilador.
func Marcar(m map[string]Placar, time string, p int) {
}

func main() {
	m := map[string]Placar{"go": {Pontos: 2}}

	Marcar(m, "go", 3)
	Marcar(m, "rust", 1)

	fmt.Println(m)
}
