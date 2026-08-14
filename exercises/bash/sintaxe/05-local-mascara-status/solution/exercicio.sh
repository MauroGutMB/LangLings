# buscar_no_dicionario está correta e não deve mudar: imprime o valor e
# devolve status 0 quando a chave existe, e devolve status 1 (sem imprimir
# nada) quando não existe.
buscar_no_dicionario() {
  case "$1" in
    host) echo "localhost"; return 0 ;;
    porta) echo "8080"; return 0 ;;
    *) return 1 ;;
  esac
}

# buscar_config devolve o valor da chave quando ela existe, e status 1 sem
# imprimir nada quando não existe.
#
# "local" e a atribuição são separadas: assim $? logo depois é o de
# buscar_no_dicionario, não o de local.
buscar_config() {
  local valor
  valor=$(buscar_no_dicionario "$1")
  if [[ $? -ne 0 ]]; then
    return 1
  fi
  echo "$valor"
}
