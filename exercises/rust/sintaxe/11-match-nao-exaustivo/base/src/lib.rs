// TODO: isto não compila. O match não cobre `Estado::Suspenso` — o
// compilador recusa um match que deixa uma variante de fora.
#[derive(Debug, PartialEq)]
pub enum Estado {
    Ativo,
    Inativo,
    Suspenso,
}

pub fn pode_logar(estado: &Estado) -> bool {
    match estado {
        Estado::Ativo => true,
        Estado::Inativo => false,
    }
}
