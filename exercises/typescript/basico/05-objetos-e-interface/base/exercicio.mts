// Uma interface descreve o FORMATO de um objeto: quais campos existem e o
// tipo de cada um. Não é uma classe — não tem implementação, só a forma.
interface Produto {
  nome: string;
  preco: number;
}

export function exemplos(): void {
  const produto: Produto = { nome: "Caneta", preco: 2.5 };
  console.log(produto.nome, produto.preco);

  // Um objeto que não bate com a interface é erro de compilação: falta
  // campo, campo a mais, ou tipo errado num campo.
  // const invalido: Produto = { nome: "Lápis" }; // erro: falta 'preco'

  function precoComImposto(p: Produto): number {
    return p.preco * 1.1;
  }
  console.log(precoComImposto(produto));

  // Um campo opcional (?) pode faltar no objeto — e dentro da função o tipo
  // desse campo inclui undefined, igual a um parâmetro opcional de função.
  interface ProdutoComDesconto extends Produto {
    percentualDesconto?: number;
  }

  function precoFinal(p: ProdutoComDesconto): number {
    if (p.percentualDesconto === undefined) {
      return p.preco;
    }
    return p.preco - (p.preco * p.percentualDesconto) / 100;
  }
  console.log(precoFinal({ nome: "Caderno", preco: 10 })); // 10
  console.log(precoFinal({ nome: "Caderno", preco: 10, percentualDesconto: 20 })); // 8
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

interface Pessoa {
  nome: string;
  idade?: number;
}

// SUA VEZ
//
// Devolva "nome (idade anos)" quando idade foi informada, ou só "nome"
// quando não foi.
export function apresentar(pessoa: Pessoa): string {
  return ""; // <- troque isto
}
