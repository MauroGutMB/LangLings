import 'dart:io';
import 'exercicio.dart';

int _falhas = 0;

bool _iguais(Object? a, Object? b) {
  if (a is Map && b is Map) {
    if (a.length != b.length) return false;
    for (final k in a.keys) {
      if (!b.containsKey(k) || !_iguais(a[k], b[k])) return false;
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
  verificar('contarPalavras agrupa repetições', {'go': 3, 'rust': 1},
      contarPalavras(['go', 'rust', 'go', 'go']));
  verificar('contarPalavras distingue maiúsculas', {'Go': 1, 'go': 1},
      contarPalavras(['Go', 'go']));
  verificar('contarPalavras com lista vazia', {}, contarPalavras([]));

  final got = contarPalavras(['a']);
  verificar('contarPalavras não inventa chaves', false, got.containsKey('b'));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
