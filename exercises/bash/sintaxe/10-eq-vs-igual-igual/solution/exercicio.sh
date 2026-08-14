# iguais devolve, via status, se dois números representam o mesmo valor.
# -eq compara NÚMEROS: converte os dois lados antes de comparar, então "007"
# e "7" são iguais.
iguais() {
  [[ "$1" -eq "$2" ]]
}
