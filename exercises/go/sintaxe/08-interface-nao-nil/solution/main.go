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
// O ponteiro nunca aparece no return do caminho feliz. Um error é uma
// interface, ou seja, um par (tipo, valor): devolver uma variável *ErroValidacao
// valendo nil preenche o lado do tipo, e a interface resultante é diferente de
// nil mesmo sem erro nenhum dentro dela.
func Validar(campo, valor string) error {
	if valor == "" {
		return &ErroValidacao{Campo: campo}
	}
	return nil
}

func main() {
	fmt.Println(Validar("nome", "Ana") == nil)
	fmt.Println(Validar("nome", ""))
}
