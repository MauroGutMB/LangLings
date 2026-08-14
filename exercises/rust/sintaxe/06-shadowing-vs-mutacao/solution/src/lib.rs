// Os dois `let texto` estão no mesmo nível — cada um sombreia o anterior até
// o fim da função, sem um bloco extra apagando o shadow antes da hora.
pub fn normalizar(texto: &str) -> String {
    let texto = texto.trim();
    let texto = texto.to_lowercase();
    texto
}
