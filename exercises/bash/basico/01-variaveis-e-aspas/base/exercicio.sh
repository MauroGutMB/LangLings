# Em Bash toda variável guarda texto. A atribuição não aceita espaço em volta
# do "=", e a leitura pede o "$" na frente do nome.
#
# Este arquivo só DEFINE funções: nada roda quando ele é carregado. Para ver os
# exemplos, abra o shell com [s] e rode:  source ./exercicio.sh && exemplos
exemplos() {
  local nome="Ana"
  local saudacao="Ola"

  echo "$nome"              # Ana
  echo "$saudacao, $nome!"  # Ola, Ana!

  # As chaves delimitam onde o nome da variável termina. Sem elas o Bash
  # procuraria uma variável chamada "nomes", que não existe.
  echo "${nome}s"           # Anas

  # Variável nunca definida expande para vazio, sem erro e sem aviso — o que
  # transforma um nome digitado errado numa string vazia silenciosa.
  echo "[$sobrenome]"       # []

  # Entre aspas duplas o valor chega inteiro, com os espaços que tiver. Sem
  # aspas o Bash primeiro quebra o valor em palavras, e o echo as junta de
  # volta com um espaço só.
  local frase="dois  espacos"
  echo "$frase"             # dois  espacos
  echo $frase               # dois espacos   <- o par de espaços virou um

  # local prende a variável à função. Sem ele a variável seria global e
  # continuaria existindo depois que a função terminasse.
  local visivel_so_aqui="segredo"
  echo "$visivel_so_aqui"   # segredo
}

# SUA VEZ
#
# Devolva "Nome: <texto>", com o texto exatamente como veio.
etiquetar() {
  local texto="$1"
  echo "Nome:"   # <- troque isto
}
