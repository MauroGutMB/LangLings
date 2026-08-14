// any desliga a checagem de tipos: qualquer operação em cima de um any
// compila, mesmo que exploda em tempo de execução.
export function exemplos(): void {
  function dobrarComAny(valor: any): number {
    return valor * 2; // compila mesmo se valor não for number
  }
  console.log(dobrarComAny(21)); // 42
  console.log(dobrarComAny("oi")); // NaN — compilou, mas quebrou em silêncio

  // unknown também aceita qualquer valor na entrada, mas TypeScript não
  // deixa usá-lo como se já soubesse o tipo — é preciso checar primeiro.
  function dobrarComUnknown(valor: unknown): number {
    // return valor * 2; // erro de compilação: valor é unknown

    if (typeof valor === "number") {
      // Dentro deste bloco, TypeScript já sabe que valor é number —
      // essa checagem se chama "narrowing" (estreitamento de tipo).
      return valor * 2;
    }
    return 0;
  }
  console.log(dobrarComUnknown(21)); // 42
  console.log(dobrarComUnknown("oi")); // 0, sem quebrar

  // A regra prática: prefira unknown a any na entrada de uma função — ele
  // força quem escreve a função a considerar todos os tipos possíveis antes
  // de agir, em vez de confiar cegamente no chamador.
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Devolva valor se ele já for number; caso contrário devolva 0.
export function paraNumeroSeguro(valor: unknown): number {
  return 0; // <- troque isto
}
