// TODO: devolve quantos elementos de ns foram passados pela função de
// transformação — ou seja, o tamanho de ns.
int contarProcessados(List<int> ns) {
  var contador = 0;
  ns.map((n) {
    contador++;
    return n * 2;
  });
  return contador;
}
