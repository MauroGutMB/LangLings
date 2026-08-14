class Pessoa {
  final String nome;
  Pessoa(this.nome);

  String saudacao() => 'Olá, $nome!';
}

// TODO: devolve a saudação de pessoa.
String cumprimentar(dynamic pessoa) {
  return pessoa.saudaco();
}
