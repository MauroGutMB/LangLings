// Devolva um array com n elementos, todos valendo 0.
//
// TODO: Array(n) cria posições vazias, não posições com 0 dentro — e map não
// visita posição vazia nenhuma.
export function criarZeros(n) {
  return Array(n).map(() => 0);
}
