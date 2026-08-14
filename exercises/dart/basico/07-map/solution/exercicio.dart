// Map é uma coleção de pares chave -> valor. Diferente de listas, não tem
// ordem garantida de inserção ao percorrer (embora LinkedHashMap, o Map
// literal padrão, preserve a ordem de inserção na prática).
void exemplos() {
  final idades = <String, int>{'ana': 30};
  idades['bruno'] = 25; // insere ou sobrescreve
  print(idades['ana']); // 30

  // Ler uma chave ausente NÃO é erro: devolve null (o tipo do valor precisa
  // aceitar isso, por isso o Map é <String, int?> implicitamente ao ler).
  print(idades['zoe']); // null

  print(idades.containsKey('zoe')); // false — a forma de distinguir
  // "ausente" de "presente valendo algo falsy" (0, '', etc.)

  idades.remove('ana');
  print(idades.length); // 1

  // update aplica uma função ao valor existente, ou usa ifAbsent se a chave
  // não existir ainda — é o idioma para contagem sem checar antes.
  final contagem = <String, int>{};
  for (final p in ['a', 'b', 'a']) {
    contagem.update(p, (v) => v + 1, ifAbsent: () => 1);
  }
  print(contagem); // {a: 2, b: 1}

  var total = 0;
  for (final entrada in contagem.entries) {
    total += entrada.value;
  }
  print(total); // 3
}

// Devolve um map com quantas vezes cada palavra aparece em palavras.
Map<String, int> contarPalavras(List<String> palavras) {
  final contagem = <String, int>{};
  for (final p in palavras) {
    contagem.update(p, (v) => v + 1, ifAbsent: () => 1);
  }
  return contagem;
}

// Para ver a saída de exemplos(), rode no shell do exercício ([s]):
//   dart run exercicio.dart
void main() {
  exemplos();
}
