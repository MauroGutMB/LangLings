package main

import "fmt"

// O if de Go não usa parênteses na condição, mas as chaves são obrigatórias.
func exemplos() {
	n := 7

	if n%2 == 0 {
		fmt.Println("par")
	} else {
		fmt.Println("ímpar") // ímpar
	}

	// O if aceita um statement antes da condição, separado por ponto e vírgula.
	// A variável declarada ali vive só dentro do if/else — é assim que Go evita
	// deixar variáveis temporárias vazando para o escopo de fora.
	if dobro := n * 2; dobro > 10 {
		fmt.Println(dobro, "passou de 10") // 14 passou de 10
	}

	// switch sem valor à frente: cada case é uma condição booleana. É a escada
	// de if/else escrita de forma plana.
	switch {
	case n < 0:
		fmt.Println("negativo")
	case n < 10:
		fmt.Println("um dígito") // um dígito
	default:
		fmt.Println("dois ou mais dígitos")
	}

	// switch com valor: os cases comparam contra ele, e um case aceita vários
	// rótulos. Não existe fall-through implícito — cada case para sozinho.
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("ímpar de um dígito") // ímpar de um dígito
	case 0, 2, 4, 6, 8:
		fmt.Println("par de um dígito")
	}
}

// SUA VEZ
//
// Devolva "baixo" para n < 10, "medio" para n de 10 a 99 e "alto" para
// n >= 100.
func Faixa(n int) string {
	return "" // <- troque isto
}

func main() {
	exemplos()
}
