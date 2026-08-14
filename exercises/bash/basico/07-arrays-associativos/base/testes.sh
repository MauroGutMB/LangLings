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

verificar 'idade_de Bruno' '25' "$(idade_de Bruno Ana 30 Bruno 25 Cida 40)"
verificar 'idade_de Ana' '30' "$(idade_de Ana Ana 30 Bruno 25)"
verificar 'idade_de ausente' 'desconhecida' "$(idade_de Duda Ana 30 Bruno 25)"
verificar 'idade_de sem pares' 'desconhecida' "$(idade_de Ana)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
