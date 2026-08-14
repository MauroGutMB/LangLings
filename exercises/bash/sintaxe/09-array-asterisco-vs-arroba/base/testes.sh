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

verificar 'juntar_virgula um item' 'x' "$(juntar_virgula x)"

# "${itens[*]}" junta com o PRIMEIRO CARACTERE de IFS, que por padrão é
# espaço — não vírgula. Sem ajustar IFS antes, o separador sai errado.
verificar 'juntar_virgula tres itens' 'a,b,c' "$(juntar_virgula a b c)"

# Um item com espaço dentro só sai reconhecível se a junção realmente for
# por vírgula: com espaço como separador, não dá pra saber onde um item
# termina e o outro começa.
verificar 'juntar_virgula item com espaco' 'a b,c' "$(juntar_virgula "a b" c)"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
