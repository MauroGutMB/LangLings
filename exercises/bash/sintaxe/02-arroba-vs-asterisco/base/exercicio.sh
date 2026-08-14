# contar_args deve devolver a QUANTIDADE de argumentos recebidos, mesmo
# quando algum deles tem espaço dentro.
#
# TODO: contar() está certa e não deve mudar. O problema está em como
# contar_args repassa os argumentos para ela.
contar() {
  echo $#
}

contar_args() {
  contar $@
}
