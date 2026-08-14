class Contador {
  final int valor;
  Contador(this.valor);

  // Puro: devolve um Contador NOVO com valor + 1, não muda a própria instância.
  Contador incrementar() => Contador(valor + 1);
}

// Aplica incrementar três vezes seguidas e devolve o valor final.
int incrementarTres(int inicial) {
  var c = Contador(inicial);
  c = c.incrementar();
  c = c.incrementar();
  c = c.incrementar();
  return c.valor;
}
