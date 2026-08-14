# iguais recebe dois números e devolve status 0 quando eles representam o
# MESMO valor numérico — mesmo que estejam escritos diferente, como "007" e
# "7" — e status diferente de 0 quando são valores diferentes.
#
# TODO: para números escritos do mesmo jeito, funciona. Para números iguais
# escritos diferente, não. Corrija.
iguais() {
  [[ "$1" == "$2" ]]
}
