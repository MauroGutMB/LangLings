# $(( )) e (( )) avaliam expressões aritméticas. Dentro delas as variáveis
# dispensam o "$" e só existem números inteiros — sem casas decimais.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  echo $(( 7 + 3 ))      # 10
  echo $(( 7 / 2 ))      # 3    <- divisão inteira, sem arredondar
  echo $(( 7 % 2 ))      # 1    <- resto da divisão
  echo $(( 2 ** 10 ))    # 1024 <- potência

  local x=5
  (( x += 3 ))
  echo "$x"               # 8
  (( x++ ))
  echo "$x"               # 9

  # Comparação numérica funciona dentro de (( )), sem -gt/-lt.
  local a=10 b=3
  if (( a > b )); then
    echo "a e maior"
  fi

  # (( )) também é um comando com status: 1 (falso) quando a conta dá zero,
  # 0 (verdadeiro) em qualquer outro caso. É por isso que dá pra usar em if.
  (( 0 )); echo "status de (( 0 )): $?"
  (( 1 )); echo "status de (( 1 )): $?"
}

# media soma os argumentos e divide pela quantidade deles, truncando para
# baixo. Sem argumentos, devolve 0 em vez de dividir por zero.
media() {
  if (( $# == 0 )); then
    echo 0
    return
  fi

  local total=0
  local n
  for n in "$@"; do
    (( total += n ))
  done
  echo $(( total / $# ))
}
