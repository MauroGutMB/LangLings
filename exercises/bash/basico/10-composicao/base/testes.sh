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

verificar 'contar_aprovados misto' '2' "$(contar_aprovados Ana:8 Bruno:5 Cida:9)"
verificar 'contar_aprovados todos reprovados' '0' "$(contar_aprovados Ana:3 Bruno:5)"
verificar 'contar_aprovados todos aprovados' '3' "$(contar_aprovados Ana:8 Bruno:6 Cida:10)"
verificar 'contar_aprovados sem argumentos' '0' "$(contar_aprovados)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
