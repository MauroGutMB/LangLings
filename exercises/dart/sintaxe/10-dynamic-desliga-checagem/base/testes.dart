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
  // A asserção que separa dynamic mal usado da versão correta: o typo no
  // nome do método só é descoberto ao chamar, não ao compilar.
  verificar('cumprimentar(Pessoa("Ana"))', 'Olá, Ana!', cumprimentar(Pessoa('Ana')));

  verificar('cumprimentar(Pessoa("Bruno"))', 'Olá, Bruno!',
      cumprimentar(Pessoa('Bruno')));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
