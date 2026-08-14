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

verificar 'saudar_todos Ana'            'Ola, Ana!' "$(saudar_todos Ana)"
verificar 'saudar_todos Ana Bruno Cida' $'Ola, Ana!\nOla, Bruno!\nOla, Cida!' \
  "$(saudar_todos Ana Bruno Cida)"

# Sem argumentos o laço não dá nenhuma volta e nada é impresso.
verificar 'saudar_todos sem argumentos' '' "$(saudar_todos)"

# Um nome com espaço chegou como UM argumento e tem que sair como UMA linha.
verificar 'saudar_todos nome composto' 'Ola, Ana Maria!' "$(saudar_todos 'Ana Maria')"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
