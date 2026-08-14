package main

import (
	"errors"
	"fmt"
)

// ErrVazio é devolvido (embrulhado) quando não há nada para somar.
var ErrVazio = errors.New("nenhum elemento")

// Processar soma xs e devolve o total. Com xs vazio, devolve 0 e um erro que
// embrulha ErrVazio. Nos dois casos ela precisa chamar registrar exatamente
// uma vez, com o total devolvido.
//
// TODO: implemente. O esqueleto abaixo nunca registra nada.
func Processar(xs []int, registrar func(int)) (total int, err error) {
	return 0, nil
}

func main() {
	log := func(n int) { fmt.Println("registrado:", n) }

	fmt.Println(Processar([]int{1, 2, 3}, log))
	fmt.Println(Processar(nil, log))
}
