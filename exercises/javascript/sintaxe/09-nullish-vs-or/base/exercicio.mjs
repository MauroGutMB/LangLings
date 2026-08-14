// Devolva valor quando ele não for null nem undefined; devolva padrao caso
// contrário. 0 e '' contam como informados.
//
// TODO: || descarta todo valor falsy, não só null e undefined.
export function comPadrao(valor, padrao) {
  return valor || padrao;
}
