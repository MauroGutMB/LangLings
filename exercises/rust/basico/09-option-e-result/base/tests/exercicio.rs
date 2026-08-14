use ll_basico_option_e_result::somar_textos;

#[test]
fn soma_dois_numeros_validos() {
    assert_eq!(somar_textos("2", "3"), Ok(5));
}

#[test]
fn soma_com_negativos() {
    assert_eq!(somar_textos("-2", "5"), Ok(3));
}

// O ponto do exercício: um texto que não é número precisa virar Err, não um
// pânico nem um 0 silencioso.
#[test]
fn primeiro_invalido_e_err() {
    assert!(somar_textos("abc", "3").is_err());
}

#[test]
fn segundo_invalido_e_err() {
    assert!(somar_textos("2", "xyz").is_err());
}

#[test]
fn os_dois_invalidos_e_err() {
    assert!(somar_textos("abc", "xyz").is_err());
}
