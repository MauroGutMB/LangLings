// TODO: com noUncheckedIndexedAccess ligado (veja tsconfig.json),
// palavras[indice] tem tipo `string | undefined` — mesmo palavras sendo
// string[]. O compilador está avisando que o índice pode estar fora dos
// limites do array, e chamar .charAt direto nesse resultado não compila.
export function primeiraLetra(palavras: string[], indice: number): string {
  return palavras[indice].charAt(0).toUpperCase();
}
