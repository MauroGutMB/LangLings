class Contador {
  final int valor;
  Contador(this.valor);

  // Puro: devolve um Contador NOVO com valor + 1, não muda a própria instância.
  Contador incrementar() => Contador(valor + 1);
}

// TODO: aplica incrementar três vezes seguidas e devolve o valor final.
int incrementarTres(int inicial) {
  final c = Contador(inicial)..incrementar()..incrementar()..incrementar();
  return c.valor;
}
