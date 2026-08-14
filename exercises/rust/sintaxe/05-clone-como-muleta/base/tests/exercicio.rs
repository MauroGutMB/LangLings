use ll_sintaxe_clone_como_muleta::dobrar_todos;

// Esta é a asserção que separa a versão com clone da correta: dobrar uma
// cópia deixa `v` sem alteração nenhuma, então este teste falha na versão
// ingênua mesmo ela compilando sem erro.
#[test]
fn dobra_in_place() {
    let mut v = vec![1, 2, 3];
    dobrar_todos(&mut v);
    assert_eq!(v, vec![2, 4, 6]);
}

#[test]
fn com_zero() {
    let mut v = vec![0, 5];
    dobrar_todos(&mut v);
    assert_eq!(v, vec![0, 10]);
}

#[test]
fn vetor_vazio() {
    let mut v: Vec<i32> = vec![];
    dobrar_todos(&mut v);
    assert_eq!(v, Vec::<i32>::new());
}

#[test]
fn com_negativos() {
    let mut v = vec![-3, 4];
    dobrar_todos(&mut v);
    assert_eq!(v, vec![-6, 8]);
}
