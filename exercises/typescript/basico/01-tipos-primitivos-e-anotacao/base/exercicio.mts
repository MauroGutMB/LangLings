// Os três tipos primitivos mais comuns e a forma de anotá-los explicitamente:
// nome: Tipo. A anotação não é obrigatória quando o valor inicial já deixa o
// tipo óbvio (mais sobre isso no próximo exercício), mas em parâmetros de
// função ela É obrigatória — TypeScript não infere o tipo de quem chama.
export function exemplos(): void {
  const idade: number = 30;
  const nome: string = "Ana";
  const ativo: boolean = true;

  console.log(idade, nome, ativo);

  // Parâmetros e retorno de função são anotados do mesmo jeito.
  function dobrar(valor: number): number {
    return valor * 2;
  }
  console.log(dobrar(21)); // 42

  // Misturar tipos onde não cabe é erro de COMPILAÇÃO, não de execução —
  // é isso que torna TypeScript útil: o erro aparece antes de rodar.
  // const errado: number = "trinta"; // descomente para ver tsc reclamar

  // Template strings interpolam qualquer valor, mas o resultado é sempre string.
  const mensagem: string = `${nome} tem ${idade} anos`;
  console.log(mensagem);
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Converta celsius para Fahrenheit: (celsius * 9 / 5) + 32.
export function celsiusParaFahrenheit(celsius: number): number {
  return 0; // <- troque isto
}
