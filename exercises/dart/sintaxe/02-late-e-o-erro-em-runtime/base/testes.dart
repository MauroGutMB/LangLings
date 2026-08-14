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
  verificar('formatarConfig(modo: "debug")', 'MODO DEBUG',
      formatarConfig(modo: 'debug'));
  verificar('formatarConfig(modo: "prod")', 'MODO PRODUCAO',
      formatarConfig(modo: 'prod'));

  // A asserção que separa o late descoberto da versão correta: sem modo,
  // nenhum if atribui `resultado`, e lê-la explode em vez de cair no padrão.
  verificar('formatarConfig() sem modo', 'MODO PADRAO', formatarConfig());

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
