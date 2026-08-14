// mapa é um Map cujas chaves são objetos { id: ... }. Devolva o valor cuja
// chave tem chave.id === id.
//
// TODO: um objeto criado agora nunca é === a um objeto criado antes, mesmo
// com os mesmos campos — e Map.get compara chaves por ===.
export function buscarValor(mapa, id) {
  return mapa.get({ id });
}
