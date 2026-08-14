export function somarSeNumero(valor: unknown, outro: number): number {
  if (typeof valor === "number") {
    return valor + outro;
  }
  return outro;
}
