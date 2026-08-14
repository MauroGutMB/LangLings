use ll_sintaxe_impl_trait_vs_generics::escolher;

#[test]
fn intervalo() {
    let v: Vec<i32> = escolher(true).collect();
    assert_eq!(v, vec![1, 2, 3, 4]);
}

// Esta é a asserção que exige o OUTRO tipo concreto: se a função só pudesse
// devolver um tipo fixo escolhido em tempo de compilação, este ramo do `if`
// não teria como coexistir com o de cima.
#[test]
fn vetor() {
    let v: Vec<i32> = escolher(false).collect();
    assert_eq!(v, vec![10, 20]);
}

#[test]
fn intervalo_soma() {
    let soma: i32 = escolher(true).sum();
    assert_eq!(soma, 10);
}
