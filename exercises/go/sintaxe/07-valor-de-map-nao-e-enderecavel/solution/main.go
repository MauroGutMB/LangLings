package main

import "fmt"

// Placar é a pontuação de um time.
type Placar struct {
	Pontos int
}

// Marcar soma p aos pontos de time em m.
//
// m[time].Pontos += p não compila: um map pode remanejar os próprios buckets
// ao crescer, então o valor guardado numa chave não tem endereço estável e o
// compilador se recusa a deixar você escrever "dentro" dele. O caminho é ler,
// alterar a cópia e gravá-la de volta — e a leitura de chave ausente já
// devolve o zero value, o que faz o time novo entrar com p sem caso especial.
func Marcar(m map[string]Placar, time string, p int) {
	if m == nil {
		return
	}

	atual := m[time]
	atual.Pontos += p
	m[time] = atual
}

func main() {
	m := map[string]Placar{"go": {Pontos: 2}}

	Marcar(m, "go", 3)
	Marcar(m, "rust", 1)

	fmt.Println(m)
}
