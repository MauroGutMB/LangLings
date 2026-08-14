interface Pessoa {
  nome: string;
  idade: number;
}

// TODO: pessoa está tipado como any, então o compilador não confere nenhum
// acesso a campo aqui dentro — inclusive um typo no nome do campo passaria
// batido, compilando normalmente e só quebrando (em silêncio) em produção.
export function apresentarPessoa(pessoa: any): string {
  return `Olá, ${pessoa.nomee}, você tem ${pessoa.idade} anos`;
}
