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

verificar 'buscar_config host' 'localhost' "$(buscar_config host)"
verificar 'buscar_config host status' '0' "$(status_de buscar_config host)"

# "local valor=$(...)" faz o status de "local" (sempre 0 quando a atribuição
# funciona) sobrescrever o status de buscar_no_dicionario — o "$? -ne 0" logo
# depois nunca é verdadeiro, mesmo quando a chave não existe.
verificar 'buscar_config chave inexistente status' '1' "$(status_de buscar_config chave_qualquer)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
