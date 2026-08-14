// TypeScript infere o tipo a partir do valor inicial — não é preciso anotar
// quando o valor já deixa claro o que é.
export function exemplos(): void {
  const total = 42; // inferido como number, sem anotação
  let contador = 0; // também number, mas contador pode mudar de VALOR

  contador = contador + 1; // ok: ainda é number
  // contador = "zero"; // erro de compilação: contador é number, não string

  console.log(total, contador);

  // A diferença importante: uma const guarda um VALOR que não muda, então
  // TypeScript pode inferir o tipo mais estreito possível — o literal exato.
  const status = "ativo"; // tipo inferido: "ativo" (o literal, não string)

  // Uma let pode receber outros valores depois, então TypeScript infere o
  // tipo mais amplo que caiba qualquer atribuição futura compatível.
  let etapa = "inicio"; // tipo inferido: string (não o literal "inicio")
  etapa = "meio"; // ok: qualquer string serve para uma variável do tipo string

  console.log(status, etapa);

  // Anotar um tipo que já seria inferido é redundante, não errado — mas o
  // idioma comum em TypeScript é deixar a inferência trabalhar quando ela já
  // acerta, e anotar só onde ela não tem como adivinhar (como parâmetros).
  const jaObvio: number = 10; // anotação supérflua: 10 já seria number sozinho
  console.log(jaObvio);
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Aplique um desconto percentual a um preço.
export function precoComDesconto(preco: number, percentual: number): number {
  return 0; // <- troque isto
}
