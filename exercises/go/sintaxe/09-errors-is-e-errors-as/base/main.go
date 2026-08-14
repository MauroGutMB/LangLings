package main

import (
	"errors"
	"fmt"
)

// ErrNaoEncontrado é uma sentinela: um erro comparado por identidade.
var ErrNaoEncontrado = errors.New("não encontrado")

// ErroHTTP é um erro com dados dentro — o status só é acessível se você
// conseguir chegar ao valor concreto.
type ErroHTTP struct {
	Status int
}

func (e *ErroHTTP) Error() string {
	return fmt.Sprintf("resposta http %d", e.Status)
}

// Classificar devolve "ok", "ausente", "http:<status>" ou "outro".
//
// TODO: implemente. O esqueleto abaixo chuta sempre a mesma coisa.
func Classificar(err error) string {
	return "outro"
}

func main() {
	fmt.Println(Classificar(nil))
	fmt.Println(Classificar(fmt.Errorf("buscar usuário: %w", ErrNaoEncontrado)))
	fmt.Println(Classificar(fmt.Errorf("chamar api: %w", &ErroHTTP{Status: 503})))
}
