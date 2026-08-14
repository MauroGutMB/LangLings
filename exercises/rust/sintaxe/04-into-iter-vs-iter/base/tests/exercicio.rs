use ll_sintaxe_into_iter_vs_iter::soma_e_lista;

#[test]
fn soma_correta() {
    let (soma, _) = soma_e_lista(vec![1, 2, 3]);
    assert_eq!(soma, 6);
}

// Esta é a asserção que separa a versão que compila da que não compila: ela
// depende do vetor original ainda existir DEPOIS do laço que soma seus
// elementos — exatamente o que `for x in v` (consumindo) impede.
#[test]
fn vetor_original_devolvido_intacto() {
    let (_, v) = soma_e_lista(vec![1, 2, 3]);
    assert_eq!(v, vec![1, 2, 3]);
}

#[test]
fn vetor_vazio() {
    assert_eq!(soma_e_lista(vec![]), (0, vec![]));
}

#[test]
fn com_negativos() {
    assert_eq!(soma_e_lista(vec![-1, 5, -2]), (2, vec![-1, 5, -2]));
}
