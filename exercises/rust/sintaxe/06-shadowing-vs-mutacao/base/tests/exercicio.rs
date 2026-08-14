use ll_sintaxe_shadowing_vs_mutacao::normalizar;

#[test]
fn remove_espacos_das_pontas() {
    assert_eq!(normalizar("  oi  "), "oi");
}

// Esta é a asserção que separa a versão que compila-mas-engana da correta:
// o trim sozinho já faz este teste passar, então é a MINÚSCULA que expõe o
// shadow perdido dentro do bloco.
#[test]
fn deixa_minusculo() {
    assert_eq!(normalizar("  OI  "), "oi");
}

#[test]
fn mistura_maiusculas_e_espacos() {
    assert_eq!(normalizar(" Bom DIA "), "bom dia");
}

#[test]
fn ja_normalizado() {
    assert_eq!(normalizar("ok"), "ok");
}
