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

verificar 'preparar_e_finalizar sucesso: saida' $'diretorio pronto\nfinalizado' \
  "$(preparar_e_finalizar pasta_valida)"
verificar 'preparar_e_finalizar sucesso: status' '0' "$(status_de preparar_e_finalizar pasta_valida)"

verificar 'preparar_e_finalizar falha: saida' 'finalizado' "$(preparar_e_finalizar '')"

# criar_diretorio falhou dentro de um "cmd && echo", e set -e NÃO para a
# função por causa disso: comandos que são parte de && / || / if não disparam
# set -e. A função segue até o fim e devolve o status do último comando
# (echo "finalizado", que é 0) em vez de propagar a falha.
verificar 'preparar_e_finalizar falha: status' '1' "$(status_de preparar_e_finalizar '')"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
