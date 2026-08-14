// var, final e const parecem intercambiáveis à primeira vista, mas cada um
// promete uma coisa diferente sobre a variável.
void exemplos() {
  var nome = 'Ana'; // var: tipo inferido (String) a partir do literal
  nome = 'Bruno'; // válido — var só fixa o TIPO, não o valor
  print(nome); // Bruno

  final cidade = 'Recife'; // final: só pode ser atribuído uma vez...
  // cidade = 'Olinda'; // ...isto não compilaria
  print(cidade); // Recife

  const pi = 3.14159; // const: valor conhecido em tempo de COMPILAÇÃO
  print(pi); // 3.14159

  // final também aceita um valor calculado em tempo de execução; const não.
  final agora = DateTime.now().year;
  print(agora is int); // true

  // Listas e maps também podem ser const: nesse caso a coleção inteira fica
  // imutável, não só a variável que aponta para ela.
  const cores = ['vermelho', 'verde', 'azul'];
  print(cores.length); // 3
}

// SUA VEZ
//
// Devolva a área de um quadrado de lado `lado`.
double areaDoQuadrado(num lado) {
  return 0; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
