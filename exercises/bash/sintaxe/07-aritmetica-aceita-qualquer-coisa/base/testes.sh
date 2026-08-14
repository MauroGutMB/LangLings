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

verificar 'soma_segura numeros' '6' "$(soma_segura 1 2 3)"
verificar 'soma_segura ignora vazio' '4' "$(soma_segura 1 "" 3)"

# Dentro de $(( )), um texto que não é número é tratado como NOME DE
# VARIÁVEL, e uma variável não definida vale 0 — a conta segue sem erro
# nenhum e "abc" vira zero silenciosamente, mascarando um dado inválido.
verificar 'soma_segura valor invalido: saida' '' "$(soma_segura 1 abc 3)"
verificar 'soma_segura valor invalido: status' '1' "$(status_de soma_segura 1 abc 3)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
