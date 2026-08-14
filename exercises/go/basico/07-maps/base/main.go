package main

import "fmt"

// Um map é uma coleção de pares chave -> valor. A chave pode ser qualquer tipo
// comparável (string, int, struct sem slices dentro); o valor, qualquer tipo.
func exemplos() {
	idades := map[string]int{"ana": 30}
	idades["bruno"] = 25       // insere ou sobrescreve
	fmt.Println(idades["ana"]) // 30

	// Ler uma chave ausente NÃO é erro: devolve o zero value do tipo do valor.
	// É o que faz o idioma de contagem (m[k]++) funcionar sem inicialização.
	fmt.Println(idades["zoe"]) // 0

	// A forma de duas variáveis é a única maneira de distinguir "não existe"
	// de "existe valendo zero".
	v, ok := idades["zoe"]
	fmt.Println(v, ok) // 0 false

	delete(idades, "ana")
	fmt.Println(len(idades)) // 1

	// make cria um map vazio pronto para escrita. O zero value de um map é
	// nil: dá para ler dele, mas escrever nele entra em pânico.
	contagem := make(map[string]int)
	for _, p := range []string{"a", "b", "a"} {
		contagem[p]++
	}
	fmt.Println(contagem) // map[a:2 b:1]

	// A ordem do range sobre um map é deliberadamente aleatória: se você
	// precisa de ordem, ordene as chaves você mesmo.
	total := 0
	for _, n := range contagem {
		total += n
	}
	fmt.Println(total) // 3
}

// SUA VEZ
//
// Devolva um map com quantas vezes cada palavra aparece em palavras.
func Contagem(palavras []string) map[string]int {
	return nil // <- troque isto
}

func main() {
	exemplos()
}
