package main

import "fmt"

// Um struct agrupa campos. Campos com inicial maiúscula são exportados; os
// com minúscula só são visíveis dentro do pacote.
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Método com receiver de VALOR: recebe uma cópia do struct. Serve para quem só
// lê — alterar a cópia não teria efeito nenhum lá fora.
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Método com receiver de PONTEIRO: recebe o endereço, então consegue alterar o
// struct original. Esta é a única forma de um método mutar quem o recebeu.
func (r *Retangulo) Escalar(f float64) {
	r.Largura *= f
	r.Altura *= f
}

// Um struct pode embutir outro: os campos e métodos do embutido sobem para o
// tipo de fora, como se tivessem sido declarados nele.
type Caixa struct {
	Retangulo
	Profundidade float64
}

func exemplos() {
	// Nomear os campos na criação é o padrão: resiste a reordenação e a campos
	// novos, ao contrário de Retangulo{3, 2}.
	r := Retangulo{Largura: 3, Altura: 2}
	fmt.Println(r.Area()) // 6

	// r é uma variável endereçável, então Go traduz r.Escalar(2) para
	// (&r).Escalar(2) sozinho — não é preciso escrever o & na chamada.
	r.Escalar(2)
	fmt.Println(r) // {6 4}

	// O zero value de um struct é ele com todos os campos zerados, e já é
	// utilizável: não existe "construtor obrigatório".
	var vazio Retangulo
	fmt.Println(vazio, vazio.Area()) // {0 0} 0

	// Structs são valores: atribuir COPIA todos os campos.
	copia := r
	copia.Largura = 100
	fmt.Println(r.Largura, copia.Largura) // 6 100

	c := Caixa{Retangulo: r, Profundidade: 5}
	fmt.Println(c.Largura, c.Area()) // 6 24  — campo e método vieram do embutido
}

// Perimetro devolve o perímetro do retângulo.
//
// O receiver é de valor porque o método só lê: usar ponteiro aqui daria a quem
// chama a impressão errada de que o retângulo pode sair alterado.
func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

func main() {
	exemplos()
}
