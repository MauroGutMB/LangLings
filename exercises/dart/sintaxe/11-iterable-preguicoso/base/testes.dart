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
  // A asserção que separa o map() nunca consumido da versão correta: sem
  // forçar a avaliação, o corpo do map nunca roda e contador fica em 0.
  verificar('contarProcessados([1,2,3])', 3, contarProcessados([1, 2, 3]));

  verificar('contarProcessados com 5 elementos', 5,
      contarProcessados([10, 20, 30, 40, 50]));
  verificar('contarProcessados([])', 0, contarProcessados([]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
