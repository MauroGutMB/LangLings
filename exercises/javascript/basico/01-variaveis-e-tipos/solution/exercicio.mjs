// Em JavaScript o tipo pertence ao valor, não à variável: a mesma variável
// pode guardar um número agora e uma string depois. O que a declaração decide
// é outra coisa — se o vínculo com o valor pode ser refeito.
function exemplos() {
  let contador = 0; // let: pode receber outro valor depois
  const nome = 'Ana'; // const: reatribuir nome seria erro em tempo de execução
  contador = contador + 1;
  console.log(nome, contador); // Ana 1

  // const congela o VÍNCULO, não o conteúdo. O array continua mutável; o que
  // não dá é apontar `cores` para outro array.
  const cores = ['azul'];
  cores.push('verde');
  console.log(cores); // [ 'azul', 'verde' ]

  // typeof responde o tipo do valor, sempre como string.
  console.log(typeof 42, typeof 'oi', typeof true); // number string boolean

  // Declarar sem atribuir dá undefined: o "ainda não tem valor" da linguagem.
  // null é diferente — é o vazio que alguém escreveu de propósito.
  let semValor;
  const vazio = null;
  console.log(semValor, vazio); // undefined null
  console.log(semValor === undefined); // true
  console.log(vazio === null); // true
  console.log(vazio === undefined); // false  <- são dois valores distintos

  // Ler uma propriedade que não existe não é erro: devolve undefined.
  const config = { tema: 'escuro' };
  console.log(config.fonte); // undefined

  // 0, '' e false são valores como quaisquer outros. Nenhum deles é ausência.
  console.log(0 === undefined, '' === null, false === undefined); // false false false
}

// temValor separa "não tem valor" de "tem um valor que parece pouco".
//
// A comparação é com === contra as duas formas de ausência, e não com um
// if (valor): converter para booleano trataria 0, '' e false como ausência —
// que é justamente o bug que essa distinção existe para evitar.
export function temValor(valor) {
  return valor !== undefined && valor !== null;
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
