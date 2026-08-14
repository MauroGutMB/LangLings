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
  verificar('areaDoQuadrado(3)', 9.0, areaDoQuadrado(3));
  verificar('areaDoQuadrado(2.5)', 6.25, areaDoQuadrado(2.5));
  verificar('areaDoQuadrado(0)', 0.0, areaDoQuadrado(0));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
