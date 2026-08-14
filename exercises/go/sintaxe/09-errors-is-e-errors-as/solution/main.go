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
// Os dois testes de erro fazem trabalhos diferentes: errors.Is compara a cadeia
// com um erro que já existe (a sentinela), e errors.As procura na mesma cadeia
// o primeiro erro de um tipo e o coloca numa variável — o único jeito de chegar
// ao campo Status. Nenhum dos dois pode ser trocado por == ou por asserção de
// tipo sem quebrar assim que alguém embrulhar o erro com %w.
func Classificar(err error) string {
	if err == nil {
		return "ok"
	}

	if errors.Is(err, ErrNaoEncontrado) {
		return "ausente"
	}

	var httpErr *ErroHTTP
	if errors.As(err, &httpErr) {
		return fmt.Sprintf("http:%d", httpErr.Status)
	}

	return "outro"
}

func main() {
	fmt.Println(Classificar(nil))
	fmt.Println(Classificar(fmt.Errorf("buscar usuário: %w", ErrNaoEncontrado)))
	fmt.Println(Classificar(fmt.Errorf("chamar api: %w", &ErroHTTP{Status: 503})))
}
