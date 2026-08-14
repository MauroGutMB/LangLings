# tamanho está correta e não deve mudar: devolve a quantidade de argumentos
# recebidos.
tamanho() {
  echo $#
}

# contar_itens monta um array a partir dos argumentos recebidos e devolve
# quantos itens ele tem, repassando para tamanho.
#
# TODO: para itens simples funciona. Quando algum item tem espaço dentro, a
# contagem sobe sozinha. Corrija sem mudar tamanho().
contar_itens() {
  local -a itens=("$@")
  tamanho ${itens[@]}
}
