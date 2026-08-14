// `Box<dyn Iterator<Item = i32>>` apaga o tipo concreto por trás de um
// ponteiro com despacho em tempo de execução — os dois ramos podem devolver
// iteradores de tipos diferentes, porque o chamador só enxerga o trait.
pub fn escolher(usar_intervalo: bool) -> Box<dyn Iterator<Item = i32>> {
    if usar_intervalo {
        Box::new(1..5)
    } else {
        Box::new(vec![10, 20].into_iter())
    }
}
