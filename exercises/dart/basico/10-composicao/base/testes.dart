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
  verificar(
      'totalCompra com dois itens',
      23,
      totalCompra([
        {'preco': 5, 'quantidade': 3},
        {'preco': 4, 'quantidade': 2},
      ]));
  verificar('totalCompra com lista vazia', 0, totalCompra([]));
  verificar(
      'totalCompra com preco fracionado',
      7.5,
      totalCompra([
        {'preco': 2.5, 'quantidade': 3},
      ]));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
