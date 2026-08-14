use ll_basico_vetores::pares;

#[test]
fn mistura_de_pares_e_impares() {
    assert_eq!(pares(&[1, 2, 3, 4]), vec![2, 4]);
}

#[test]
fn mantem_a_ordem_original() {
    assert_eq!(pares(&[4, 3, 2, 1]), vec![4, 2]);
}

#[test]
fn nenhum_par() {
    assert_eq!(pares(&[1, 3, 5]), Vec::<i32>::new());
}

#[test]
fn todos_pares() {
    assert_eq!(pares(&[2, 4, 6]), vec![2, 4, 6]);
}

#[test]
fn slice_vazio() {
    assert_eq!(pares(&[]), Vec::<i32>::new());
}

#[test]
fn com_negativos() {
    assert_eq!(pares(&[-4, -3, -2, -1]), vec![-4, -2]);
}
