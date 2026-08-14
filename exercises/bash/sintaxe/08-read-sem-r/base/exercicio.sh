# ler_linha lê uma linha do stdin e a devolve exatamente como veio, sem
# alterar nada nela.
#
# TODO: funciona para texto comum e altera o conteúdo quando a linha tem
# barra invertida dentro. Corrija.
ler_linha() {
  local linha
  read linha
  echo "$linha"
}
