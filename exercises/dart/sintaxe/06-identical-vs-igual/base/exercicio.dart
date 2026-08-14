class Moeda {
  final String codigo;
  Moeda(this.codigo);

  @override
  bool operator ==(Object other) => other is Moeda && codigo == other.codigo;

  @override
  int get hashCode => codigo.hashCode;
}

// TODO: devolve true quando a e b têm o mesmo código — mesmo sendo
// instâncias diferentes.
bool mesmoValor(Moeda a, Moeda b) {
  return identical(a, b);
}
