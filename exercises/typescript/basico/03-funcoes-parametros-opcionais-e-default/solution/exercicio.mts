// Um parâmetro opcional (?) pode ser omitido na chamada — e dentro da função
// seu tipo inclui undefined automaticamente.
export function exemplos(): void {
  function saudacao(nome: string, titulo?: string): string {
    // titulo tem tipo `string | undefined` aqui dentro, por causa do ?.
    if (titulo === undefined) {
      return `Olá, ${nome}`;
    }
    return `Olá, ${titulo} ${nome}`;
  }
  console.log(saudacao("Ana")); // Olá, Ana
  console.log(saudacao("Ana", "Dra.")); // Olá, Dra. Ana

  // Um parâmetro com valor default é diferente: quando omitido, ele recebe o
  // default automaticamente, e dentro da função o tipo NÃO inclui undefined.
  function repetir(texto: string, vezes: number = 2): string {
    return texto.repeat(vezes); // vezes é sempre number aqui, nunca undefined
  }
  console.log(repetir("ha")); // haha
  console.log(repetir("ha", 3)); // hahaha

  // Parâmetros com default podem vir depois de obrigatórios, mas nunca antes
  // de um parâmetro obrigatório sem default — a ordem importa na assinatura.
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Devolva "nome sobrenome" quando sobrenome foi informado, ou só "nome"
// quando não foi.
export function formatarNome(nome: string, sobrenome?: string): string {
  if (sobrenome === undefined) {
    return nome;
  }
  return `${nome} ${sobrenome}`;
}
