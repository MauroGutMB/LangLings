// if/else é a forma clássica; switch expression (Dart 3) é mais direto
// quando o resultado é sempre um valor calculado a partir de casos.
void exemplos() {
  final idade = 20;
  if (idade < 12) {
    print('criança');
  } else if (idade < 18) {
    print('adolescente');
  } else {
    print('adulto'); // adulto
  }

  // switch expression: cada `case` produz um valor: o switch INTEIRO é uma
  // expressão. `when` adiciona uma condição extra ao padrão.
  final temperatura = 28;
  final sensacao = switch (temperatura) {
    < 10 => 'frio',
    >= 10 && < 25 => 'agradável',
    int t when t >= 25 => 'quente',
    _ => 'desconhecido',
  };
  print(sensacao); // quente

  // switch também casa por valor exato, como um if/else if em cadeia.
  final dia = 'sáb';
  final tipo = switch (dia) {
    'sáb' || 'dom' => 'fim de semana',
    _ => 'dia útil',
  };
  print(tipo); // fim de semana
}

// Classifica nota (0 a 10): 'aprovado' se >= 7, 'recuperação' se >= 5 e < 7,
// 'reprovado' caso contrário.
String classificarNota(num nota) {
  if (nota >= 7) return 'aprovado';
  if (nota >= 5) return 'recuperação';
  return 'reprovado';
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
