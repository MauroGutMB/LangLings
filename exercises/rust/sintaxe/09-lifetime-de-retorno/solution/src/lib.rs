// Os dois parâmetros compartilham o mesmo lifetime 'a: o retorno passa a
// viver tanto quanto o MENOR dos dois textos recebidos, o que é sempre
// verdade não importa qual dos dois seja devolvido.
pub fn maior_str<'a>(a: &'a str, b: &'a str) -> &'a str {
    if a.len() >= b.len() { a } else { b }
}
