import 'dart:math' as math;

sealed class Forma {}

class Circulo extends Forma {
  final double raio;
  Circulo(this.raio);
}

class Quadrado extends Forma {
  final double lado;
  Quadrado(this.lado);
}

class Retangulo extends Forma {
  final double largura;
  final double altura;
  Retangulo(this.largura, this.altura);
}

// TODO: calcula a área de f, seja qual for a subclasse de Forma.
double area(Forma f) {
  return switch (f) {
    Circulo c => math.pi * c.raio * c.raio,
    Quadrado q => q.lado * q.lado,
    _ => 0,
  };
}
