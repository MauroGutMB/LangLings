package main

import "fmt"

// Em Go toda variável tem um tipo, e toda variável declarada sem valor recebe
// o "zero value" do tipo dela. Não existe variável não inicializada.
func exemplos() {
	var nome string = "Ana"  // forma completa: var, nome, tipo, valor
	idade := 30              // := infere o tipo a partir do valor (int)
	fmt.Println(nome, idade) // Ana 30

	// Sem valor inicial, cada tipo tem o seu zero.
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	fmt.Println(i, f, b)        // 0 0 false
	fmt.Printf("%q %v\n", s, p) // "" <nil>

	// Conversão entre tipos numéricos é sempre explícita: Go não promove int
	// para float64 sozinho. Repare que a conversão vem ANTES da divisão — se
	// viesse depois, a divisão inteira já teria descartado a fração.
	metade := float64(idade) / 2
	fmt.Println(metade)             // 15
	fmt.Println(float64(idade / 4)) // 7  <- fração perdida na divisão inteira
	fmt.Println(float64(idade) / 4) // 7.5

	// const é resolvida em tempo de compilação.
	const maiorIdade = 18
	fmt.Println(idade >= maiorIdade) // true
}

// SUA VEZ
//
// Devolva a média de a e b como float64: Media(1, 2) é 1.5.
func Media(a, b int) float64 {
	return 0 // <- troque isto
}

func main() {
	exemplos()
}
