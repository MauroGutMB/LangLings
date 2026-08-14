# contar_pares recebe números como argumentos e devolve quantos deles são
# pares.
#
# < <(...) (process substitution) alimenta o while sem colocá-lo no fim de um
# pipe: o laço continua rodando na shell atual, e "contador" sobrevive.
contar_pares() {
  local contador=0
  while read -r n; do
    if (( n % 2 == 0 )); then
      (( contador++ ))
    fi
  done < <(printf '%s\n' "$@")
  echo "$contador"
}
