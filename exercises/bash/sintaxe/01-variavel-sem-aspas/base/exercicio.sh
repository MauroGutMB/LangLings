# colchetes deve devolver o texto recebido, exatamente como chegou, entre
# colchetes: "abc" vira "[abc]".
#
# TODO: funciona para uma palavra só e quebra quando o texto tem mais de um
# espaço entre as palavras ou espaço nas pontas. Corrija.
colchetes() {
  echo [$1]
}
