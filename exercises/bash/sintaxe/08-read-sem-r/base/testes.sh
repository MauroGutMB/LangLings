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

verificar 'ler_linha texto simples' 'ola mundo' "$(printf 'ola mundo\n' | ler_linha)"

# Sem -r, read trata a barra invertida como escape: ela some e o caractere
# seguinte fica sozinho. Um caminho estilo Windows perde as barras.
verificar 'ler_linha com barra invertida' 'C:\notas\arquivo.txt' \
  "$(printf 'C:\\notas\\arquivo.txt\n' | ler_linha)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
