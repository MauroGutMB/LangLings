// Um array é uma lista ordenada e indexada a partir do zero. Ele cresce e
// encolhe sozinho, e guarda valores de qualquer tipo — inclusive misturados.
function exemplos() {
  const notas = [7, 8, 10];

  console.log(notas[0], notas.length); // 7 3
  console.log(notas[99]); // undefined  <- fora do array não é erro

  // push acrescenta no fim e pop tira do fim. Os dois ALTERAM o array.
  notas.push(6);
  console.log(notas); // [ 7, 8, 10, 6 ]
  console.log(notas.pop(), notas); // 6 [ 7, 8, 10 ]

  // Busca.
  console.log(notas.includes(8)); // true
  console.log(notas.indexOf(10)); // 2
  console.log(notas.indexOf(99)); // -1  <- não achou

  // map devolve um array NOVO, do mesmo tamanho, com cada elemento
  // transformado pela função.
  console.log(notas.map((n) => n * 10)); // [ 70, 80, 100 ]

  // filter devolve um array novo só com os elementos cuja função deu
  // verdadeiro. O tamanho muda; os elementos, não.
  console.log(notas.filter((n) => n >= 8)); // [ 8, 10 ]

  // reduce reduz o array inteiro a um único valor. O segundo argumento é o
  // valor inicial do acumulador — e é ele que faz o array vazio devolver 0 em
  // vez de estourar.
  const soma = notas.reduce((acumulado, n) => acumulado + n, 0);
  console.log(soma); // 25
  console.log([].reduce((a, n) => a + n, 0)); // 0

  // Os três devolvem arrays/valores novos e podem ser encadeados. O array de
  // origem continua o mesmo.
  console.log(notas.filter((n) => n >= 8).map((n) => `nota ${n}`)); // [ 'nota 8', 'nota 10' ]
  console.log(notas); // [ 7, 8, 10 ]
}

// media devolve a média aritmética dos números.
//
// O array vazio sai por cima, antes da divisão: é o único caminho em que a
// conta produziria NaN, e NaN se propaga silenciosamente por tudo que vier
// depois. O reduce não encosta no array de origem, então nada é alterado.
export function media(numeros) {
  if (numeros.length === 0) return 0;

  const soma = numeros.reduce((acumulado, n) => acumulado + n, 0);
  return soma / numeros.length;
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
