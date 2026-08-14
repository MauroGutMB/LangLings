# eh_producao NÃO imprime nada: devolve status 0 se o argumento recebido for
# exatamente o texto "prod", e status diferente de 0 em qualquer outro caso —
# sem quebrar não importa o que vier no argumento.
#
# TODO: funciona para os casos óbvios e se comporta mal quando o argumento
# não é uma palavra só. Corrija.
eh_producao() {
  [ $1 = "prod" ]
}
