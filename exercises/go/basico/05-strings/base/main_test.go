package main

import "testing"

func TestContarAsciiPuro(t *testing.T) {
	if got := Contar("abc"); got != 3 {
		t.Fatalf("Contar(%q) = %d, quero 3", "abc", got)
	}
}

// Aqui bytes e caracteres deixam de coincidir: "olá" tem 4 bytes e 3 runas.
func TestContarComAcento(t *testing.T) {
	if got := Contar("olá"); got != 3 {
		t.Fatalf("Contar(%q) = %d, quero 3 (contou bytes em vez de caracteres?)", "olá", got)
	}
}

func TestContarComEmoji(t *testing.T) {
	if got := Contar("oi🐹"); got != 3 {
		t.Fatalf("Contar(%q) = %d, quero 3", "oi🐹", got)
	}
}

func TestContarStringVazia(t *testing.T) {
	if got := Contar(""); got != 0 {
		t.Fatalf("Contar(%q) = %d, quero 0", "", got)
	}
}
