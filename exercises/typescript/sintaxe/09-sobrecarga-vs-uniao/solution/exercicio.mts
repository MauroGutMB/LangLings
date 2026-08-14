export function converterTipo(valor: string): number;
export function converterTipo(valor: number): string;
export function converterTipo(valor: string | number): string | number {
  if (typeof valor === "string") {
    return Number(valor);
  }
  return valor.toString();
}
