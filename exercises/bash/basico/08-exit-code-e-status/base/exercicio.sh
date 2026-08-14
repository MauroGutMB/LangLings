# Todo comando termina com um STATUS numérico: 0 é sucesso, qualquer outro
# valor (1 a 255) é algum tipo de falha. $? guarda o status do último comando.
#
# Este arquivo só DEFINE funções. Para ver os exemplos, abra o shell com [s]:
#   source ./exercicio.sh && exemplos
exemplos() {
  true
  echo "true: $?"      # 0

  false
  echo "false: $?"      # 1

  # $? é sobrescrito a CADA comando. Se você quer guardá-lo, capture antes
  # que outro comando (até um echo) tome o lugar dele.
  false
  local status=$?
  echo "guardado antes que sumisse: $status"

  # && só roda o comando da direita se o da esquerda deu status 0.
  # || só roda o da direita se o da esquerda deu status diferente de 0.
  true && echo "roda: true deu certo"
  false || echo "roda: false deu errado"

  # Uma função também tem status: sem return explícito, é o do último comando
  # executado dentro dela — por isso um [[ ]] sozinho no corpo já basta.
  eh_positivo() {
    [[ $1 -gt 0 ]]
  }
  eh_positivo 5
  echo "eh_positivo 5: $?"      # 0
  eh_positivo -3
  echo "eh_positivo -3: $?"     # 1
}

# SUA VEZ
#
# numero_par NÃO imprime nada: devolve status 0 se o argumento for um inteiro
# par, e status diferente de 0 se for ímpar.
numero_par() {
  return 1   # <- troque isto
}
