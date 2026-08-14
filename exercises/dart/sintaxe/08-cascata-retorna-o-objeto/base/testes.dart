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
  // A asserção que separa a cascata da versão correta: incrementar() devolve
  // um Contador novo a cada vez, e a cascata descarta esses três retornos.
  verificar('incrementarTres(5)', 8, incrementarTres(5));

  verificar('incrementarTres(0)', 3, incrementarTres(0));
  verificar('incrementarTres(-2)', 1, incrementarTres(-2));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
