// Três formas de repetir: for clássico (quando o índice importa), for-in
// (quando só o valor importa) e while (quando a condição de parada não é
// "andei N vezes").
void exemplos() {
  for (var i = 0; i < 3; i++) {
    print('índice $i'); // índice 0, índice 1, índice 2
  }

  final frutas = ['maçã', 'pera', 'uva'];
  for (final fruta in frutas) {
    print(fruta); // maçã, pera, uva — sem precisar de índice
  }

  var restante = 3;
  while (restante > 0) {
    print('faltam $restante'); // faltam 3, faltam 2, faltam 1
    restante--;
  }

  // break sai do laço; continue pula para a próxima iteração.
  for (final n in [1, 2, 3, 4, 5]) {
    if (n == 4) break; // para antes de imprimir 4 e 5
    if (n.isOdd) continue; // pula os ímpares
    print(n); // 2
  }
}

// SUA VEZ
//
// Devolva a soma dos números pares de ns. Lista vazia devolve 0.
int somaPares(List<int> ns) {
  return 0; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
