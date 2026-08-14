// Um array tipado se anota com Tipo[] (ou Array<Tipo>, equivalente). Toda
// posição tem o mesmo tipo, e o tamanho não é fixo.
export function exemplos(): void {
  const numeros: number[] = [10, 20, 30];
  numeros.push(40); // ok: array pode crescer
  console.log(numeros[0], numeros.length); // 10 4

  const nomes: string[] = ["ana", "bruno"];
  console.log(nomes.join(", ")); // ana, bruno

  // Uma tupla é diferente: tamanho FIXO e cada posição tem seu próprio tipo.
  // [number, string] não é o mesmo que (number | string)[] — a ordem e a
  // quantidade importam.
  const coordenada: [number, number] = [10, 20];
  const x = coordenada[0]; // tipo number, especificamente essa posição
  const y = coordenada[1]; // tipo number também, mas é outra posição
  console.log(x, y);

  const par: [string, number] = ["idade", 30];
  const chave = par[0]; // tipo string
  const valor = par[1]; // tipo number
  console.log(chave, valor);

  // Desestruturar uma tupla preserva o tipo de cada posição.
  const [primeiroNumero, segundoNumero] = coordenada;
  console.log(primeiroNumero, segundoNumero);
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Devolva uma tupla com o primeiro e o último elemento de numeros.
// numeros nunca vem vazio.
export function primeiroEUltimo(numeros: number[]): [number, number] {
  return [0, 0]; // <- troque isto
}
