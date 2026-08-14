use ll_basico_variaveis_e_mutabilidade::media;

#[test]
fn media_de_inteiros_pares() {
    assert_eq!(media(2, 4), 3.0);
}

// A fração é o ponto do exercício: uma média que devolve 1.0 em vez de 1.5
// dividiu antes de converter.
#[test]
fn media_mantem_a_fracao() {
    assert_eq!(media(1, 2), 1.5, "a parte fracionária foi perdida?");
}

#[test]
fn media_com_negativos() {
    assert_eq!(media(-3, -4), -3.5);
}

#[test]
fn media_de_zeros() {
    assert_eq!(media(0, 0), 0.0);
}
