# Uma função em Bash não declara parâmetros: os parênteses da definição são
# sempre vazios. Os argumentos chegam nas variáveis posicionais $1, $2, ... e a
# chamada é feita como a de um comando, separada por espaços e sem vírgulas.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  apresentar Ana 30          # sem parênteses e sem vírgula na chamada
  apresentar Bruno           # o segundo argumento fica ausente
  contar a b "c d"           # 3 argumentos: as aspas mantêm "c d" inteiro

  # Um VALOR sai pela saída padrão e é capturado com $(...).
  local dobrado
  dobrado=$(dobro 21)
  echo "dobro: $dobrado"     # dobro: 42

  # Um STATUS sai pelo return e é lido em $?. Ele é um código de sucesso ou
  # falha (0 = deu certo), não um número de resultado — por isso o if lê a
  # função direto, sem $(...).
  if eh_maior 10 3; then
    echo "10 e maior"
  fi
  eh_maior 3 10
  echo "status: $?"          # status: 1
}

apresentar() {
  local nome="$1"
  # ${2:-...} usa o valor de $2, ou o texto depois do :- quando $2 está ausente
  # ou vazio. Sem isso, ler um argumento que não veio daria erro no modo
  # estrito que os testes usam.
  local idade="${2:-idade nao informada}"
  echo "$nome ($idade)"
}

contar() {
  echo "recebi $# argumentos"   # $# é a quantidade
  echo "o primeiro e: $1"
  # "$@" expande para a lista de argumentos, cada um como uma palavra separada
  # e inteira. É a forma correta de repassar argumentos adiante.
  for arg in "$@"; do
    echo " - [$arg]"
  done
}

dobro() {
  echo $(( $1 * 2 ))   # devolve VALOR: quem chama captura com $(...)
}

eh_maior() {
  [[ $1 -gt $2 ]]      # devolve STATUS: o do próprio [[ ]], sem return explícito
}

# saudar_todos cumprimenta cada argumento recebido.
#
# "$@" com aspas é o que garante uma linha por argumento: sem elas, um nome
# composto seria quebrado em duas palavras e viraria duas saudações.
saudar_todos() {
  for nome in "$@"; do
    echo "Ola, $nome!"
  done
}
