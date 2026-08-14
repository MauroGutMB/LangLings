# soma_segura recebe vários valores como argumentos. Cada um pode ser um
# número, uma string vazia (representando "sem valor", que deve ser ignorada
# na soma) ou um texto que não é número nenhum (que deve fazer a função
# falhar: devolver status 1, sem imprimir nada, em vez de inventar um total).
#
# TODO: soma os números certo e ignora os vazios certo. O problema é o que
# acontece com um valor que não é número — e não é vazio.
soma_segura() {
  local total=0
  local v
  for v in "$@"; do
    (( total += v ))
  done
  echo "$total"
}
