interface Preco {
  dolar: number;
  euro: number;
  real: number;
}

// TODO: Object.keys(precos) tem tipo string[] — SEMPRE, mesmo sabendo que
// precos é um Preco e só pode ter essas três chaves. TypeScript não estreita
// esse retorno para keyof Preco, então indexar precos com uma chave string
// qualquer não compila.
export function somarTodosOsPrecos(precos: Preco): number {
  let total = 0;
  for (const chave of Object.keys(precos)) {
    total += precos[chave];
  }
  return total;
}
