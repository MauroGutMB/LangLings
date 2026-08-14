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

verificar 'eh_producao prod' '0' "$(status_de eh_producao prod)"
verificar 'eh_producao dev' '1' "$(status_de eh_producao dev)"

# O argumento tem espaço dentro: sem aspas, [ ] recebe palavras demais e
# devolve status 2 (erro de sintaxe do próprio [ ]), não o "1" de "não é
# prod" que a função promete.
verificar 'eh_producao com espaco' '1' "$(status_de eh_producao "prod extra")"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
