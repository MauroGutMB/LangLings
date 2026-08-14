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

verificar 'contagem_regressiva 3' $'3\n2\n1' "$(contagem_regressiva 3)"
verificar 'contagem_regressiva 1' '1'        "$(contagem_regressiva 1)"

# n = 0 tem que sair sem imprimir nada: um laço cuja condição já nasce falsa
# não executa o corpo nenhuma vez.
verificar 'contagem_regressiva 0' ''         "$(contagem_regressiva 0)"
verificar 'contagem_regressiva -2' ''        "$(contagem_regressiva -2)"
verificar 'contagem_regressiva 5' $'5\n4\n3\n2\n1' "$(contagem_regressiva 5)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
