# calcular_bonus devolve, como VALOR pela saída padrão, o dobro do salário
# base recebido.
#
# return só devolve STATUS (0 a 255, truncado nesse intervalo) — não serve
# para comunicar um número calculado. echo é quem entrega um valor para quem
# chama capturar com $(...).
calcular_bonus() {
  local base="$1"
  echo $(( base * 2 ))
}
