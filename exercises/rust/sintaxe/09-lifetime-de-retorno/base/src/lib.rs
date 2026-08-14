// TODO: isto não compila. `'a` só amarra o retorno a `a`, e a função pode
// devolver `b` — cujo lifetime não tem nenhuma relação provada com `'a`.
pub fn maior_str<'a>(a: &'a str, b: &str) -> &'a str {
    if a.len() >= b.len() { a } else { b }
}
