// TODO: isto não compila. `s.parse::<i32>()` erra com ParseIntError, e não
// existe conversão automática de ParseIntError para String — o `?` sozinho
// não sabe como propagar esse erro no tipo que esta função promete devolver.
pub fn dobrar_texto(s: &str) -> Result<i32, String> {
    let n: i32 = s.parse()?;
    Ok(n * 2)
}
