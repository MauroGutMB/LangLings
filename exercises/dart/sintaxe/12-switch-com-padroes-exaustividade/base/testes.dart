import 'dart:io';
import 'dart:math' as math;
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
  verificar('area(Circulo(3))', math.pi * 3 * 3, area(Circulo(3)));
  verificar('area(Quadrado(4))', 16.0, area(Quadrado(4)));

  // A asserção que separa o caso curinga da versão exaustiva: Retangulo cai
  // no `_ => 0` em vez de ter seu próprio cálculo.
  verificar('area(Retangulo(2,3))', 6.0, area(Retangulo(2, 3)));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
