use ll_basico_lacos::soma_pares;

#[test]
fn soma_ate_dez() {
    assert_eq!(soma_pares(10), 30);
}

// O limite entra na conta quando ele próprio é par: por isso 10 vale 30 e não
// 20. Um intervalo exclusivo (`0..ate`) erraria exatamente aqui.
#[test]
fn o_limite_e_inclusivo() {
    assert_eq!(soma_pares(2), 2);
    assert_eq!(soma_pares(4), 6);
}

#[test]
fn limite_impar_nao_muda_nada() {
    assert_eq!(soma_pares(9), 20);
    assert_eq!(soma_pares(1), 0);
}

#[test]
fn soma_de_zero() {
    assert_eq!(soma_pares(0), 0);
}
