use ll_sintaxe_lifetime_de_retorno::maior_str;

#[test]
fn primeiro_e_maior() {
    assert_eq!(maior_str("banana", "uva"), "banana");
}

// Esta é a asserção que exige a função devolver `b`: se o lifetime dela não
// estivesse ligado ao retorno, isto não teria como compilar.
#[test]
fn segundo_e_maior() {
    assert_eq!(maior_str("uva", "abacaxi"), "abacaxi");
}

#[test]
fn tamanhos_iguais_devolve_o_primeiro() {
    assert_eq!(maior_str("abc", "xyz"), "abc");
}

#[test]
fn com_textos_vazios() {
    assert_eq!(maior_str("", "a"), "a");
    assert_eq!(maior_str("a", ""), "a");
}
