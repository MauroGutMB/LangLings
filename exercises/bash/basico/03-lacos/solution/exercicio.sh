# Bash tem duas famílias de laço: o for, que percorre uma lista de palavras, e
# o while, que repete enquanto um comando devolver status 0.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  # for sobre uma lista literal de palavras.
  for fruta in maca banana uva; do
    echo "fruta: $fruta"
  done

  # {1..5} é expansão de chaves: o Bash a substitui pelas palavras 1 2 3 4 5
  # ANTES de o for começar. Aceita passo: {0..10..2} dá 0 2 4 6 8 10.
  for i in {1..5}; do
    echo -n "$i "
  done
  echo

  # for aritmético, no formato do C: inicialização, condição, passo. Dentro de
  # (( )) as variáveis dispensam o "$" e a comparação é numérica de verdade.
  local n=3
  for (( i = 0; i < n; i++ )); do
    echo "volta $i"
  done

  # while repete enquanto o comando da condição devolver 0. Aqui a condição é
  # o próprio (( )), que devolve 0 quando a conta dá diferente de zero.
  local restante=3
  while (( restante > 0 )); do
    echo "faltam $restante"
    (( restante-- ))
  done

  # continue pula para a próxima volta; break abandona o laço inteiro.
  for i in {1..10}; do
    (( i % 2 == 0 )) && continue   # ignora os pares
    (( i > 7 )) && break           # e para depois do 7
    echo "impar: $i"
  done
}

# contagem_regressiva imprime de n até 1, uma linha por número.
#
# O for aritmético já dá conta do caso n <= 0: a condição é avaliada antes da
# primeira volta, então o corpo simplesmente não roda. Um if extra para isso
# seria código morto.
contagem_regressiva() {
  local n="$1"
  for (( i = n; i >= 1; i-- )); do
    echo "$i"
  done
}
