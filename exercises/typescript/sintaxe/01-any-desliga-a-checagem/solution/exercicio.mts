interface Pessoa {
  nome: string;
  idade: number;
}

export function apresentarPessoa(pessoa: Pessoa): string {
  return `Olá, ${pessoa.nome}, você tem ${pessoa.idade} anos`;
}
