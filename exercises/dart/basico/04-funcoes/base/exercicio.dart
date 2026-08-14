// Parâmetros nomeados (entre {}) tornam a chamada autoexplicativa; opcionais
// posicionais (entre []) mantêm a ordem mas dispensam o valor. => é açúcar
// para uma função cujo corpo é uma única expressão.
void exemplos() {
  String apresentar(String nome, {int idade = 0, String cidade = '?'}) {
    return '$nome, $idade anos, de $cidade';
  }

  print(apresentar('Ana')); // Ana, 0 anos, de ?
  print(apresentar('Bruno', idade: 30)); // Bruno, 30 anos, de ?
  print(apresentar('Cris', idade: 25, cidade: 'Recife'));
  // Cris, 25 anos, de Recife — nomeados podem vir em qualquer ordem

  String repetir(String texto, [int vezes = 2]) => texto * vezes;
  print(repetir('ab')); // abab — usa o padrão
  print(repetir('ab', 3)); // ababab — sobrescreve o padrão

  // => funciona para qualquer função de expressão única, incluindo as de
  // nível superior e os métodos de classe.
  int dobro(int n) => n * 2;
  print(dobro(21)); // 42
}

// SUA VEZ
//
// Devolva '$saudacao, $nome!' — saudacao é nomeado, opcional, padrão 'Olá'.
String cumprimentar(String nome, {String saudacao = 'Olá'}) {
  return ''; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
