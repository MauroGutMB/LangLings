# Em Bash não existe "expressão booleana": o if executa um COMANDO e olha o
# status de saída dele. [[ ... ]] é um comando embutido que devolve 0 quando a
# condição vale e 1 quando não vale.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  local nome="Ana"
  local vazio=""

  # Testes de texto.
  [[ -n $nome ]]  && echo "nome tem conteudo"   # -n: não vazio
  [[ -z $vazio ]] && echo "vazio esta vazio"    # -z: vazio
  [[ $nome == "Ana" ]] && echo "igual a Ana"
  [[ $nome != "Bruno" ]] && echo "diferente de Bruno"

  # Testes numéricos usam palavras, não símbolos: -eq -ne -lt -le -gt -ge.
  # Dentro de [[ ]] o < e o > existem, mas comparam TEXTO — e em texto "10"
  # vem antes de "9".
  local idade=30
  [[ $idade -ge 18 ]] && echo "maior de idade"
  [[ 10 -gt 9 ]] && echo "10 e maior que 9 numericamente"
  [[ 10 < 9 ]]  && echo "e menor que 9 alfabeticamente"

  # Testes de arquivo: -f (arquivo comum), -d (diretório), -e (existe).
  [[ -f ./exercicio.sh ]] && echo "o proprio arquivo existe"
  [[ -d /tmp ]] && echo "/tmp e um diretorio"

  # if/elif/else para mais de dois caminhos.
  if [[ $idade -lt 13 ]]; then
    echo "crianca"
  elif [[ $idade -lt 18 ]]; then
    echo "adolescente"
  else
    echo "adulto"
  fi

  # case compara contra PADRÕES (os mesmos do glob), na ordem escrita, e para
  # no primeiro que casar. O *) final é o "senão".
  local arquivo="relatorio.txt"
  case $arquivo in
    *.txt) echo "texto" ;;
    *.sh)  echo "script" ;;
    *)     echo "desconhecido" ;;
  esac

  # && roda o próximo comando só se o anterior deu certo; || só se deu errado.
  [[ -f /nao/existe ]] || echo "nao achei o arquivo"
}

# SUA VEZ
#
# Devolva "negativo", "zero" ou "positivo" conforme o número recebido.
classificar() {
  local n="$1"
  echo "positivo"   # <- troque isto
}
