package main

import (
	"errors"
	"fmt"
)

// Uma "sentinela" é um erro declarado uma vez e comparado por identidade. Ela
// é parte da API do pacote: quem chama pode reagir a ESTE erro em específico.
var ErrNaoEncontrado = errors.New("não encontrado")

func buscar(m map[string]int, k string) (int, error) {
	v, ok := m[k]
	if !ok {
		// %w embrulha o erro: a mensagem ganha contexto e a causa original
		// continua acessível a quem quiser interrogá-la.
		return 0, fmt.Errorf("buscar %q: %w", k, ErrNaoEncontrado)
	}
	return v, nil
}

func exemplos() {
	m := map[string]int{"a": 1}

	v, err := buscar(m, "a")
	fmt.Println(v, err) // 1 <nil>

	_, err = buscar(m, "z")
	fmt.Println(err)                              // buscar "z": não encontrado
	fmt.Println(err == ErrNaoEncontrado)          // false — o erro devolvido é o embrulho
	fmt.Println(errors.Is(err, ErrNaoEncontrado)) // true  — mas a causa está lá dentro

	// Com %v em vez de %w a mensagem sai idêntica e a causa se perde. Este é o
	// motivo de o verbo importar: o texto não é a informação.
	solto := fmt.Errorf("buscar %q: %v", "z", ErrNaoEncontrado)
	fmt.Println(solto)                              // buscar "z": não encontrado
	fmt.Println(errors.Is(solto, ErrNaoEncontrado)) // false

	// O idioma é checar o erro na hora e sair cedo.
	if _, err := buscar(m, "z"); err != nil {
		fmt.Println("falhou:", err) // falhou: buscar "z": não encontrado
	}
}

// ErrDivisaoPorZero é a sentinela que Dividir precisa embrulhar.
var ErrDivisaoPorZero = errors.New("divisão por zero")

// SUA VEZ
//
// Devolva a/b e nil. Com b igual a zero, devolva 0 e um erro que embrulhe
// ErrDivisaoPorZero e acrescente contexto à mensagem.
func Dividir(a, b int) (int, error) {
	return 0, nil // <- troque isto
}

func main() {
	exemplos()
}
