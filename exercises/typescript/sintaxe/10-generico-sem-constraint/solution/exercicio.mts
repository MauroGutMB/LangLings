export function maiorPorTamanho<T extends { length: number }>(a: T, b: T): T {
  return a.length > b.length ? a : b;
}
