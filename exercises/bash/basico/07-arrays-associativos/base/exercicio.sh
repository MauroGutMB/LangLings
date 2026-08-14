# Um array associativo usa texto como chave, não posição. Precisa ser
# declarado explicitamente com -A antes de receber a primeira chave.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  local -A idades
  idades[Ana]=30
  idades[Bruno]=25

  echo "${idades[Ana]}"        # 30
  echo "${#idades[@]}"         # 2   <- quantidade de pares

  # ${!m[@]} dá as CHAVES; ${m[@]} dá os VALORES. A ordem entre elas não é
  # garantida ser a de inserção.
  local chave
  for chave in "${!idades[@]}"; do
    echo "$chave tem ${idades[$chave]} anos"
  done

  # -v testa se a chave existe, mesmo que o valor associado seja vazio.
  if [[ -v idades[Ana] ]]; then
    echo "Ana esta cadastrada"
  fi

  # Ler uma chave ausente não dá erro: devolve vazio, ou o padrão do :-.
  echo "${idades[Duda]:-desconhecida}"   # desconhecida
}

# SUA VEZ
#
# idade_de recebe um nome procurado e, depois dele, pares nome/idade
# (nome idade nome idade ...). Devolva a idade correspondente ao nome
# procurado, ou "desconhecida" se ele não aparecer em nenhum par.
idade_de() {
  local procurado="$1"
  echo "$procurado"   # <- troque isto
}
