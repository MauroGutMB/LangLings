const _base = ['a', 'b'];

// TODO: devolve uma nova lista com _base mais item no final, sem acumular
// chamadas anteriores.
List<String> comItem(String item) {
  final lista = _base;
  lista.add(item);
  return lista;
}
