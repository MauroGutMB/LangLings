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

verificar 'sem_extensao relatorio.txt' 'relatorio' "$(sem_extensao relatorio.txt)"

# arquivo.tar.gz é o caso que separa o corte curto (%) do corte longo (%%):
# com %% sobraria só "arquivo".
verificar 'sem_extensao arquivo.tar.gz' 'arquivo.tar' "$(sem_extensao arquivo.tar.gz)"

# Sem ponto não há o que cortar, e o padrão que não casa devolve o texto inteiro.
verificar 'sem_extensao LEIAME' 'LEIAME' "$(sem_extensao LEIAME)"

verificar 'sem_extensao com caminho' '/home/ana/notas' "$(sem_extensao /home/ana/notas.md)"
verificar 'sem_extensao nome com espaco' 'meu relatorio' "$(sem_extensao 'meu relatorio.txt')"

if (( falhas > 0 )); then
  printf '\n%d verificação(ões) falharam\n' "$falhas"
  exit 1
fi
printf '\ntodas as verificações passaram\n'
