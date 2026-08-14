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
  verificar('tamanhoOuZero("abc")', 3, tamanhoOuZero('abc'));
  verificar('tamanhoOuZero("")', 0, tamanhoOuZero(''));

  // A asserção que separa a versão com ! da correta: null nunca chega a ter
  // .length lido, então a versão ingênua explode aqui em vez de devolver 0.
  verificar('tamanhoOuZero(null)', 0, tamanhoOuZero(null));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
