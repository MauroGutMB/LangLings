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

verificar 'etiquetar Ana' 'Nome: Ana' "$(etiquetar Ana)"

# Os espaços internos são o ponto do exercício: um echo sem aspas devolveria
# "Nome: dois espacos", com um espaço só.
verificar 'etiquetar preserva os espacos' 'Nome: dois  espacos' "$(etiquetar 'dois  espacos')"

verificar 'etiquetar texto vazio' 'Nome: ' "$(etiquetar '')"
verificar 'etiquetar acentuado' 'Nome: José da Silva' "$(etiquetar 'José da Silva')"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
