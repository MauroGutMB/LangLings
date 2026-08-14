// TODO: isto não compila. O compilador vai reclamar de `titulo` usado depois
// de movido — mas nem `resumo` nem o que ela devolve estão errados.
pub fn tamanho(s: String) -> usize {
    s.len()
}

pub fn resumo(titulo: String, corpo: String) -> String {
    let n = tamanho(titulo);
    format!("{titulo}: {n} caracteres. {corpo}")
}
