package main

import "testing"

func drenar(ch <-chan int) []int {
	out := []int{}
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func TestEnfileirarEntregaTudoEmOrdem(t *testing.T) {
	got := drenar(Enfileirar([]int{1, 2, 3}))

	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("recebi %v, quero %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("recebi %v, quero %v", got, want)
		}
	}
}

func TestEnfileirarFechaOCanal(t *testing.T) {
	ch := Enfileirar([]int{7})

	if v := <-ch; v != 7 {
		t.Fatalf("primeiro valor = %d, quero 7", v)
	}

	// A segunda recepção só volta se o canal estiver fechado; num canal aberto
	// e vazio ela bloquearia para sempre.
	v, aberto := <-ch
	if aberto {
		t.Fatalf("o canal ainda está aberto e entregou %d", v)
	}
}

func TestEnfileirarComMuitosValores(t *testing.T) {
	xs := make([]int, 1000)
	for i := range xs {
		xs[i] = i
	}

	got := drenar(Enfileirar(xs))

	if len(got) != len(xs) {
		t.Fatalf("recebi %d valores, quero %d", len(got), len(xs))
	}
	for i := range xs {
		if got[i] != xs[i] {
			t.Fatalf("posição %d = %d, quero %d", i, got[i], xs[i])
		}
	}
}

func TestEnfileirarComSliceVazio(t *testing.T) {
	if got := drenar(Enfileirar(nil)); len(got) != 0 {
		t.Fatalf("recebi %v, quero nada", got)
	}
	if got := drenar(Enfileirar([]int{})); len(got) != 0 {
		t.Fatalf("recebi %v, quero nada", got)
	}
}

// Enfileirar tem que voltar antes de qualquer leitura: quem chama recebe o
// canal e só então decide quando drenar.
func TestEnfileirarRetornaAntesDeAlguemLer(t *testing.T) {
	ch := Enfileirar([]int{1, 2, 3})

	// Chegar até aqui já prova que Enfileirar não ficou presa num envio.
	if got := drenar(ch); len(got) != 3 {
		t.Fatalf("recebi %v, quero 3 valores", got)
	}
}
