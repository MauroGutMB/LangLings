// `for x in &v` empresta cada elemento em vez de consumir o vetor — `v`
// continua existindo depois do laço para ser devolvido.
pub fn soma_e_lista(v: Vec<i32>) -> (i32, Vec<i32>) {
    let mut soma = 0;
    for x in &v {
        soma += x;
    }
    (soma, v)
}
