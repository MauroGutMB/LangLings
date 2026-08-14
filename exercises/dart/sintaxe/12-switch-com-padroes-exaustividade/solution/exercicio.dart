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

// Calcula a área de f, seja qual for a subclasse de Forma. Sem case
// curinga: Forma é sealed, e o compilador exige que todas as subclasses
// estejam cobertas.
double area(Forma f) {
  return switch (f) {
    Circulo c => math.pi * c.raio * c.raio,
    Quadrado q => q.lado * q.lado,
    Retangulo r => r.largura * r.altura,
  };
}
