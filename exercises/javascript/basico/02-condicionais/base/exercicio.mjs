// Toda decisão em JavaScript passa por um valor booleano — ou por um valor que
// a linguagem converte em booleano quando pergunta "isso é verdadeiro?".
function exemplos() {
  const temperatura = 31;

  // A cadeia if / else if / else para em cima do PRIMEIRO teste verdadeiro.
  // Por isso a ordem importa: aqui, do maior limite para o menor.
  if (temperatura >= 30) {
    console.log('calor'); // calor
  } else if (temperatura >= 20) {
    console.log('agradável');
  } else {
    console.log('frio');
  }

  // O ternário é uma EXPRESSÃO: ele devolve um valor, então pode ser
  // atribuído. Use quando o if inteiro só serviria para escolher um valor.
  const roupa = temperatura >= 30 ? 'camiseta' : 'casaco';
  console.log(roupa); // camiseta

  // Comparações devolvem booleano; && e || combinam booleanos.
  console.log(temperatura > 25 && temperatura < 40); // true
  console.log(temperatura < 0 || temperatura > 35); // false
  console.log(!(temperatura === 31)); // false

  // switch compara com === contra cada case. Sem o break, a execução
  // ESCORREGA para o case seguinte — o que às vezes é o que se quer.
  const dia = 'sabado';
  switch (dia) {
    case 'sabado':
    case 'domingo':
      console.log('fim de semana'); // fim de semana
      break;
    default:
      console.log('dia útil');
  }
}

// SUA VEZ
//
// Devolva a letra da nota: 'A' a partir de 90, 'B' a partir de 80, 'C' a
// partir de 70, e 'F' abaixo disso.
export function classificar(nota) {
  return ''; // <- troque isto
}

// Só executa os exemplos quando você roda o arquivo direto (node exercicio.mjs).
// Sob `node --test` o arquivo é importado, e essa saída se misturaria à dos testes.
if (import.meta.main) exemplos();
