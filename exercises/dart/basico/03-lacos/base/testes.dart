import 'dart:io';
import 'exercicio.dart';

int _falhas = 0;

void verificar(String oQue, Object? esperado, Object? obtido) {
  if (esperado == obtido) {
    print('ok    $oQue');
    return;
  }
  print('FALHA $oQue\n      esperado: $esperado\n      obtido:   $obtido');
  _falhas++;
}

void main() {
  verificar('somaPares([1,2,3,4,5,6])', 12, somaPares([1, 2, 3, 4, 5, 6]));
  verificar('somaPares([1,3,5])', 0, somaPares([1, 3, 5]));
  verificar('somaPares([])', 0, somaPares([]));
  verificar('somaPares com negativos', 2, somaPares([-2, 1, 3, 4, -3]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
