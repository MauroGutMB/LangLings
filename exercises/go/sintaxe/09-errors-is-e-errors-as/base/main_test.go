package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestClassificarNil(t *testing.T) {
	if got := Classificar(nil); got != "ok" {
		t.Fatalf("Classificar(nil) = %q, quero %q", got, "ok")
	}
}

func TestClassificarSentinelaDireta(t *testing.T) {
	if got := Classificar(ErrNaoEncontrado); got != "ausente" {
		t.Fatalf("Classificar(ErrNaoEncontrado) = %q, quero %q", got, "ausente")
	}
}

// Comparar com == passa no teste acima e falha aqui.
func TestClassificarSentinelaEmbrulhada(t *testing.T) {
	err := fmt.Errorf("buscar usuário 7: %w", ErrNaoEncontrado)

	if got := Classificar(err); got != "ausente" {
		t.Fatalf("Classificar(%v) = %q, quero %q", err, got, "ausente")
	}
}

func TestClassificarSentinelaEmbrulhadaDuasVezes(t *testing.T) {
	err := fmt.Errorf("camada externa: %w", fmt.Errorf("camada interna: %w", ErrNaoEncontrado))

	if got := Classificar(err); got != "ausente" {
		t.Fatalf("Classificar(%v) = %q, quero %q", err, got, "ausente")
	}
}

func TestClassificarErroHTTPDireto(t *testing.T) {
	if got := Classificar(&ErroHTTP{Status: 404}); got != "http:404" {
		t.Fatalf("Classificar(&ErroHTTP{404}) = %q, quero %q", got, "http:404")
	}
}

// Uma asserção de tipo (err.(*ErroHTTP)) passa no teste acima e falha aqui:
// o erro embrulhado não tem o tipo do erro embrulhado.
func TestClassificarErroHTTPEmbrulhado(t *testing.T) {
	err := fmt.Errorf("chamar api: %w", &ErroHTTP{Status: 500})

	if got := Classificar(err); got != "http:500" {
		t.Fatalf("Classificar(%v) = %q, quero %q", err, got, "http:500")
	}
}

func TestClassificarErroQualquer(t *testing.T) {
	if got := Classificar(errors.New("cabo desconectado")); got != "outro" {
		t.Fatalf("Classificar(erro qualquer) = %q, quero %q", got, "outro")
	}
}

func TestClassificarErroQualquerEmbrulhado(t *testing.T) {
	err := fmt.Errorf("contexto: %w", errors.New("cabo desconectado"))

	if got := Classificar(err); got != "outro" {
		t.Fatalf("Classificar(%v) = %q, quero %q", err, got, "outro")
	}
}
