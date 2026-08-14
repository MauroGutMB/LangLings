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

# buscar_config deve devolver o valor da chave (via echo) quando ela existe,
# e status 1 sem imprimir nada quando não existe.
#
# TODO: para as chaves que existem funciona. Para as que não existem, a
# função deveria devolver status 1 e não devolve. Corrija.
buscar_config() {
  local valor=$(buscar_no_dicionario "$1")
  if [[ $? -ne 0 ]]; then
    return 1
  fi
  echo "$valor"
}
