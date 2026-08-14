// Nada de novo daqui para baixo: objeto guardando dados com nome, array
// guardando a lista deles, Map agrupando por chave e erro recusando entrada
// inválida — tudo junto, que é como o código de verdade aparece.

export class ErroDeEstoque extends Error {
  constructor(mensagem) {
    super(mensagem);
    this.name = 'ErroDeEstoque';
  }
}

// O preço está em centavos porque dinheiro em ponto flutuante acumula erro de
// arredondamento: 0.1 + 0.2 não dá 0.3. Inteiro não tem esse problema.
const estoque = [
  { nome: 'caneta', categoria: 'papelaria', preco: 250, qtd: 4 },
  { nome: 'caderno', categoria: 'papelaria', preco: 1200, qtd: 2 },
  { nome: 'café', categoria: 'cozinha', preco: 1850, qtd: 3 },
];

function exemplos() {
  // Desestruturação direto no parâmetro: a função declara quais campos usa.
  const valor = ({ preco, qtd }) => preco * qtd;
  console.log(valor(estoque[0])); // 1000

  // reduce sobre objetos: o total geral do estoque.
  const total = estoque.reduce((acumulado, item) => acumulado + valor(item), 0);
  console.log(total); // 9000

  // filter + map + join: um relatório em texto dos itens caros.
  const caros = estoque
    .filter((item) => valor(item) > 2000)
    .map(({ nome, qtd }) => `${nome} x${qtd}`)
    .join(', ');
  console.log(caros); // café x3

  // Map como índice: agrupa os itens por categoria. O `?? []` cobre a
  // primeira vez que a categoria aparece.
  const porCategoria = new Map();
  for (const item of estoque) {
    porCategoria.set(item.categoria, [...(porCategoria.get(item.categoria) ?? []), item]);
  }
  console.log(porCategoria.get('papelaria').length); // 2
  console.log([...porCategoria.keys()]); // [ 'papelaria', 'cozinha' ]

  // Erro como recusa: a validação lança, quem chama decide o que fazer.
  try {
    validar([]);
  } catch (erro) {
    console.log(erro.name, '|', erro.message); // ErroDeEstoque | nenhum item
  }
}

function validar(itens) {
  if (itens.length === 0) {
    throw new ErroDeEstoque('nenhum item');
  }
  return itens;
}

// resumir devolve o valor total de estoque por categoria.
//
// Os dois erros saem cedo, cada um no ponto em que dá para detectá-lo: o array
// vazio antes do laço, a categoria vazia dentro dele. O `?? 0` cobre a
// primeira aparição de cada categoria, onde o get ainda devolve undefined.
export function resumir(itens) {
  if (itens.length === 0) {
    throw new ErroDeEstoque('nenhum item');
  }

  const total = new Map();

  for (const { nome, categoria, preco, qtd } of itens) {
    if (!categoria) {
      throw new ErroDeEstoque(`item sem categoria: ${nome}`);
    }
    total.set(categoria, (total.get(categoria) ?? 0) + preco * qtd);
  }

  return total;
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
