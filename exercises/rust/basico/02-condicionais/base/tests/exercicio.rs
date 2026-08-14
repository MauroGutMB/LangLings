use ll_basico_condicionais::classificar;

#[test]
fn nota_baixa_reprova() {
    assert_eq!(classificar(0), "reprovado");
    assert_eq!(classificar(42), "reprovado");
}

#[test]
fn nota_media_aprova() {
    assert_eq!(classificar(60), "aprovado");
    assert_eq!(classificar(79), "aprovado");
}

#[test]
fn nota_alta_e_destaque() {
    assert_eq!(classificar(80), "destaque");
    assert_eq!(classificar(100), "destaque");
}

// As fronteiras são o que separa uma cadeia de condições certa de uma que
// erra por um: 59 ainda reprova, 60 já aprova.
#[test]
fn as_fronteiras_ficam_do_lado_certo() {
    assert_eq!(classificar(59), "reprovado");
    assert_eq!(classificar(60), "aprovado");
    assert_eq!(classificar(79), "aprovado");
    assert_eq!(classificar(80), "destaque");
}
