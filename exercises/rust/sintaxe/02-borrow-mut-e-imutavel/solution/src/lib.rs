// `maior` é um i32 independente, não uma referência para dentro de `v` — o
// `*` desreferencia (e copia, porque i32 é Copy) no momento da leitura, então
// o empréstimo imutável de `v.iter()` já terminou quando o `push` acontece.
pub fn maior_e_dobro(v: &mut Vec<i32>) -> i32 {
    let maior = *v.iter().max().unwrap();
    v.push(maior * 2);
    maior
}
