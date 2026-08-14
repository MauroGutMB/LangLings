// Strings em JavaScript são imutáveis: todo método que "muda" uma string na
// verdade devolve outra. Se você não guardar o retorno, nada aconteceu.
function exemplos() {
  const nome = 'Ana Maria Silva';

  // Template literal: crase no lugar das aspas, e ${} interpola qualquer
  // expressão. Substitui a soma de strings com + em quase todo caso.
  const idade = 30;
  console.log(`${nome} tem ${idade} anos`); // Ana Maria Silva tem 30 anos
  console.log(`ano que vem: ${idade + 1}`); // ano que vem: 31

  // Tamanho e acesso por posição. A primeira letra é a posição 0.
  console.log(nome.length, nome[0]); // 15 A

  // Busca.
  console.log(nome.includes('Maria')); // true
  console.log(nome.indexOf('Maria')); // 4
  console.log(nome.indexOf('Joana')); // -1  <- não achou

  // Recorte: slice(inicio, fim), com o fim de fora.
  console.log(nome.slice(0, 3)); // Ana
  console.log(nome.slice(-5)); // Silva  <- negativo conta do fim

  // Transformações devolvem uma string nova. A original continua intacta.
  console.log(nome.toUpperCase()); // ANA MARIA SILVA
  console.log('  espaço  '.trim()); // espaço
  console.log(nome.replaceAll('a', '@')); // An@ M@ri@ Silv@
  console.log(nome); // Ana Maria Silva  <- nada acima a alterou

  // split quebra a string num array; join cola o array de volta numa string.
  const partes = nome.split(' ');
  console.log(partes); // [ 'Ana', 'Maria', 'Silva' ]
  console.log(partes.join('-')); // Ana-Maria-Silva
  console.log(partes.map((p) => p.length).join(',')); // 3,5,5
}

// iniciais devolve a primeira letra de cada palavra do nome, em maiúsculas.
//
// O trim vem antes do split porque um espaço na ponta viraria uma palavra
// vazia, e a inicial dela seria undefined. O toUpperCase fica no fim, sobre o
// resultado inteiro: uma chamada em vez de uma por palavra.
export function iniciais(nomeCompleto) {
  return nomeCompleto
    .trim()
    .split(' ')
    .map((parte) => parte[0])
    .join('')
    .toUpperCase();
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
