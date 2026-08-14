// TODO: `as string` afirma para o compilador "confie em mim, isso é uma
// string" — mas não converte nem checa nada em tempo de execução. Quando
// valor não é mesmo uma string, o programa quebra silenciosamente.
export function comprimento(valor: unknown): number {
  return (valor as string).length;
}
