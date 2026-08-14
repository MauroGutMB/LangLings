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
  // A asserção que separa a versão que muta _base da correta: um `const []`
  // rejeita .add em tempo de execução — isto nem chega a devolver um valor.
  verificar('primeira chamada', ['a', 'b', 'c'], comItem('c'));

  // Mesmo se alguém trocar const por var, sem copiar a lista continua
  // acumulando entre chamadas — esta pega esse caso também.
  verificar('segunda chamada não acumula', ['a', 'b', 'd'], comItem('d'));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
