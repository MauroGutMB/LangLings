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
  verificar('Retangulo(3,4).area()', 12.0, Retangulo(3, 4).area());
  verificar('Retangulo(2.5,2).area()', 5.0, Retangulo(2.5, 2).area());
  verificar('Retangulo(0,5).area()', 0.0, Retangulo(0, 5).area());

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
