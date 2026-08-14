// Converta cada string de arr para um número inteiro em base 10, na mesma
// ordem.
export function paraNumeros(arr) {
  return arr.map((s) => parseInt(s, 10));
}
