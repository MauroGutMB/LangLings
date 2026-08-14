// Junta cliente e quantidade num texto: '$cliente: $quantidade unidade(s)'.
String formatarPedido({required String cliente, int quantidade = 1}) {
  return '$cliente: $quantidade unidade(s)';
}

// Versão de formatarPedido para quem não tem cadastro, com nome padrão
// 'convidado'.
String pedidoConvidado(int quantidade) =>
    formatarPedido(cliente: 'convidado', quantidade: quantidade);
