# colchetes devolve o texto recebido, exatamente como chegou, entre
# colchetes. A variável precisa estar entre aspas: sem elas, a expansão passa
# pelo word splitting do shell e espaços extras (ou nas pontas) somem.
colchetes() {
  echo "[$1]"
}
