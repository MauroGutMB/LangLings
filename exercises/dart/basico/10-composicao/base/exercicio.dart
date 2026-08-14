// Um relatório simples: uma lista de maps, cada um representando um
// funcionário, percorrida para montar um resumo em texto.
void exemplos() {
  final funcionarios = <Map<String, Object>>[
    {'nome': 'Ana', 'salario': 5000},
    {'nome': 'Bruno', 'salario': 4200},
    {'nome': 'Cris', 'salario': 6100},
  ];

  num totalFolha = 0;
  for (final f in funcionarios) {
    final nome = f['nome'] as String;
    final salario = f['salario'] as num;
    print('$nome ganha $salario'); // Ana ganha 5000, Bruno ganha 4200, ...
    totalFolha += salario;
  }
  print('folha total: $totalFolha'); // folha total: 15300
}

// SUA VEZ
//
// Cada item de itens é um Map com as chaves 'preco' (num) e 'quantidade'
// (int). Devolva a soma de preco * quantidade de todos os itens. Lista
// vazia devolve 0.
num totalCompra(List<Map<String, num>> itens) {
  return 0; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
