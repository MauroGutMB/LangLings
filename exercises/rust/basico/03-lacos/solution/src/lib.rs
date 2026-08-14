// Rust tem três laços, e a escolha entre eles é sobre o que você sabe de
// antemão: `loop` quando a saída é decidida no meio, `while` quando existe uma
// condição de continuação, `for` quando existe uma sequência a percorrer.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    // `loop` repete para sempre até um `break`. Como ele é uma expressão, o
    // valor passado ao `break` é o valor do laço inteiro — o que evita a
    // variável temporária declarada só para escapar de dentro.
    let mut tentativa = 0;
    let achado = loop {
        tentativa += 1;
        if tentativa * tentativa > 50 {
            break tentativa;
        }
    };
    println!("{achado}"); // 8

    // `while` testa a condição antes de cada volta.
    let mut restante = 3;
    while restante > 0 {
        println!("faltam {restante}"); // 3, 2, 1
        restante -= 1;
    }

    // `for` sobre um intervalo. `0..4` exclui o 4; `0..=4` inclui.
    let mut soma = 0;
    for i in 0..4 {
        soma += i;
    }
    println!("{soma}"); // 6  (0+1+2+3)

    let mut soma_inclusiva = 0;
    for i in 0..=4 {
        soma_inclusiva += i;
    }
    println!("{soma_inclusiva}"); // 10  (0+1+2+3+4)

    // `for` sobre uma coleção. O `&` é o que faz o laço EMPRESTAR o vetor em
    // vez de consumi-lo: sem ele, `frutas` seria movido para dentro do laço e
    // não existiria mais depois.
    let frutas = vec!["maçã", "pera", "uva"];
    for f in &frutas {
        println!("{f}"); // maçã, pera, uva
    }
    println!("{}", frutas.len()); // 3  <- ainda dá para usar

    // `enumerate` acrescenta o índice a qualquer iterador.
    for (i, f) in frutas.iter().enumerate() {
        println!("{i}: {f}"); // 0: maçã, 1: pera, 2: uva
    }

    // `continue` pula para a próxima volta; `break` encerra o laço.
    for i in 0..10 {
        if i % 2 == 1 {
            continue;
        }
        if i > 5 {
            break;
        }
        println!("{i}"); // 0, 2, 4
    }
}

// SUA VEZ
//
// Devolva a soma dos números pares de 0 até `ate`, inclusive.
// soma_pares(10) é 30.
pub fn soma_pares(ate: u32) -> u32 {
    // O acumulador vive fora do laço para sobreviver às iterações, e é `mut`
    // porque muda a cada volta. O intervalo é inclusivo (`..=`): quando `ate`
    // é par, ele próprio entra na soma.
    let mut soma = 0;
    for i in 0..=ate {
        if i % 2 == 0 {
            soma += i;
        }
    }
    soma
}
