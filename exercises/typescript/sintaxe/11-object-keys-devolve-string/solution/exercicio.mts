interface Preco {
  dolar: number;
  euro: number;
  real: number;
}

export function somarTodosOsPrecos(precos: Preco): number {
  let total = 0;
  for (const chave of Object.keys(precos) as (keyof Preco)[]) {
    total += precos[chave];
  }
  return total;
}
