# soma_segura recebe vários valores como argumentos. Cada um pode ser um
# número, uma string vazia (representando "sem valor", que deve ser ignorada
# na soma) ou um texto que não é número nenhum (que deve fazer a função
# falhar: devolver status 1, sem imprimir nada, em vez de inventar um total).
#
# Um valor precisa ser validado ANTES de entrar em $(( )): dentro da
# aritmética, um texto que não é número vira o nome de uma variável, e uma
# variável não definida vale 0 — silenciosamente.
soma_segura() {
  local total=0
  local v
  for v in "$@"; do
    [[ -z "$v" ]] && continue
    if [[ ! "$v" =~ ^-?[0-9]+$ ]]; then
      return 1
    fi
    (( total += v ))
  done
  echo "$total"
}
