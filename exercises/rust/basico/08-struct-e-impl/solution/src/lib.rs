// `struct` agrupa campos nomeados; `impl` dá comportamento a esse tipo, num
// bloco separado da declaração dos dados.
#[derive(Debug)]
pub struct Retangulo {
    pub largura: u32,
    pub altura: u32,
}

impl Retangulo {
    // Uma função ASSOCIADA (sem `self`) funciona como construtor. Chama-se
    // `Retangulo::novo(...)`, com `::`, porque não há instância ainda para
    // chamar com `.`.
    pub fn novo(largura: u32, altura: u32) -> Self {
        Self { largura, altura }
    }

    // `&self` empresta a instância só para leitura — este método não muda
    // largura nem altura, então não precisa de mais que isso.
    pub fn area(&self) -> u32 {
        self.largura * self.altura
    }

    // `&mut self` empresta a instância para ESCRITA. Sem o `mut` aqui, nem
    // `self.largura *= fator` compilaria.
    pub fn escalar(&mut self, fator: u32) {
        self.largura *= fator;
        self.altura *= fator;
    }

    // eh_quadrado devolve true quando largura e altura são iguais.
    pub fn eh_quadrado(&self) -> bool {
        self.largura == self.altura
    }
}

// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let r = Retangulo::novo(3, 4);
    println!("{:?}", r); // Retangulo { largura: 3, altura: 4 }
    println!("{}", r.area()); // 12

    let mut r2 = Retangulo::novo(2, 5);
    r2.escalar(3);
    println!("{} {}", r2.largura, r2.altura); // 6 15
}
