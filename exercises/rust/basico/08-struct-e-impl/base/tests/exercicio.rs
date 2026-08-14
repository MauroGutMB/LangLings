use ll_basico_struct_e_impl::Retangulo;

#[test]
fn quadrado_e_verdadeiro() {
    let r = Retangulo::novo(5, 5);
    assert!(r.eh_quadrado());
}

#[test]
fn nao_quadrado_e_falso() {
    let r = Retangulo::novo(3, 4);
    assert!(!r.eh_quadrado());
}

#[test]
fn area_continua_correta() {
    let r = Retangulo::novo(3, 4);
    assert_eq!(r.area(), 12);
}

#[test]
fn escalar_continua_correto() {
    let mut r = Retangulo::novo(2, 5);
    r.escalar(3);
    assert_eq!(r.largura, 6);
    assert_eq!(r.altura, 15);
}

#[test]
fn quadrado_apos_escalar_uniforme() {
    let mut r = Retangulo::novo(2, 2);
    r.escalar(4);
    assert!(r.eh_quadrado());
}
