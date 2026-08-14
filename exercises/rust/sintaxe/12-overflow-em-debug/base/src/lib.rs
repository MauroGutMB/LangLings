// TODO: isto compila e roda para entradas pequenas — e entra em pânico por
// overflow para n >= 13, mesmo o retorno sendo u64. O acumulador é u32, e é
// nele que a multiplicação acontece antes do `as u64` converter no final.
pub fn fatorial(n: u32) -> u64 {
    let mut resultado: u32 = 1;
    for i in 1..=n {
        resultado *= i;
    }
    resultado as u64
}
