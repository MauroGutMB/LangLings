// Devolva valor quando ele não for null nem undefined; devolva padrao caso
// contrário. 0 e '' contam como informados.
export function comPadrao(valor, padrao) {
  return valor ?? padrao;
}
