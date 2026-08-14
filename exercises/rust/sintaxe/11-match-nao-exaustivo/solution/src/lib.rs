// As três variantes de Estado, todas cobertas: uma conta suspensa não pode
// logar, do mesmo jeito que uma inativa.
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
        Estado::Suspenso => false,
    }
}
