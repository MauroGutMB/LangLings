// Devolva um objeto NOVO com o nome de pessoa em maiúsculas, sem alterar o
// objeto recebido.
//
// TODO: o retorno está certo, mas pessoa sai modificado também.
export function comNomeMaiusculo(pessoa) {
  pessoa.nome = pessoa.nome.toUpperCase();
  return pessoa;
}
