# contar_pares recebe números como argumentos e devolve quantos deles são
# pares.
#
# TODO: parece direto — soma um contador dentro de um laço — mas o resultado
# nunca sai do zero. Corrija sem trocar o jeito de gerar as linhas com printf.
contar_pares() {
  local contador=0
  printf '%s\n' "$@" | while read -r n; do
    if (( n % 2 == 0 )); then
      (( contador++ ))
    fi
  done
  echo "$contador"
}
