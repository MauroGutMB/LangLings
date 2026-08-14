// `HashMap<K, V>` associa chaves a valores. Diferente do `Vec`, não tem
// ordem garantida — por isso os exemplos abaixo consultam chaves específicas
// em vez de imprimir o mapa inteiro.
use std::collections::HashMap;

// Rode `cargo test --offline -- --nocapture` no shell do container ([s]) para
// ver a saída daqui. O harness NÃO chama esta função: assim a saída dos
// exemplos não se mistura com o resultado dos testes.
pub fn exemplos() {
    let mut idades: HashMap<String, u32> = HashMap::new();
    idades.insert(String::from("Ana"), 30);
    idades.insert(String::from("Bruno"), 25);
    println!("{:?}", idades.get("Ana")); // Some(30)

    // `insert` na mesma chave SOBRESCREVE o valor anterior; não acumula.
    idades.insert(String::from("Ana"), 31);
    println!("{:?}", idades.get("Ana")); // Some(31)

    // `get` devolve `Option<&V>` — nunca panica com uma chave ausente.
    println!("{:?}", idades.get("Carla")); // None

    println!("{}", idades.contains_key("Bruno")); // true
    println!("{}", idades.contains_key("Carla")); // false

    // `entry().or_insert()` só usa o valor padrão quando a chave ainda não
    // existe; numa chave já presente, devolve o valor que já estava lá.
    *idades.entry(String::from("Carla")).or_insert(0) += 1;
    *idades.entry(String::from("Carla")).or_insert(0) += 1;
    println!("{:?}", idades.get("Carla")); // Some(2)

    // O mesmo padrão é como se conta ocorrências: a chave nasce em 0 na
    // primeira vez e cada `+= 1` seguinte soma na MESMA entrada.
    let letras = ['a', 'b', 'a', 'c', 'b', 'a'];
    let mut contagem: HashMap<char, u32> = HashMap::new();
    for letra in letras {
        *contagem.entry(letra).or_insert(0) += 1;
    }
    println!("{:?}", contagem.get(&'a')); // Some(3)
    println!("{:?}", contagem.get(&'b')); // Some(2)
    println!("{:?}", contagem.get(&'c')); // Some(1)
}

// SUA VEZ
//
// Devolva um HashMap com quantas vezes cada palavra de texto aparece.
// As palavras são separadas por espaço.
// contar_palavras("a b a") tem "a" valendo 2 e "b" valendo 1.
pub fn contar_palavras(texto: &str) -> HashMap<String, u32> {
    HashMap::new() // <- troque isto
}
