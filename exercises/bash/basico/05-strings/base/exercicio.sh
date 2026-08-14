# Expansão de parâmetro é o canivete de texto do Bash: recortar, substituir e
# dar valor padrão sem chamar sed, cut ou awk. Tudo mora dentro de ${}.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  local caminho="/home/ana/relatorio.final.txt"

  echo "${#caminho}"          # 29  <- tamanho em caracteres

  # Valor padrão: usa o conteúdo se houver, senão o texto depois do :-.
  local apelido=""
  echo "${apelido:-sem apelido}"   # sem apelido
  echo "${apelido}"                # (linha vazia: o :- não altera a variável)

  # Remover PREFIXO com #. O padrão é um glob, não uma expressão regular.
  # Um # corta o menor pedaço que casa; dois ## cortam o maior.
  echo "${caminho#*/}"        # home/ana/relatorio.final.txt
  echo "${caminho##*/}"       # relatorio.final.txt   <- o nome do arquivo

  # Remover SUFIXO com %. Vale a mesma regra: % é o menor, %% é o maior.
  echo "${caminho%.*}"        # /home/ana/relatorio.final
  echo "${caminho%%.*}"       # /home/ana/relatorio

  # Um padrão que não casa com nada devolve o texto inalterado, sem erro.
  local sem_ponto="LEIAME"
  echo "${sem_ponto%.*}"      # LEIAME

  # Substituir: uma barra troca a primeira ocorrência, duas trocam todas.
  local frase="um dois um dois"
  echo "${frase/um/UM}"       # UM dois um dois
  echo "${frase//um/UM}"      # UM dois UM dois

  # Recorte por posição: ${v:inicio:quantos}. A contagem começa no zero.
  echo "${frase:0:2}"         # um
  echo "${frase:3}"           # dois um dois

  # Caixa alta e caixa baixa.
  local nome="Ana"
  echo "${nome^^}"            # ANA
  echo "${nome,,}"            # ana
}

# SUA VEZ
#
# Devolva o nome de arquivo sem a última extensão.
sem_extensao() {
  local arquivo="$1"
  echo "$arquivo"   # <- troque isto
}
