package main

import "fmt"

// Sem valor inicial de propósito: uma variável string de pacote nesse estado é
// exatamente o que a flag -X do linker consegue preencher.
var versao string

func main() {
	fmt.Println("olá do LangLings")
	fmt.Println("versão:", versao)
}
