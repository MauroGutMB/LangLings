// Devolva um objeto NOVO com o nome de pessoa em maiúsculas, sem alterar o
// objeto recebido.
export function comNomeMaiusculo(pessoa) {
  return { ...pessoa, nome: pessoa.nome.toUpperCase() };
}
