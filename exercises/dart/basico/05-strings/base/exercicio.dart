// Interpolação evita concatenar com + na mão; String tem uma bateria de
// métodos prontos para os casos comuns.
void exemplos() {
  final nome = 'Ana';
  final idade = 30;
  print('$nome tem $idade anos'); // Ana tem 30 anos
  print('daqui a 1 ano, ${idade + 1}'); // expressão precisa de ${} — daqui a 1 ano, 31

  final frase = '  Dart é divertido  ';
  print(frase.trim()); // Dart é divertido — sem os espaços das pontas
  print(frase.trim().toUpperCase()); // DART É DIVERTIDO

  final palavras = frase.trim().split(' ');
  print(palavras); // [Dart, é, divertido]
  print(palavras.length); // 3

  print(frase.contains('divertido')); // true
  print('Dart'.substring(1)); // art — a partir do índice 1
  print('Dart'.substring(1, 3)); // ar — do índice 1 até (exclusive) 3
}

// SUA VEZ
//
// Devolva as iniciais de cada palavra de nomeCompleto, em maiúsculas, sem
// separador. Palavras separadas por um único espaço. "Ana Maria Silva" -> "AMS".
String iniciais(String nomeCompleto) {
  return ''; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
