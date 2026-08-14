export function comprimento(valor: unknown): number {
  if (typeof valor === "string") {
    return valor.length;
  }
  return 0;
}
