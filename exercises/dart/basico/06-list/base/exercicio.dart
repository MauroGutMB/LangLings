// List é uma coleção ordenada, indexada, que aceita repetição.
void exemplos() {
  final numeros = <int>[10, 20, 30];
  numeros.add(40);
  print(numeros); // [10, 20, 30, 40]
  print(numeros[0]); // 10 — acesso por índice, começando em 0

  numeros.remove(20); // remove o VALOR 20, não o índice 20
  print(numeros); // [10, 30, 40]

  numeros.sort((a, b) => b.compareTo(a)); // decrescente
  print(numeros); // [40, 30, 10]

  // spread (...) explode uma lista dentro de outro literal.
  final combinada = [0, ...numeros, 100];
  print(combinada); // [0, 40, 30, 10, 100]

  print(numeros.length); // 3
  print(numeros.contains(30)); // true
}

// SUA VEZ
//
// Devolva uma nova lista só com os valores de ns maiores que zero, mantendo
// a ordem original.
List<int> apenasPositivos(List<int> ns) {
  return []; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
