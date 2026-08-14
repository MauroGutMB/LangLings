// Em Rust toda ligação criada com `let` é imutável por padrão: o nome fica
// preso ao valor. Mutar é a exceção, e a exceção é escrita à mão com `mut`.
//
// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let nome = "Ana"; // &str, inferido do literal
    let idade = 30; // i32, o inteiro padrão
    println!("{nome} {idade}"); // Ana 30

    // `nome = "Bruno";` aqui seria erro de compilação: a ligação é imutável.
    // Com `mut`, a reatribuição passa a ser permitida — mas o TIPO continua
    // fixo, então só cabe outro i32 ali.
    let mut contador = 0;
    contador += 1;
    contador = contador * 10;
    println!("{contador}"); // 10

    // A anotação de tipo é obrigatória quando o valor sozinho não decide.
    // Aqui o literal 3 poderia ser i8, u64, i32...; a anotação escolhe.
    let pequeno: u8 = 3;
    let grande: i64 = 9_000_000_000;
    println!("{pequeno} {grande}"); // 3 9000000000

    // Os tipos escalares: inteiro, ponto flutuante, booleano e caractere.
    // `char` é um ponto de código Unicode inteiro, não um byte.
    let ativo: bool = true;
    let inicial: char = 'A';
    let taxa: f64 = 0.5;
    println!("{ativo} {inicial} {taxa}"); // true A 0.5

    // Conversão entre números é sempre explícita, com `as`. Repare que ela vem
    // ANTES da divisão — se viesse depois, a divisão inteira já teria
    // descartado a fração.
    let metade = idade as f64 / 2.0;
    println!("{metade}"); // 15
    println!("{}", (idade / 4) as f64); // 7  <- fração perdida na divisão inteira
    println!("{}", idade as f64 / 4.0); // 7.5

    // `const` é resolvida em tempo de compilação e exige o tipo escrito.
    const MAIOR_IDADE: i32 = 18;
    println!("{}", idade >= MAIOR_IDADE); // true
}

// media devolve a média de a e b.
//
// A soma acontece em i32 e só então vira f64: converter os dois lados antes de
// somar daria o mesmo resultado, mas é a DIVISÃO que precisa acontecer em
// ponto flutuante para a fração sobreviver.
pub fn media(a: i32, b: i32) -> f64 {
    (a + b) as f64 / 2.0
}
