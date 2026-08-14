// TODO: esta única assinatura diz "entra string ou number, sai string ou
// number" — verdade, mas incompleta: ela não capta que uma entrada string
// SEMPRE devolve number, e uma entrada number SEMPRE devolve string. Quem
// chama com um tipo concreto perde essa precisão e o teste não compila.
export function converterTipo(valor: string | number): string | number {
  if (typeof valor === "string") {
    return Number(valor);
  }
  return valor.toString();
}
