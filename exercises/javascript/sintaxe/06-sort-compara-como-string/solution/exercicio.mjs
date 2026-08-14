// Devolva uma cópia de arr ordenada em ordem numérica crescente, sem alterar
// o array recebido.
export function ordenarNumeros(arr) {
  return [...arr].sort((a, b) => a - b);
}
