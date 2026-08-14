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
  final a = Moeda('BRL');
  final b = Moeda('BRL');

  // A asserção que separa identical de ==: a e b têm o mesmo valor, mas são
  // duas instâncias diferentes — identical(a, b) é false.
  verificar('mesmo código, instâncias diferentes', true, mesmoValor(a, b));

  verificar('códigos diferentes', false, mesmoValor(Moeda('BRL'), Moeda('USD')));
  verificar('mesma instância', true, mesmoValor(a, a));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
