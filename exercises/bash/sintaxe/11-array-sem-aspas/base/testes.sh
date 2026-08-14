# Sem `set -e`: queremos acumular todas as falhas, não parar na primeira.
set -uo pipefail

if [[ ! -f ./exercicio.sh ]]; then
  echo "exercicio.sh não encontrado" >&2
  exit 1
fi
source ./exercicio.sh

falhas=0

verificar() {  # verificar <descrição> <esperado> <obtido>
  local o_que=$1 esperado=$2 obtido=$3
  if [[ "$esperado" == "$obtido" ]]; then
    printf 'ok    %s\n' "$o_que"
  else
    printf 'FALHA %s\n      esperado: %q\n      obtido:   %q\n' "$o_que" "$esperado" "$obtido"
    falhas=$((falhas + 1))
  fi
}

verificar 'contar_itens tres simples' '3' "$(contar_itens a b c)"

# "${itens[@]}" sem aspas reparte cada elemento com espaço pelo espaço,
# antes mesmo de chegar em tamanho: "a b" e "c" viram três palavras, não
# duas.
verificar 'contar_itens com espaco' '2' "$(contar_itens "a b" c)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
