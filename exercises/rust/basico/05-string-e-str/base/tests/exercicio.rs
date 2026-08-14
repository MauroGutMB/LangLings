use ll_basico_string_e_str::gritar;

#[test]
fn palavra_simples() {
    assert_eq!(gritar("oi"), "OI!");
}

#[test]
fn ja_maiuscula() {
    assert_eq!(gritar("OI"), "OI!");
}

#[test]
fn frase_com_espacos() {
    assert_eq!(gritar("bom dia"), "BOM DIA!");
}

#[test]
fn string_vazia() {
    assert_eq!(gritar(""), "!");
}

#[test]
fn aceita_literal_e_string_emprestada() {
    let s = String::from("ok");
    assert_eq!(gritar(&s), "OK!");
    assert_eq!(gritar("ok"), "OK!");
}
