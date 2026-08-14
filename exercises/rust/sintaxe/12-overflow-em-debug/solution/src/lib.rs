// O acumulador é u64 desde o início: a multiplicação acontece no tipo largo,
// não só a conversão do resultado final.
pub fn fatorial(n: u32) -> u64 {
    let mut resultado: u64 = 1;
    for i in 1..=n as u64 {
        resultado *= i;
    }
    resultado
}
