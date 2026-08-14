// TODO: isto compila e roda, mas não faz o que o nome promete. `v.clone()`
// dobra uma CÓPIA; o `v` original que quem chamou passou continua o mesmo.
pub fn dobrar_todos(v: &mut Vec<i32>) {
    for mut x in v.clone() {
        x *= 2;
    }
}
