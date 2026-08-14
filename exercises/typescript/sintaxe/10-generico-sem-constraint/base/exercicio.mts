// TODO: T aqui não tem constraint nenhuma — pode ser um number, um boolean,
// qualquer coisa. Nem todo tipo tem um campo length, e o compilador recusa
// o acesso até T ser restrito a algo que garanta esse campo.
export function maiorPorTamanho<T>(a: T, b: T): T {
  return a.length > b.length ? a : b;
}
