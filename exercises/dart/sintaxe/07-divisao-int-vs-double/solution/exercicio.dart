// Devolve a média inteira (arredondada para baixo) de ns.
int mediaInteira(List<int> ns) {
  final soma = ns.reduce((a, b) => a + b);
  return soma ~/ ns.length;
}
