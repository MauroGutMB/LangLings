// Uma função em JavaScript é um valor: ela pode ser guardada numa variável,
// passada como argumento e devolvida por outra função. As três sintaxes abaixo
// produzem funções equivalentes para o que este exercício faz.
function exemplos() {
  // Forma declarada.
  function dobro(n) {
    return n * 2;
  }

  // Arrow function guardada numa const. Com corpo de uma expressão só, as
  // chaves e o return somem: o valor da expressão já é o retorno.
  const triplo = (n) => n * 3;

  console.log(dobro(5), triplo(5)); // 10 15

  // Parâmetro com valor padrão: usado quando o argumento não vem — ou seja,
  // quando ele é undefined. Passar 0 é passar um valor, e o padrão não entra.
  function saudar(nome, saudacao = 'Olá') {
    return `${saudacao}, ${nome}!`;
  }
  console.log(saudar('Ana')); // Olá, Ana!
  console.log(saudar('Ana', 'Bom dia')); // Bom dia, Ana!

  // Rest: junta todos os argumentos restantes num array de verdade.
  function somar(...numeros) {
    let total = 0;
    for (const n of numeros) total += n;
    return total;
  }
  console.log(somar()); // 0
  console.log(somar(1, 2, 3)); // 6

  // Função como valor: map recebe outra função e a aplica a cada elemento.
  console.log([1, 2, 3].map(dobro)); // [ 2, 4, 6 ]

  // Chamar sem um argumento que não tem padrão não é erro: ele vale undefined.
  console.log(dobro()); // NaN  <- undefined * 2
}

// repetir emenda texto consigo mesmo `vezes` vezes.
//
// O padrão fica na assinatura, e não num `vezes = vezes || 2` no corpo: o ||
// trataria o 0 como ausência e repetiria duas vezes um texto que o chamador
// pediu zero vezes.
export function repetir(texto, vezes = 2) {
  return texto.repeat(vezes);
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
