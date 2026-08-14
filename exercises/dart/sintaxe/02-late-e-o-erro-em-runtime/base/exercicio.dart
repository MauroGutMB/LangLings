// TODO: devolve 'MODO DEBUG', 'MODO PRODUCAO' ou 'MODO PADRAO' quando modo
// for nulo.
String formatarConfig({String? modo}) {
  late String resultado;
  if (modo == 'debug') {
    resultado = 'MODO DEBUG';
  }
  if (modo == 'prod') {
    resultado = 'MODO PRODUCAO';
  }
  return resultado;
}
