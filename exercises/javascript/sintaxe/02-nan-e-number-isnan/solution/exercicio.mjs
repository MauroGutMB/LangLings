// Devolva true quando arr contiver algum NaN.
export function temNaN(arr) {
  return arr.some((x) => Number.isNaN(x));
}
