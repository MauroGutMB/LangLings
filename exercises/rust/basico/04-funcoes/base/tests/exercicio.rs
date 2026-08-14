use ll_basico_funcoes::maior_de_tres;

#[test]
fn maior_no_meio() {
    assert_eq!(maior_de_tres(3, 9, 5), 9);
}

#[test]
fn maior_no_comeco() {
    assert_eq!(maior_de_tres(9, 3, 5), 9);
}

#[test]
fn maior_no_fim() {
    assert_eq!(maior_de_tres(3, 5, 9), 9);
}

#[test]
fn todos_iguais() {
    assert_eq!(maior_de_tres(4, 4, 4), 4);
}

#[test]
fn com_negativos() {
    assert_eq!(maior_de_tres(-1, -9, -5), -1);
}
