// mapa é um Map cujas chaves são objetos { id: ... }. Devolva o valor cuja
// chave tem chave.id === id.
export function buscarValor(mapa, id) {
  for (const [chave, valor] of mapa.entries()) {
    if (chave.id === id) return valor;
  }
  return undefined;
}
