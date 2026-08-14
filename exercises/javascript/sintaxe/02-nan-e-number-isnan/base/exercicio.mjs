// Devolva true quando arr contiver algum NaN.
//
// TODO: indexOf compara por igualdade estrita, e isso não é o bastante aqui.
export function temNaN(arr) {
  return arr.indexOf(NaN) !== -1;
}
