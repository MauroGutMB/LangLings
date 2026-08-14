use ll_basico_composicao::{mais_caro, Produto};

#[test]
fn encontra_o_mais_caro_no_meio() {
    let estoque = vec![
        Produto::novo("caneta", 2.5, 100),
        Produto::novo("mochila", 90.0, 5),
        Produto::novo("caderno", 8.0, 30),
    ];
    assert_eq!(mais_caro(&estoque), Some(&estoque[1]));
}

#[test]
fn encontra_o_mais_caro_no_fim() {
    let estoque = vec![
        Produto::novo("caneta", 2.5, 100),
        Produto::novo("caderno", 8.0, 30),
        Produto::novo("mochila", 90.0, 5),
    ];
    assert_eq!(mais_caro(&estoque), Some(&estoque[2]));
}

// Lista vazia é o caso que separa uma implementação ingênua (que panica ou
// devolve um valor inventado) da correta.
#[test]
fn lista_vazia_e_none() {
    let estoque: Vec<Produto> = vec![];
    assert_eq!(mais_caro(&estoque), None);
}

#[test]
fn um_unico_produto() {
    let estoque = vec![Produto::novo("caneta", 2.5, 100)];
    assert_eq!(mais_caro(&estoque), Some(&estoque[0]));
}
