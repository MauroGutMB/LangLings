package main

import (
	"errors"
	"testing"
)

func TestValidarComValorPreenchido(t *testing.T) {
	err := Validar("nome", "Ana")

	if err != nil {
		t.Fatalf("Validar devolveu %#v, quero nil", err)
	}
}

func TestValidarComValorVazio(t *testing.T) {
	err := Validar("nome", "")

	if err == nil {
		t.Fatal("Validar com valor vazio devolveu nil, quero um erro")
	}

	var alvo *ErroValidacao
	if !errors.As(err, &alvo) {
		t.Fatalf("Validar devolveu %v, quero um *ErroValidacao", err)
	}
	if alvo.Campo != "nome" {
		t.Fatalf("Campo = %q, quero %q", alvo.Campo, "nome")
	}
}

// Quem chama costuma testar o erro dentro de um if com statement. Se a
// interface devolvida carregar um ponteiro nil, este bloco entra.
func TestValidarNoIdiomaDeQuemChama(t *testing.T) {
	campos := map[string]string{"nome": "Ana", "email": "ana@exemplo.com"}

	for campo, valor := range campos {
		if err := Validar(campo, valor); err != nil {
			t.Fatalf("Validar(%q, %q) devolveu erro %v, não queria nenhum", campo, valor, err)
		}
	}
}

func TestValidarComEspacoNaoEVazio(t *testing.T) {
	if err := Validar("nome", " "); err != nil {
		t.Fatalf("Validar(%q, %q) devolveu %v, quero nil", "nome", " ", err)
	}
}

// O erro devolvido precisa ser utilizável: chamar Error() num ponteiro nil
// embrulhado numa interface entra em pânico.
func TestErroDevolvidoTemMensagem(t *testing.T) {
	err := Validar("email", "")
	if err == nil {
		t.Fatal("Validar com valor vazio devolveu nil")
	}

	if err.Error() == "" {
		t.Fatal("o erro devolvido não tem mensagem")
	}
}
