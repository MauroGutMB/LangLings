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
  final pontos = [Ponto(1, 1), Ponto(2, 2), Ponto(1, 1)];
  final unicos = pontosUnicos(pontos);

  // A asserção que separa a versão sem == e hashCode da correta: sem eles, o
  // Set trata as duas instâncias Ponto(1, 1) como diferentes.
  verificar('pontosUnicos remove duplicata por valor', 2, unicos.length);

  final coordenadas = unicos.map((p) => [p.x, p.y]).toSet();
  verificar('contém (1,1)', true, coordenadas.any((c) => c[0] == 1 && c[1] == 1));
  verificar('contém (2,2)', true, coordenadas.any((c) => c[0] == 2 && c[1] == 2));

  verificar('lista sem duplicatas mantém tamanho', 2,
      pontosUnicos([Ponto(3, 3), Ponto(4, 4)]).length);

  if (_falhas > 0) {
    print('\n$_falhas verificação(ões) falharam');
    exit(1);
  }
  print('\ntodas as verificações passaram');
}
