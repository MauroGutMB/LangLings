package main

import (
	"fmt"
	"strings"
)

// Uma string em Go é uma sequência imutável de BYTES. Ela costuma conter texto
// UTF-8, mas o tipo não sabe disso: quem interpreta os bytes como caracteres é
// quem lê.
func exemplos() {
	s := "olá"

	fmt.Println(len(s))      // 4  — o "á" ocupa 2 bytes em UTF-8
	fmt.Println(s[0])        // 111 — indexar devolve um byte (um uint8)
	fmt.Printf("%c\n", s[0]) // o   — %c interpreta esse número como caractere

	// range sobre string entrega runas (code points) e o índice EM BYTES.
	// Repare que o índice pula de 1 para 2 e o próximo seria 4.
	for i, r := range s {
		fmt.Printf("%d:%c ", i, r) // 0:o 1:l 2:á
	}
	fmt.Println()

	// Converter para []rune dá acesso posicional por caractere, ao custo de
	// uma cópia — por isso não é o padrão.
	rs := []rune(s)
	fmt.Println(len(rs), string(rs[2])) // 3 á

	// Strings são imutáveis: cada concatenação cria uma string nova. Num laço
	// isso vira lixo proporcional ao quadrado do tamanho. strings.Builder
	// escreve num buffer único e só materializa a string no fim.
	var b strings.Builder
	for _, r := range rs {
		b.WriteRune(r)
		b.WriteByte('-')
	}
	fmt.Println(b.String()) // o-l-á-

	fmt.Println(strings.ToUpper(s))       // OLÁ
	fmt.Println(strings.Contains(s, "l")) // true
}

// SUA VEZ
//
// Devolva quantos caracteres existem em s: Contar("olá") é 3.
func Contar(s string) int {
	return len(s) // <- troque isto
}

func main() {
	exemplos()
}
