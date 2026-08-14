// Devolva um array com n funções; a função na posição i deve devolver i
// quando chamada.
//
// TODO: var é local à função inteira, não ao corpo do for — todas as
// closures acabam compartilhando a mesma variável.
export function criarContadores(n) {
  const funcoes = [];
  for (var i = 0; i < n; i++) {
    funcoes.push(() => i);
  }
  return funcoes;
}
