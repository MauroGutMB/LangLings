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
  verificar('formatarPedido com cliente', 'Ana: 3 unidade(s)',
      formatarPedido(cliente: 'Ana', quantidade: 3));

  // A asserção que separa a versão opcional da correta: pedidoConvidado
  // esqueceu de passar cliente — exatamente o que `required` capturaria em
  // tempo de compilação, em vez de estourar em tempo de execução.
  verificar(
      'pedidoConvidado tem nome padrão', 'convidado: 2 unidade(s)', pedidoConvidado(2));

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
