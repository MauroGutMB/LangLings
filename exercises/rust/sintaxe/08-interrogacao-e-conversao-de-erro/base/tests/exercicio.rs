use ll_sintaxe_interrogacao_e_conversao_de_erro::dobrar_texto;

#[test]
fn dobra_um_numero_valido() {
    assert_eq!(dobrar_texto("21"), Ok(42));
}

// A asserção que só faz sentido depois do código compilar: um texto que não
// é número precisa virar Err, não interromper o build inteiro.
#[test]
fn texto_invalido_e_err() {
    assert!(dobrar_texto("abc").is_err());
}

#[test]
fn negativo_valido() {
    assert_eq!(dobrar_texto("-5"), Ok(-10));
}

#[test]
fn texto_vazio_e_err() {
    assert!(dobrar_texto("").is_err());
}
