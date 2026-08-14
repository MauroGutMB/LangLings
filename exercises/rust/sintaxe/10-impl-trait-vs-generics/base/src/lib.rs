// TODO: isto não compila. `impl Trait` no retorno exige UM tipo concreto só;
// os dois ramos do `if` devolvem tipos diferentes (Range<i32> e
// std::vec::IntoIter<i32>), mesmo os dois implementando Iterator<Item = i32>.
pub fn escolher(usar_intervalo: bool) -> impl Iterator<Item = i32> {
    if usar_intervalo {
        1..5
    } else {
        vec![10, 20].into_iter()
    }
}
