use ll_sintaxe_match_nao_exaustivo::{pode_logar, Estado};

#[test]
fn ativo_pode_logar() {
    assert!(pode_logar(&Estado::Ativo));
}

#[test]
fn inativo_nao_pode_logar() {
    assert!(!pode_logar(&Estado::Inativo));
}

// A variante que falta no match: sem cobri-la, o código nem compila, então
// esta asserção só roda depois do braço estar escrito.
#[test]
fn suspenso_nao_pode_logar() {
    assert!(!pode_logar(&Estado::Suspenso));
}
