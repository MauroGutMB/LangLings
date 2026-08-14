# Um array em Bash é uma lista indexada por número, começando em 0. Criar um
# não exige tamanho fixo nem tipo.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  local frutas=(maca banana uva)

  echo "${frutas[0]}"        # maca      <- índice começa em 0
  echo "${#frutas[@]}"       # 3         <- quantidade de elementos

  frutas+=(pera)             # += acrescenta ao final, sem apagar o resto
  echo "${frutas[@]}"        # maca banana uva pera
  echo "${frutas[-1]}"       # pera      <- índice negativo conta do fim

  # "${frutas[@]}" com aspas percorre cada elemento inteiro, mesmo os que têm
  # espaço. Sem aspas, o for reparte cada elemento pelo espaço — o mesmo risco
  # de "$@" sem aspas.
  local fruta
  for fruta in "${frutas[@]}"; do
    echo "- $fruta"
  done

  # Fatiar: ${arr[@]:inicio:quantos}, a mesma ideia do recorte de string.
  echo "${frutas[@]:1:2}"    # banana uva
}

# soma_array soma os números recebidos como argumentos.
soma_array() {
  local total=0
  local n
  for n in "$@"; do
    (( total += n ))
  done
  echo "$total"
}
