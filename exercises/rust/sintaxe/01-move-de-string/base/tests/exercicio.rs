use ll_sintaxe_move_de_string::resumo;

#[test]
fn junta_titulo_tamanho_e_corpo() {
    // Esta é a asserção que separa a versão que compila da que não compila:
    // ela usa `titulo` depois da chamada a `tamanho`, exatamente como o
    // corpo de `resumo` precisa fazer.
    assert_eq!(
        resumo("Rust".to_string(), "linguagem de sistemas".to_string()),
        "Rust: 4 caracteres. linguagem de sistemas"
    );
}

#[test]
fn titulo_vazio() {
    assert_eq!(resumo("".to_string(), "x".to_string()), ": 0 caracteres. x");
}

#[test]
fn titulo_com_acentos_conta_bytes() {
    // 'é' ocupa 2 bytes em UTF-8, então "café" tem 5, não 4.
    assert_eq!(resumo("café".to_string(), "".to_string()), "café: 5 caracteres. ");
}
