// Devolve 'MODO DEBUG', 'MODO PRODUCAO' ou 'MODO PADRAO' quando modo for
// nulo.
String formatarConfig({String? modo}) {
  var resultado = 'MODO PADRAO';
  if (modo == 'debug') {
    resultado = 'MODO DEBUG';
  }
  if (modo == 'prod') {
    resultado = 'MODO PRODUCAO';
  }
  return resultado;
}
