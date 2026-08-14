class Ponto {
  final int x;
  final int y;
  Ponto(this.x, this.y);

  @override
  bool operator ==(Object other) => other is Ponto && x == other.x && y == other.y;

  @override
  int get hashCode => Object.hash(x, y);
}

// Devolve um Set sem pontos duplicados — dois pontos com o mesmo x e o
// mesmo y contam como o mesmo ponto.
Set<Ponto> pontosUnicos(List<Ponto> pontos) => pontos.toSet();
