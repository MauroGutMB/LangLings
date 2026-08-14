// Map e Set são as coleções que o objeto e o array não cobrem bem: um Map
// aceita chave de qualquer tipo e mantém a ordem de inserção; um Set guarda
// cada valor uma única vez.
function exemplos() {
  const idades = new Map();

  idades.set('ana', 30); // insere ou sobrescreve
  idades.set('bruno', 25);
  console.log(idades.get('ana')); // 30
  console.log(idades.size); // 2

  // Ler uma chave ausente devolve undefined. Note a diferença para o Go ou
  // para uma contagem ingênua: não existe "zero automático" aqui.
  console.log(idades.get('zoe')); // undefined
  console.log(idades.has('zoe')); // false  <- a pergunta explícita

  idades.delete('ana');
  console.log(idades.size); // 1

  // O percurso segue a ordem de inserção, e cada volta entrega [chave, valor].
  const contagem = new Map();
  for (const p of ['a', 'b', 'a']) {
    contagem.set(p, (contagem.get(p) ?? 0) + 1); // ?? 0 cobre a primeira vez
  }
  for (const [chave, valor] of contagem) {
    console.log(chave, valor); // a 2 / b 1
  }

  // Diferente de um objeto, a chave de um Map pode ser qualquer valor.
  const porNumero = new Map([[1, 'um']]); // dá para construir já com pares
  console.log(porNumero.get(1)); // um
  console.log(porNumero.get('1')); // undefined  <- 1 e '1' são chaves distintas

  // Um Set ignora repetições silenciosamente.
  const vistos = new Set();
  vistos.add('go');
  vistos.add('go');
  console.log(vistos.size, vistos.has('go')); // 1 true

  // Daí o idioma para tirar duplicados de um array: joga tudo num Set e
  // espalha de volta num array, preservando a ordem da primeira aparição.
  console.log([...new Set([3, 1, 3, 2, 1])]); // [ 3, 1, 2 ]
}

// contarOcorrencias conta quantas vezes cada palavra aparece.
//
// O `?? 0` é o que substitui o zero automático que um Map não tem: sem ele, a
// primeira ocorrência somaria 1 a undefined e gravaria NaN, que continuaria
// NaN em todas as ocorrências seguintes.
export function contarOcorrencias(palavras) {
  const contagem = new Map();

  for (const palavra of palavras) {
    contagem.set(palavra, (contagem.get(palavra) ?? 0) + 1);
  }

  return contagem;
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
