# criar_diretorio está correta e não deve mudar: cria o diretório recebido e
# devolve status 1 (sem criar nada) se o nome vier vazio.
criar_diretorio() {
  [[ -n "$1" ]] || return 1
  mkdir -p "$1"
}

# preparar_e_finalizar deve imprimir "diretorio pronto" e depois "finalizado"
# quando criar_diretorio dá certo, e SÓ "finalizado" — devolvendo status 1 —
# quando criar_diretorio falha.
#
# TODO: com `set -e` ligado, dava pra imaginar que uma falha no meio já para
# a função sozinha. Rode os testes: a saída de texto está certa, mas o status
# devolvido no caso de falha não é o que o resto do código espera.
preparar_e_finalizar() {
  set -e
  criar_diretorio "$1" && echo "diretorio pronto"
  echo "finalizado"
}
