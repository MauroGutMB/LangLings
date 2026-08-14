// this.x no construtor é atalho para "receba um argumento chamado x e
// atribua direto ao campo x" — poupa o `this.x = x;` manual.
class Pessoa {
  final String nome;
  final int idade;

  Pessoa(this.nome, this.idade);

  // Construtor nomeado: uma forma alternativa de criar a instância, com
  // outro nome e outra lógica de inicialização.
  Pessoa.recemNascido(String nome) : this(nome, 0);

  String apresentar() => '$nome, $idade anos';
}

void exemplos() {
  final ana = Pessoa('Ana', 30);
  print(ana.apresentar()); // Ana, 30 anos

  final bebe = Pessoa.recemNascido('Bruno');
  print(bebe.apresentar()); // Bruno, 0 anos
}

// SUA VEZ
//
// Complete area(), que devolve largura * altura.
class Retangulo {
  final double largura;
  final double altura;

  Retangulo(this.largura, this.altura);

  double area() {
    return 0; // <- troque isto
  }
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
