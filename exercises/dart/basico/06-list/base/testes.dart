import 'dart:io';
import 'exercicio.dart';

int _falhas = 0;

bool _iguais(Object? a, Object? b) {
  if (a is List && b is List) {
    if (a.length != b.length) return false;
    for (var i = 0; i < a.length; i++) {
      if (!_iguais(a[i], b[i])) return false;
    }
    return true;
  }
  return a == b;
}

void verificar(String oQue, Object? esperado, Object? obtido) {
  if (_iguais(esperado, obtido)) {
    print('ok    $oQue');
    return;
  }
  print('FALHA $oQue\n      esperado: $esperado\n      obtido:   $obtido');
  _falhas++;
}

void main() {
  verificar('apenasPositivos([-2,-1,0,1,2])', [1, 2],
      apenasPositivos([-2, -1, 0, 1, 2]));
  verificar('apenasPositivos([3,1,2])', [3, 1, 2], apenasPositivos([3, 1, 2]));
  verificar('apenasPositivos([])', [], apenasPositivos([]));
  verificar('apenasPositivos([-5,-6])', [], apenasPositivos([-5, -6]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
