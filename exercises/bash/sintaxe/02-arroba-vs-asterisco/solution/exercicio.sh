# contar_args devolve a quantidade de argumentos recebidos. "$@" com aspas
# repassa cada argumento como veio, mesmo os que têm espaço dentro.
contar() {
  echo $#
}

contar_args() {
  contar "$@"
}
