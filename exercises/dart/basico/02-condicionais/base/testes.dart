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
  verificar('classificarNota(10)', 'aprovado', classificarNota(10));
  verificar('classificarNota(7)', 'aprovado', classificarNota(7));
  verificar('classificarNota(6.9)', 'recuperação', classificarNota(6.9));
  verificar('classificarNota(5)', 'recuperação', classificarNota(5));
  verificar('classificarNota(4.9)', 'reprovado', classificarNota(4.9));
  verificar('classificarNota(0)', 'reprovado', classificarNota(0));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
