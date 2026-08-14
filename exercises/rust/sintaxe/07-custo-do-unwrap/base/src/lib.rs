// TODO: isto compila e funciona quando existe um positivo — e entra em
// pânico quando não existe, mesmo a função já devolvendo Option<i32>.
pub fn primeiro_positivo(v: &[i32]) -> Option<i32> {
    Some(*v.iter().find(|&&x| x > 0).unwrap())
}
