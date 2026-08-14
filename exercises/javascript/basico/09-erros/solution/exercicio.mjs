// Erro em JavaScript é uma interrupção: throw abandona a função na hora, e o
// controle sobe até o try/catch mais próximo. Quem não trata, propaga.

// Um Error próprio é uma classe que estende Error. O valor dele está no
// instanceof: quem captura consegue separar "entrada inválida" de qualquer
// outra falha, sem comparar mensagens.
export class ErroDeEntrada extends Error {
  constructor(mensagem) {
    super(mensagem);
    this.name = 'ErroDeEntrada';
  }
}

function exemplos() {
  function idadeValida(idade) {
    if (typeof idade !== 'number') {
      throw new ErroDeEntrada('idade precisa ser número');
    }
    if (idade < 0) {
      throw new ErroDeEntrada(`idade negativa: ${idade}`);
    }
    return idade;
  }

  // O caminho feliz não passa por nada disso.
  console.log(idadeValida(30)); // 30

  // try/catch: o catch recebe o valor lançado. Nada depois do throw, dentro
  // do try, chega a rodar.
  try {
    idadeValida(-1);
    console.log('nunca chega aqui');
  } catch (erro) {
    console.log(erro.name, '|', erro.message); // ErroDeEntrada | idade negativa: -1
    console.log(erro instanceof ErroDeEntrada, erro instanceof Error); // true true
  }

  // finally roda tendo dado erro ou não — é onde vai a limpeza.
  try {
    idadeValida('trinta');
  } catch (erro) {
    console.log('capturado:', erro.message); // capturado: idade precisa ser número
  } finally {
    console.log('sempre roda'); // sempre roda
  }

  // instanceof é o que permite tratar só o que você sabe tratar e deixar o
  // resto subir.
  try {
    JSON.parse('{isso não é json}');
  } catch (erro) {
    console.log(erro instanceof ErroDeEntrada); // false  <- é um SyntaxError
    console.log(erro instanceof Error); // true
  }

  // Nem toda operação inválida lança: dividir por zero devolve Infinity, e o
  // programa segue com um número que não é número nenhum.
  console.log(1 / 0); // Infinity
}

// dividir devolve a / b, recusando o divisor zero.
//
// A recusa é explícita porque a linguagem não recusa: a / 0 devolve Infinity,
// e o valor inválido só apareceria muitas linhas depois, longe da causa.
export function dividir(a, b) {
  if (b === 0) {
    throw new ErroDeEntrada('divisão por zero');
  }
  return a / b;
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
