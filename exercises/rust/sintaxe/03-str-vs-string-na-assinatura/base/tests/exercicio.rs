use ll_sintaxe_str_vs_string_na_assinatura::{contem_arroba, validar_emails};

#[test]
fn contem_arroba_com_arroba() {
    assert!(contem_arroba("a@b.com"));
}

#[test]
fn contem_arroba_sem_arroba() {
    assert!(!contem_arroba("nao-e-email"));
}

// Esta é a asserção que separa a versão que compila da que não compila:
// `validar_emails` só existe para chamar `contem_arroba` com literais de
// dentro de um `&[&str]`, e é exatamente essa chamada que quebra com
// `&String` na assinatura.
#[test]
fn valida_uma_lista_de_literais() {
    let emails = ["a@b.com", "sem-arroba", "c@d.com"];
    assert_eq!(validar_emails(&emails), vec![true, false, true]);
}

#[test]
fn tambem_aceita_uma_string_emprestada() {
    let email = String::from("x@y.com");
    assert!(contem_arroba(&email));
}

#[test]
fn lista_vazia() {
    let emails: [&str; 0] = [];
    assert_eq!(validar_emails(&emails), Vec::<bool>::new());
}
