class Ponto {
  final int x;
  final int y;
  Ponto(this.x, this.y);
}

// TODO: devolve um Set sem pontos duplicados — dois pontos com o mesmo x e
// o mesmo y contam como o mesmo ponto.
Set<Ponto> pontosUnicos(List<Ponto> pontos) => pontos.toSet();
