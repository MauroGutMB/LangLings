// <T> declara um parâmetro de tipo: T fica "em aberto" até a chamada, onde
// TypeScript o preenche a partir do argumento recebido.
export function exemplos(): void {
  function primeiro<T>(itens: T[]): T | undefined {
    return itens.length === 0 ? undefined : itens[0];
  }

  // T é inferido a cada chamada — aqui vira number...
  const n = primeiro([10, 20, 30]);
  console.log(n); // 10, com tipo number | undefined

  // ...e aqui vira string, sem precisar de duas funções nem de any.
  const s = primeiro(["a", "b"]);
  console.log(s); // "a", com tipo string | undefined

  // Com any no lugar de T, a função ainda funcionaria em tempo de execução,
  // mas o tipo do retorno seria any — perderíamos a garantia de que
  // primeiro(numeros) devolve um number, especificamente.
  function primeiroComAny(itens: any[]): any {
    return itens[0];
  }
  const perdido = primeiroComAny([10, 20]); // tipo any, não number
  console.log(perdido);

  // T também pode aparecer em mais de um lugar da assinatura, amarrando os
  // tipos entre si: aqui, o valor default tem que ser do mesmo T do array.
  function primeiroOuDefault<T>(itens: T[], valorDefault: T): T {
    return itens.length === 0 ? valorDefault : itens[0];
  }
  console.log(primeiroOuDefault([1, 2], 0)); // 1
  console.log(primeiroOuDefault([], 0)); // 0
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Devolva o último elemento de itens, ou undefined se vier vazio.
export function ultimo<T>(itens: T[]): T | undefined {
  return itens.length === 0 ? undefined : itens[itens.length - 1];
}
