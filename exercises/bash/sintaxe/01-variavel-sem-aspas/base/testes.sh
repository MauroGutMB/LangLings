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

verificar 'colchetes palavra simples' '[ola]' "$(colchetes "ola")"

# Duas palavras separadas por um espaço só: a expansão sem aspas ainda dá
# certo aqui, por coincidência (o split cai exatamente onde já havia um
# espaço).
verificar 'colchetes duas palavras' '[ola mundo]' "$(colchetes "ola mundo")"

# Espaço DUPLO entre as palavras é o que separa a versão sem aspas da
# correta: sem aspas, a expansão espalha o texto em palavras e o eco junta
# de volta com um espaço só, apagando o espaço extra.
verificar 'colchetes espaco duplo' '[a  b]' "$(colchetes "a  b")"

# Espaço nas pontas do texto: sem aspas ele é descartado inteiro.
verificar 'colchetes espacos nas pontas' '[  ana  ]' "$(colchetes "  ana  ")"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
