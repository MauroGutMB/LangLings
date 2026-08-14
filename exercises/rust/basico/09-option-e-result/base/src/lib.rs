// `Option<T>` é para "pode não ter valor" (`Some(T)` ou `None`).
// `Result<T, E>` é para "pode falhar com um motivo" (`Ok(T)` ou `Err(E)`).
// Nenhum dos dois é `null`: o compilador obriga a tratar os dois casos antes
// de usar o valor de dentro.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let numeros = vec![10, 20, 30];

    // `match` cobre os dois casos de Option de forma exaustiva — esquecer um
    // ramo é erro de compilação, não bug em produção.
    match numeros.get(1) {
        Some(n) => println!("achei {n}"), // achei 20
        None => println!("não achei"),
    }
    match numeros.get(99) {
        Some(n) => println!("achei {n}"),
        None => println!("não achei"), // não achei
    }

    // `unwrap_or` dá um valor padrão sem precisar de `match` — útil quando o
    // caso `None` já tem uma resposta óbvia.
    let decimo: i32 = *numeros.get(9).unwrap_or(&-1);
    println!("{decimo}"); // -1

    // `str::parse` devolve Result: Ok com o número, Err com o motivo da
    // falha. `match` trata os dois lados.
    let entrada = "42";
    match entrada.parse::<i32>() {
        Ok(n) => println!("número: {n}"), // número: 42
        Err(_) => println!("não é um número"),
    }
    let invalida = "abc";
    match invalida.parse::<i32>() {
        Ok(n) => println!("número: {n}"),
        Err(_) => println!("não é um número"), // não é um número
    }

    // `?` propaga um Err para fora da função no ponto exato em que ele
    // aparece — sem `if let Err(e) = ... { return Err(e) }` escrito à mão.
    // `map_err` converte o tipo do erro ANTES do `?`, porque o erro de
    // parse (ParseIntError) não é o mesmo tipo que esta função devolve.
    fn dobro_do_texto(s: &str) -> Result<i32, String> {
        let n: i32 = s.parse().map_err(|_| format!("'{s}' não é um número"))?;
        Ok(n * 2)
    }
    println!("{:?}", dobro_do_texto("21")); // Ok(42)
    println!("{:?}", dobro_do_texto("xyz")); // Err("'xyz' não é um número")
}

// SUA VEZ
//
// Converta a e b para i32 e devolva a soma como Ok. Se algum dos dois não
// for um número válido, devolva Err com qualquer mensagem.
pub fn somar_textos(a: &str, b: &str) -> Result<i32, String> {
    Ok(0) // <- troque isto
}
