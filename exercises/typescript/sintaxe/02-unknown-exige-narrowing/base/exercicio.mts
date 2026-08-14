// TODO: valor é unknown — usá-lo direto num + é erro de compilação até que
// o tipo seja estreitado (narrowing) com uma checagem, como typeof.
export function somarSeNumero(valor: unknown, outro: number): number {
  return valor + outro;
}
