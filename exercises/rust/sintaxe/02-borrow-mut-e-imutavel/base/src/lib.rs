// TODO: isto não compila. O compilador vai reclamar de emprestar `v` como
// mutável enquanto `maior` — uma referência para dentro de `v` — ainda está
// viva. A lógica está certa; o que precisa mudar é o tipo de `maior`.
pub fn maior_e_dobro(v: &mut Vec<i32>) -> i32 {
    let maior = v.iter().max().unwrap();
    v.push(*maior * 2);
    *maior
}
