class Pessoa {
  final String nome;
  Pessoa(this.nome);

  String saudacao() => 'Olá, $nome!';
}

// Devolve a saudação de pessoa.
String cumprimentar(Pessoa pessoa) {
  return pessoa.saudacao();
}
