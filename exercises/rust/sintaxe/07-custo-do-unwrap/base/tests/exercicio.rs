use ll_sintaxe_custo_do_unwrap::primeiro_positivo;

#[test]
fn acha_o_primeiro_positivo() {
    assert_eq!(primeiro_positivo(&[-3, -1, 4, 7]), Some(4));
}

// Esta é a asserção que separa a versão com unwrap da correta: sem nenhum
// positivo, a versão ingênua panica em vez de devolver None.
#[test]
fn sem_nenhum_positivo_e_none() {
    assert_eq!(primeiro_positivo(&[-3, -1, 0, -7]), None);
}

#[test]
fn vetor_vazio_e_none() {
    assert_eq!(primeiro_positivo(&[]), None);
}

#[test]
fn primeiro_elemento_ja_positivo() {
    assert_eq!(primeiro_positivo(&[9, -1, 4]), Some(9));
}
