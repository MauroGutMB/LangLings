// Um objeto é uma coleção de pares chave -> valor, com chaves em string. É a
// estrutura que carrega dados com nome em JavaScript.
function exemplos() {
  const usuario = {
    nome: 'Ana',
    idade: 30,
    endereco: { cidade: 'Recife' }, // objetos aninham
  };

  // Acesso por ponto quando a chave é conhecida e literal; por colchetes
  // quando ela vem de uma variável ou tem caracteres fora do comum.
  const campo = 'idade';
  console.log(usuario.nome, usuario[campo]); // Ana 30
  console.log(usuario.endereco.cidade); // Recife

  // Ler uma chave que não existe devolve undefined, sem erro.
  console.log(usuario.email); // undefined
  console.log('email' in usuario); // false  <- a pergunta explícita

  // Escrever cria a chave se ela não existir.
  usuario.email = 'ana@exemplo.com';
  console.log(Object.keys(usuario)); // [ 'nome', 'idade', 'endereco', 'email' ]

  // Desestruturação arranca campos para variáveis de mesmo nome. O = ali
  // dentro é o valor padrão para quando a chave não existe.
  const { nome, idade, apelido = 'sem apelido' } = usuario;
  console.log(nome, idade, apelido); // Ana 30 sem apelido

  // As três formas de percorrer um objeto.
  console.log(Object.keys({ a: 1, b: 2 })); // [ 'a', 'b' ]
  console.log(Object.values({ a: 1, b: 2 })); // [ 1, 2 ]
  console.log(Object.entries({ a: 1 })); // [ [ 'a', 1 ] ]

  // Spread copia as chaves de um objeto para dentro de outro. Quando a mesma
  // chave aparece duas vezes, a ÚLTIMA vence — é isso que faz o padrão ficar
  // antes e o valor específico depois.
  const padrao = { cor: 'preto', tamanho: 'M' };
  const escolha = { tamanho: 'G' };
  console.log({ ...padrao, ...escolha }); // { cor: 'preto', tamanho: 'G' }
  console.log({ ...escolha, ...padrao }); // { cor: 'preto', tamanho: 'M' }  <- invertido
  console.log(padrao); // { cor: 'preto', tamanho: 'M' }  <- o original nunca muda
}

// comPadroes preenche o que a configuração não disse.
//
// O literal novo é o que garante que config saia intacta: espalhar é copiar
// chave a chave, e a ordem escolhe o vencedor de cada repetida — padrões
// primeiro, configuração depois.
export function comPadroes(config) {
  return { tema: 'claro', fonte: 14, animacoes: true, ...config };
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
