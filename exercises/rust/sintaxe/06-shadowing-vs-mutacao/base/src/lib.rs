// TODO: isto compila e roda, mas devolve o texto SEM minúsculas. O segundo
// `let texto` está dentro de um bloco `{ }` — o shadow morre ali, e o
// `texto` usado no retorno é o de fora, só com trim.
pub fn normalizar(texto: &str) -> String {
    let texto = texto.trim();
    {
        let texto = texto.to_lowercase();
        let _ = texto; // evita o warning de variável não usada
    }
    texto.to_string()
}
