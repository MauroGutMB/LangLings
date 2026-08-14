# tamanho está correta e não deve mudar: devolve a quantidade de argumentos
# recebidos.
tamanho() {
  echo $#
}

# contar_itens monta um array a partir dos argumentos recebidos e devolve
# quantos itens ele tem. "${itens[@]}" com aspas repassa cada elemento
# inteiro, mesmo os que têm espaço dentro.
contar_itens() {
  local -a itens=("$@")
  tamanho "${itens[@]}"
}
