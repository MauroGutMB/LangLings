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
  // A asserção que separa `as int` de ~/: mesmo numa divisão exata, / devolve
  // double, e `as int` num double explode em vez de converter.
  verificar('mediaInteira([4,4])', 4, mediaInteira([4, 4]));

  verificar('mediaInteira([1,2,4])', 2, mediaInteira([1, 2, 4]));
  verificar('mediaInteira([7])', 7, mediaInteira([7]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
