// TODO: junta cliente e quantidade num texto: '$cliente: $quantidade unidade(s)'.
String formatarPedido({String? cliente, int quantidade = 1}) {
  return '${cliente!}: $quantidade unidade(s)';
}

// TODO: versão de formatarPedido para quem não tem cadastro, com nome
// padrão 'convidado'.
String pedidoConvidado(int quantidade) => formatarPedido(quantidade: quantidade);
