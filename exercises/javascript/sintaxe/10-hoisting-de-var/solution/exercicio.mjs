// Devolva um array com n funções; a função na posição i deve devolver i
// quando chamada.
export function criarContadores(n) {
  const funcoes = [];
  for (let i = 0; i < n; i++) {
    funcoes.push(() => i);
  }
  return funcoes;
}
