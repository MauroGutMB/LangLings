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

status_de() {  # status_de <comando...> — roda e devolve $? como texto
  "$@" >/dev/null 2>&1
  echo $?
}

verificar 'iguais 7 e 7' '0' "$(status_de iguais 7 7)"
verificar 'iguais 3 e 4' '1' "$(status_de iguais 3 4)"

# == compara TEXTO: "007" e "7" são strings diferentes, mesmo sendo o mesmo
# número.
verificar 'iguais 007 e 7' '0' "$(status_de iguais 007 7)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
