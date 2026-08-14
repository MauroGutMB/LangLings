# juntar_virgula recebe os itens como argumentos e devolve todos juntos numa
# única string, separados por vírgula.
#
# "${itens[*]}" junta os elementos com o primeiro caractere de IFS — ajustar
# IFS para "," antes de usar [*] é o que controla o separador.
juntar_virgula() {
  local -a itens=("$@")
  local IFS=,
  echo "${itens[*]}"
}
