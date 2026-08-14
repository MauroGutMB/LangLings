use ll_basico_hashmap::contar_palavras;

#[test]
fn palavra_repetida() {
    let resultado = contar_palavras("a b a");
    assert_eq!(resultado.get("a"), Some(&2));
    assert_eq!(resultado.get("b"), Some(&1));
    assert_eq!(resultado.len(), 2);
}

#[test]
fn todas_diferentes() {
    let resultado = contar_palavras("um dois tres");
    assert_eq!(resultado.get("um"), Some(&1));
    assert_eq!(resultado.get("dois"), Some(&1));
    assert_eq!(resultado.get("tres"), Some(&1));
}

#[test]
fn texto_vazio() {
    let resultado = contar_palavras("");
    assert_eq!(resultado.len(), 0);
}

#[test]
fn palavra_unica_repetida_varias_vezes() {
    let resultado = contar_palavras("oi oi oi oi");
    assert_eq!(resultado.get("oi"), Some(&4));
    assert_eq!(resultado.len(), 1);
}

#[test]
fn chave_ausente_e_none() {
    let resultado = contar_palavras("a b a");
    assert_eq!(resultado.get("z"), None);
}
