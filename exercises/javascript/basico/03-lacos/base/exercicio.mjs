// Repetir em JavaScript é escolher entre percorrer uma coleção e repetir
// enquanto uma condição valer. As duas formas fazem o mesmo trabalho; o que
// muda é o quanto de contabilidade sobra para você.
function exemplos() {
  const notas = [7, 8, 10];

  // for clássico: você controla o índice. Vale quando o índice importa.
  for (let i = 0; i < notas.length; i++) {
    console.log(i, notas[i]); // 0 7 / 1 8 / 2 10
  }

  // for...of entrega o VALOR de cada posição — sem índice, sem length, sem
  // errar o limite. É a forma padrão quando só os valores interessam.
  let total = 0;
  for (const nota of notas) {
    total += nota;
  }
  console.log(total); // 25

  // Quando o índice também importa, entries() devolve os dois.
  for (const [i, nota] of notas.entries()) {
    console.log(`posição ${i} vale ${nota}`); // posição 0 vale 7 ...
  }

  // while repete enquanto a condição for verdadeira. Aqui a condição depende
  // de algo que muda no corpo — esquecer de mudar é o laço infinito clássico.
  let restante = 10;
  while (restante > 0) {
    restante -= 3;
  }
  console.log(restante); // -2

  // continue pula para a próxima volta; break abandona o laço inteiro.
  const numeros = [1, 2, 3, 4, 5, 6];
  let primeiroGrande = null;
  for (const n of numeros) {
    if (n % 2 !== 0) continue; // ímpar não interessa
    if (n > 3) {
      primeiroGrande = n;
      break; // achou o que queria, não precisa ver o resto
    }
  }
  console.log(primeiroGrande); // 4
}

// SUA VEZ
//
// Devolva a soma apenas dos números pares de numeros. Array vazio soma 0.
export function somarPares(numeros) {
  return -1; // <- troque isto
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
