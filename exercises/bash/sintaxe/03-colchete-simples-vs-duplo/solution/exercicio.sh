# eh_producao devolve, via status, se o argumento é exatamente "prod".
# [[ ]] não faz word splitting nem expansão de glob no que está dentro dele:
# o valor de $1 é comparado como uma coisa só, tenha espaço ou não.
eh_producao() {
  [[ $1 = "prod" ]]
}
