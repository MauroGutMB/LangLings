import 'dart:io';
import 'exercicio.dart';

int _falhas = 0;

bool _iguaisConjunto(Set<int> a, List<int> b) {
  return a.length == b.length && b.every(a.contains);
}

void verificar(String oQue, List<int> esperado, Set<int> obtido) {
  if (_iguaisConjunto(obtido, esperado)) {
    print('ok    $oQue');
    return;
  }
  print('FALHA $oQue\n      esperado: $esperado\n      obtido:   $obtido');
  _falhas++;
}

void main() {
  verificar('elementosComuns([1,2,3],[2,3,4])', [2, 3],
      elementosComuns([1, 2, 3], [2, 3, 4]));
  verificar('elementosComuns sem interseção', [],
      elementosComuns([1, 2], [3, 4]));
  verificar('elementosComuns com repetições', [5],
      elementosComuns([5, 5, 5], [5, 6]));
  verificar('elementosComuns com lista vazia', [],
      elementosComuns([], [1, 2]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
