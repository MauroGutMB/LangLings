// Devolve uma NOVA lista igual a base, com um 0 a mais no final — sem
// alterar base.
List<int> comZeroExtra(List<int> base) {
  final r = [...base, 0];
  return r;
}
