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

# O último estágio de um pipeline roda numa subshell: incrementar "contador"
# lá dentro não altera o "contador" de fora, que continua em 0.
verificar 'contar_pares mistura' '2' "$(contar_pares 1 2 3 4)"
verificar 'contar_pares todos pares' '4' "$(contar_pares 2 4 6 8)"
verificar 'contar_pares nenhum par' '0' "$(contar_pares 1 3 5)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
