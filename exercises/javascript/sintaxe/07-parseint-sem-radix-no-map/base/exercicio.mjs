// Converta cada string de arr para um número inteiro em base 10, na mesma
// ordem.
//
// TODO: map chama o callback com (elemento, índice, array); parseInt usa o
// segundo argumento como base numérica, não como algo a ignorar.
export function paraNumeros(arr) {
  return arr.map(parseInt);
}
