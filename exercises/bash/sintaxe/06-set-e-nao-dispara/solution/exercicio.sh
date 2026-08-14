# criar_diretorio está correta e não deve mudar: cria o diretório recebido e
# devolve status 1 (sem criar nada) se o nome vier vazio.
criar_diretorio() {
  [[ -n "$1" ]] || return 1
  mkdir -p "$1"
}

# preparar_e_finalizar imprime "diretorio pronto" e depois "finalizado"
# quando criar_diretorio dá certo, e só "finalizado" — devolvendo status 1 —
# quando criar_diretorio falha.
#
# A checagem é explícita com if/return: set -e não ajuda aqui, porque
# criar_diretorio roda como condição de um && / if, e set -e nunca dispara
# para comandos nessa posição.
preparar_e_finalizar() {
  local alvo="$1"
  if ! criar_diretorio "$alvo"; then
    echo "finalizado"
    return 1
  fi
  echo "diretorio pronto"
  echo "finalizado"
}
