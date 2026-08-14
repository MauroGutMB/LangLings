// tamanho pede o texto emprestado — ela só precisa ler, nunca precisou ser
// dona dele. Isso é o que deixa `resumo` livre para usar `titulo` de novo
// depois de chamar esta função.
pub fn tamanho(s: &str) -> usize {
    s.len()
}

pub fn resumo(titulo: String, corpo: String) -> String {
    let n = tamanho(&titulo);
    format!("{titulo}: {n} caracteres. {corpo}")
}
