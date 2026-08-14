// TODO: isto não compila. O `for x in v` consome `v` (chama `into_iter`
// implicitamente), então `v` não existe mais para ser devolvido no final.
pub fn soma_e_lista(v: Vec<i32>) -> (i32, Vec<i32>) {
    let mut soma = 0;
    for x in v {
        soma += x;
    }
    (soma, v)
}
