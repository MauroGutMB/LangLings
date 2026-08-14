package main

import "fmt"

// Contador guarda um número.
type Contador struct {
	N int
}

// Incrementar soma 1 ao N deste contador.
//
// O receiver é ponteiro porque o método muta: com receiver de valor, o c++
// aconteceria numa cópia descartada no fim da chamada — e o compilador não
// teria nada a reclamar.
func (c *Contador) Incrementar() {
	c.N++
}

// Dobro devolve um contador novo com o dobro de N.
//
// Receiver de valor: o método declara, na assinatura, que não altera a origem.
// Como c já é uma cópia, nem seria preciso construir outro struct — mas montar
// o retorno explicitamente deixa a intenção visível.
func (c Contador) Dobro() Contador {
	return Contador{N: c.N * 2}
}

func main() {
	c := Contador{N: 3}

	c.Incrementar()
	fmt.Println("depois de Incrementar:", c.N)

	d := c.Dobro()
	fmt.Println("dobro:", d.N, "original:", c.N)
}
