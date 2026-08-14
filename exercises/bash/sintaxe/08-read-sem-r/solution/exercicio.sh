# ler_linha lê uma linha do stdin e a devolve exatamente como veio.
#
# -r desliga o tratamento de barra invertida como escape: read pega a linha
# ao pé da letra, sem descartar nenhum caractere.
ler_linha() {
  local linha
  read -r linha
  echo "$linha"
}
