// Set é uma coleção sem ordem garantida e sem repetição: adicionar um valor
// já presente não faz nada.
void exemplos() {
  final cores = <String>{'vermelho', 'verde'};
  cores.add('azul');
  cores.add('vermelho'); // já existe — ignorado silenciosamente
  print(cores.length); // 3

  print(cores.contains('verde')); // true

  final primarias = {'vermelho', 'azul', 'amarelo'};
  print(cores.union(primarias)); // {vermelho, verde, azul, amarelo} (ordem pode variar)
  print(cores.intersection(primarias).length); // 2 — vermelho e azul
  print(cores.difference(primarias)); // {verde} — só em cores, não em primarias

  // toSet() é a forma usual de remover duplicatas de uma lista.
  final repetidos = [1, 2, 2, 3, 3, 3];
  print(repetidos.toSet().length); // 3
}

// SUA VEZ
//
// Devolva um Set com os valores que aparecem em a e em b.
Set<int> elementosComuns(List<int> a, List<int> b) {
  return {}; // <- troque isto
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
