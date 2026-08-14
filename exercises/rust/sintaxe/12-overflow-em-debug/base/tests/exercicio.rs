use ll_sintaxe_overflow_em_debug::fatorial;

#[test]
fn fatorial_pequeno() {
    assert_eq!(fatorial(5), 120);
}

#[test]
fn fatorial_de_zero() {
    assert_eq!(fatorial(0), 1);
}

#[test]
fn fatorial_de_um() {
    assert_eq!(fatorial(1), 1);
}

// A entrada que estoura um acumulador u32 (13! = 6227020800) mas cabe
// tranquilamente num u64 — esta é a asserção que separa as duas versões.
#[test]
fn fatorial_que_estoura_u32_mas_nao_u64() {
    assert_eq!(fatorial(13), 6_227_020_800);
}

#[test]
fn fatorial_ainda_maior() {
    assert_eq!(fatorial(15), 1_307_674_368_000);
}
