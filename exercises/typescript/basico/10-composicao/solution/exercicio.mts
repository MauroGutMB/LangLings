interface Produto {
  nome: string;
  preco: number;
}

// Um Cupom é uma union de dois formatos de objeto: o campo tipo funciona
// como "discriminador" — ele é o que diferencia qual dos dois formatos você
// tem nas mãos, e cada formato pode ter campos próprios.
type Cupom =
  | { tipo: "percentual"; valor: number }
  | { tipo: "fixo"; valor: number };

export function exemplos(): void {
  const produtos: Produto[] = [
    { nome: "Caneta", preco: 2 },
    { nome: "Caderno", preco: 10 },
  ];

  const total = produtos.reduce((soma, p) => soma + p.preco, 0);
  console.log(total); // 12

  function aplicarCupom(total: number, cupom?: Cupom): number {
    if (cupom === undefined) {
      return total;
    }
    // Checar cupom.tipo estreita qual dos dois formatos da union estamos
    // olhando — depois do if, TypeScript sabe que é o ramo "percentual".
    if (cupom.tipo === "percentual") {
      return total - (total * cupom.valor) / 100;
    }
    return Math.max(0, total - cupom.valor);
  }

  console.log(aplicarCupom(total)); // 12, sem cupom
  console.log(aplicarCupom(total, { tipo: "percentual", valor: 50 })); // 6
  console.log(aplicarCupom(total, { tipo: "fixo", valor: 5 })); // 7

  // Optional chaining ainda cabe aqui: se cupom fosse dentro de um Pedido
  // opcional, pedido?.cupom?.valor encadearia naturalmente.
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Some o preco de cada produto e aplique o cupom, se houver: "percentual"
// reduz o total em cupom.valor por cento; "fixo" subtrai cupom.valor. O
// resultado nunca é negativo.
export function calcularTotal(produtos: Produto[], cupom?: Cupom): number {
  const total = produtos.reduce((soma, p) => soma + p.preco, 0);
  if (cupom === undefined) {
    return total;
  }
  if (cupom.tipo === "percentual") {
    return Math.max(0, total - (total * cupom.valor) / 100);
  }
  return Math.max(0, total - cupom.valor);
}
