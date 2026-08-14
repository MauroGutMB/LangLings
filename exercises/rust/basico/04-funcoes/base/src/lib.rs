// O corpo de uma função é uma sequência de expressões. A ÚLTIMA delas, sem
// ponto e vírgula, é o valor devolvido — não existe um `return` implícito
// escondido, é literalmente o valor daquela expressão.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    // Sem `return`, sem ponto e vírgula na última linha: `a + b` É o valor.
    fn soma(a: i32, b: i32) -> i32 {
        a + b
    }
    println!("{}", soma(2, 3)); // 5

    // `return` continua existindo, mas é para sair ANTES do fim — uma guarda,
    // não o caminho normal.
    fn dividir(a: i32, b: i32) -> i32 {
        if b == 0 {
            return 0; // saída antecipada: aqui o `;` é obrigatório
        }
        a / b // caminho normal: expressão final, sem `;`
    }
    println!("{}", dividir(10, 2)); // 5
    println!("{}", dividir(10, 0)); // 0

    // `if`/`else` é uma expressão: os dois ramos precisam devolver o mesmo
    // tipo, e o valor escolhido vira o valor do `let`.
    let a = 7;
    let b = 12;
    let maior = if a > b { a } else { b };
    println!("{maior}"); // 12

    // Um bloco `{ }` também é uma expressão. A última linha sem `;` é o valor
    // do bloco inteiro — o resto de dentro é só preparação.
    let dobro_do_maior = {
        let m = if a > b { a } else { b };
        m * 2
    };
    println!("{dobro_do_maior}"); // 24

    // Uma função sem seta de retorno devolve `()`, o tipo unit — usada só
    // pelo efeito colateral, não pelo valor.
    fn avisar(nome: &str) {
        println!("oi, {nome}"); // oi, Ana
    }
    avisar("Ana");
}

// SUA VEZ
//
// Devolva o maior entre a, b e c, sem usar `return`.
// maior_de_tres(3, 9, 5) é 9.
pub fn maior_de_tres(a: i32, b: i32, c: i32) -> i32 {
    0 // <- troque isto
}
