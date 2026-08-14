package main

import "fmt"

// ErroValidacao descreve um campo que não passou na validação.
type ErroValidacao struct {
	Campo string
}

func (e *ErroValidacao) Error() string {
	return "campo obrigatório não preenchido: " + e.Campo
}

// Validar devolve um *ErroValidacao quando valor está vazio e nil quando não
// está.
//
// TODO: o if está certo, o erro devolvido está certo, e mesmo assim quem chama
// vê um erro no caminho feliz. Corrija.
func Validar(campo, valor string) error {
	var err *ErroValidacao

	if valor == "" {
		err = &ErroValidacao{Campo: campo}
	}

	return err
}

func main() {
	fmt.Println(Validar("nome", "Ana") == nil)
	fmt.Println(Validar("nome", ""))
}
