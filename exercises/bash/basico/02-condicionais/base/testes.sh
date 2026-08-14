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

verificar 'classificar 7'    'positivo' "$(classificar 7)"
verificar 'classificar 0'    'zero'     "$(classificar 0)"
verificar 'classificar -3'   'negativo' "$(classificar -3)"
verificar 'classificar 1'    'positivo' "$(classificar 1)"
verificar 'classificar -1'   'negativo' "$(classificar -1)"
verificar 'classificar 1000' 'positivo' "$(classificar 1000)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
