// `.copied()` transforma o `Option<&i32>` de `find` em `Option<i32>` sem
// desembrulhar nada — o `None` sobrevive até virar o retorno da função.
pub fn primeiro_positivo(v: &[i32]) -> Option<i32> {
    v.iter().find(|&&x| x > 0).copied()
}
