// Este exercício não ensina nada novo: junta struct+impl, Vec, HashMap e
// Option, que já apareceram separados nos exercícios anteriores.
use std::collections::HashMap;

#[derive(Debug, Clone, PartialEq)]
pub struct Produto {
    pub nome: String,
    pub preco: f64,
    pub quantidade: u32,
}

impl Produto {
    pub fn novo(nome: &str, preco: f64, quantidade: u32) -> Self {
        Self {
            nome: nome.to_string(),
            preco,
            quantidade,
        }
    }

    // Quanto vale, no total, o estoque deste único produto.
    pub fn valor_total(&self) -> f64 {
        self.preco * self.quantidade as f64
    }
}

// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let estoque = vec![
        Produto::novo("caneta", 2.5, 100),
        Produto::novo("caderno", 8.0, 30),
        Produto::novo("mochila", 90.0, 5),
    ];

    // Vec de structs, percorrido com iterador: soma o valor_total de cada um.
    let total: f64 = estoque.iter().map(|p| p.valor_total()).sum();
    println!("{total}"); // 940 (250 + 240 + 450)

    // HashMap construído a partir do Vec, para consulta por nome em vez de
    // percorrer a lista inteira de novo a cada busca.
    let mut precos: HashMap<String, f64> = HashMap::new();
    for produto in &estoque {
        precos.insert(produto.nome.clone(), produto.preco);
    }
    println!("{:?}", precos.get("caderno")); // Some(8.0)
    println!("{:?}", precos.get("tablet")); // None, não está no estoque

    // Option encadeado com unwrap_or: preço padrão quando o produto não existe.
    let preco_tablet = precos.get("tablet").copied().unwrap_or(0.0);
    println!("{preco_tablet}"); // 0
}

// mais_caro devolve o produto de maior preço, ou None se a lista for vazia.
//
// `iter().max_by(...)` já resolve "lista vazia" devolvendo None sozinho —
// não precisa do `if is_empty` que a dica sugeria como caminho alternativo.
// `partial_cmp` é necessário porque f64 não tem Ord total (por causa de NaN).
pub fn mais_caro(produtos: &[Produto]) -> Option<&Produto> {
    produtos
        .iter()
        .max_by(|a, b| a.preco.partial_cmp(&b.preco).unwrap())
}
