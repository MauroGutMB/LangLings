export function primeiraLetra(palavras: string[], indice: number): string {
  const palavra = palavras[indice];
  if (palavra === undefined) {
    return "";
  }
  return palavra.charAt(0).toUpperCase();
}
