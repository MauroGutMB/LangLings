// map_err converte o ParseIntError numa String ANTES do `?` decidir se
// propaga — depois disso os dois lados do Result batem com a assinatura.
pub fn dobrar_texto(s: &str) -> Result<i32, String> {
    let n: i32 = s.parse().map_err(|_| format!("'{s}' não é um número"))?;
    Ok(n * 2)
}
