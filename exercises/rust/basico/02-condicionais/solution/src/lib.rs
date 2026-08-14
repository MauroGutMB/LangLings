// Em Rust quase tudo é expressão: um `if` e um `match` produzem valor, e é
// isso que dispensa a variável temporária que outras linguagens exigem.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let n = 7;

    // A condição precisa ser bool. `if n` não compila: um inteiro não vira
    // condição sozinho, ao contrário do que acontece em C ou Python.
    if n % 2 == 0 {
        println!("par");
    } else if n % 3 == 0 {
        println!("múltiplo de 3");
    } else {
        println!("nem par nem múltiplo de 3"); // este
    }

    // `if` como expressão. Os dois ramos precisam ter o MESMO tipo, porque o
    // tipo de `rotulo` tem que ser decidido em tempo de compilação.
    let rotulo = if n % 2 == 0 { "par" } else { "ímpar" };
    println!("{rotulo}"); // ímpar

    // O tipo de um literal de texto é &'static str: uma fatia emprestada que
    // vive pelo programa inteiro. Por isso ela pode ser devolvida de uma
    // função sem que ninguém precise ser dono dela.
    let fixo: &'static str = "sempre aqui";
    println!("{fixo}"); // sempre aqui

    // `match` compara o valor contra padrões, de cima para baixo, e para no
    // primeiro que casa. Ele é EXAUSTIVO: se algum valor possível ficar de
    // fora, o programa nem compila.
    let dia = 3;
    let nome = match dia {
        1 => "domingo",
        2 => "segunda",
        3 => "terça", // este
        _ => "outro", // o curinga cobre todo o resto
    };
    println!("{nome}"); // terça

    // Padrões podem ser intervalos inclusivos e alternativas com `|`.
    let temperatura = 24;
    let clima = match temperatura {
        i32::MIN..=9 => "frio",
        10..=25 => "ameno", // este
        _ => "quente",
    };
    println!("{clima}"); // ameno

    let letra = 'k';
    let tipo = match letra {
        'a' | 'e' | 'i' | 'o' | 'u' => "vogal",
        'a'..='z' => "consoante", // este
        _ => "não é letra minúscula",
    };
    println!("{tipo}"); // consoante
}

// classificar traduz a nota numérica no rótulo correspondente.
//
// O `match` diz na estrutura o que uma cadeia de `if` só diria no comentário:
// as faixas são alternativas mutuamente exclusivas de uma coisa só. O `_` no
// fim é obrigatório porque u32 tem valores acima de 100, e o compilador não
// aceita deixá-los sem resposta.
pub fn classificar(nota: u32) -> &'static str {
    match nota {
        0..=59 => "reprovado",
        60..=79 => "aprovado",
        _ => "destaque",
    }
}
