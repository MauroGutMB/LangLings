const _base = ['a', 'b'];

// Devolve uma nova lista com _base mais item no final, sem acumular chamadas
// anteriores.
List<String> comItem(String item) {
  return [..._base, item];
}
