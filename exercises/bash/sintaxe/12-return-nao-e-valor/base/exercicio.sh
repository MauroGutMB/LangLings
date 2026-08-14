# calcular_bonus deve devolver, como VALOR (pela saída padrão, para quem
# chama capturar com $(...)), o dobro do salário base recebido.
#
# TODO: a conta está certa, mas quem chama a função nunca recebe o número.
# Corrija sem mudar a conta em si.
calcular_bonus() {
  local base="$1"
  return $(( base * 2 ))
}
