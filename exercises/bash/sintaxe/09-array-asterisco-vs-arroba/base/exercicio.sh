# juntar_virgula recebe os itens como argumentos e devolve todos juntos numa
# única string, separados por vírgula — sem vírgula sobrando nas pontas.
#
# TODO: para um item só funciona. Para mais de um, o separador não é o que
# devia ser. Corrija.
juntar_virgula() {
  local -a itens=("$@")
  echo "${itens[*]}"
}
