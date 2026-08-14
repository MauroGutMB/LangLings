// `iter_mut()` dá acesso mutável a cada elemento DENTRO do vetor original —
// nenhuma cópia envolvida, e a mutação através de `*x` acontece no lugar.
pub fn dobrar_todos(v: &mut Vec<i32>) {
    for x in v.iter_mut() {
        *x *= 2;
    }
}
