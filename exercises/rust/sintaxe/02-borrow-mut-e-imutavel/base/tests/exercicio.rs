use ll_sintaxe_borrow_mut_e_imutavel::maior_e_dobro;

#[test]
fn devolve_o_maior_original() {
    let mut v = vec![1, 5, 3];
    assert_eq!(maior_e_dobro(&mut v), 5);
}

// Esta é a asserção que separa a versão que compila da que não compila: ela
// depende de `v` ter sido alterado DEPOIS de `maior` ter sido calculado a
// partir dele — exatamente o push que dispara o conflito de empréstimo.
#[test]
fn empurra_o_dobro_para_o_fim() {
    let mut v = vec![1, 5, 3];
    maior_e_dobro(&mut v);
    assert_eq!(v, vec![1, 5, 3, 10]);
}

#[test]
fn vetor_de_um_elemento() {
    let mut v = vec![7];
    assert_eq!(maior_e_dobro(&mut v), 7);
    assert_eq!(v, vec![7, 14]);
}

#[test]
fn maior_ja_no_fim() {
    let mut v = vec![2, 9];
    assert_eq!(maior_e_dobro(&mut v), 9);
    assert_eq!(v, vec![2, 9, 18]);
}
