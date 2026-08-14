# Este exercício não traz conceito novo: combina array, expansão de string e
# aritmética, os três anteriores.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  local registros=(Ana:8 Bruno:5 Cida:9)

  local registro nome nota
  for registro in "${registros[@]}"; do
    nome="${registro%%:*}"     # tudo antes do primeiro ':' (expansão de string)
    nota="${registro#*:}"      # tudo depois do primeiro ':'
    if (( nota >= 6 )); then   # comparação numérica (aritmética)
      echo "$nome: aprovado ($nota)"
    else
      echo "$nome: reprovado ($nota)"
    fi
  done

  # Soma e média, do jeito que já apareceu no exercício de aritmética —
  # "${registros[@]}" é o percorrer de array de sempre.
  local total=0
  for registro in "${registros[@]}"; do
    total=$(( total + ${registro#*:} ))
  done
  echo "media: $(( total / ${#registros[@]} ))"
}

# contar_aprovados extrai a nota de cada registro "nome:nota" e conta as que
# são >= 6.
contar_aprovados() {
  local total=0
  local registro nota
  for registro in "$@"; do
    nota="${registro#*:}"
    (( nota >= 6 )) && (( total++ ))
  done
  echo "$total"
}
