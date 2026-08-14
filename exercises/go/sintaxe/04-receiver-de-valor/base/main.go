package main

import "fmt"

// Contador guarda um número.
type Contador struct {
	N int
}

// Incrementar deve somar 1 ao N deste contador.
//
// TODO: o corpo está certo e mesmo assim nada acontece lá fora. Por quê?
func (c Contador) Incrementar() {
	c.N++
}

// Dobro deve devolver um contador NOVO com o dobro de N, deixando este intacto.
//
// TODO: o valor devolvido está certo e o contador de origem sai alterado.
func (c *Contador) Dobro() Contador {
	c.N *= 2
	return *c
}

func main() {
	c := Contador{N: 3}

	c.Incrementar()
	fmt.Println("depois de Incrementar:", c.N)

	d := c.Dobro()
	fmt.Println("dobro:", d.N, "original:", c.N)
}
