// `Vec<T>` é um vetor de tamanho dinâmico, alocado no heap. `vec![...]` cria
// um já populado; `Vec::new()` cria um vazio para crescer com `push`.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let mut numeros: Vec<i32> = Vec::new();
    numeros.push(10);
    numeros.push(20);
    numeros.push(30);
    println!("{numeros:?}"); // [10, 20, 30]

    let literal = vec![1, 2, 3, 4, 5];
    println!("{literal:?}"); // [1, 2, 3, 4, 5]

    // `[]` acessa por índice e PANICA se estiver fora do intervalo — por isso
    // só serve quando o índice já é conhecido como válido.
    println!("{}", literal[0]); // 1

    // `get` devolve `Option<&T>`: `Some` quando o índice existe, `None`
    // quando não — sem pânico, e o chamador decide o que fazer com o `None`.
    println!("{:?}", literal.get(10)); // None
    println!("{:?}", literal.get(2)); // Some(3)

    // `for` sobre `&vetor` empresta cada elemento sem consumir o vetor.
    let mut soma = 0;
    for n in &literal {
        soma += n;
    }
    println!("{soma}"); // 15

    // O mesmo com iteradores: `sum` e `max` evitam o laço escrito à mão.
    println!("{}", literal.iter().sum::<i32>()); // 15
    println!("{:?}", literal.iter().max()); // Some(5)

    // `filter` + `collect` monta um Vec novo só com o que passa na condição.
    // `map` transformaria cada elemento; aqui a tarefa é ESCOLHER, não mudar.
    let maiores_que_dois: Vec<&i32> = literal.iter().filter(|n| **n > 2).collect();
    println!("{maiores_que_dois:?}"); // [3, 4, 5]
}

// SUA VEZ
//
// Devolva um Vec novo só com os números pares de ns, na mesma ordem.
// pares(&[1, 2, 3, 4]) é [2, 4].
pub fn pares(ns: &[i32]) -> Vec<i32> {
    Vec::new() // <- troque isto
}
