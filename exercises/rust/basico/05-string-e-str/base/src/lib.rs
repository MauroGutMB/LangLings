// `String` é dona do texto: alocada no heap, pode crescer, pode ser passada
// adiante. `&str` é uma VISÃO emprestada de texto — de um literal, que vive
// no binário, ou de um pedaço de dentro de uma `String`.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let literal: &str = "café"; // vive no binário, nunca é dono
    let dona: String = String::from("café"); // cópia própria, no heap
    println!("{literal} {dona}"); // café café

    // `&dona` empresta uma visão da String — não move, não copia o conteúdo.
    // A coerção de `&String` para `&str` é automática aqui.
    let emprestado: &str = &dona;
    println!("{emprestado}"); // café

    // Uma função que só LÊ texto deve pedir `&str`: assim aceita tanto um
    // literal quanto uma String emprestada, sem forçar quem chama a alocar.
    fn tamanho_em_bytes(s: &str) -> usize {
        s.len()
    }
    println!("{}", tamanho_em_bytes(literal)); // 5, não 4: 'é' ocupa 2 bytes em UTF-8
    println!("{}", tamanho_em_bytes(&dona)); // 5

    // Concatenar cria uma String nova; o `&str` original não muda.
    let saudacao = String::from("olá, ") + literal;
    println!("{saudacao}"); // olá, café

    // `push_str` empresta um &str e anexa ao final de uma String existente.
    let mut mensagem = String::from("status: ");
    mensagem.push_str("ok");
    println!("{mensagem}"); // status: ok

    // Fatiar um &str por índice de BYTE, não de caractere — por isso fatiar
    // no meio de um caractere multibyte gera pânico, e "café" cortado em 3
    // pega só o 'a' e metade do 'f'... então aqui cortamos num limite seguro.
    let apenas_cafe: &str = &dona[0..3];
    println!("{apenas_cafe}"); // caf
}

// SUA VEZ
//
// Devolva uma String nova com s em maiúsculas seguido de "!".
// gritar("oi") é "OI!".
pub fn gritar(s: &str) -> String {
    String::new() // <- troque isto
}
