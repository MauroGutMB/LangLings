#ifndef EXERCICIO_H
#define EXERCICIO_H

/* Imprime os exemplos comentados. O harness de teste nao chama esta funcao:
   a saida dos exemplos ficaria misturada com o resultado das verificacoes. */
void exemplos(void);

/* SUA VEZ: soma dos n primeiros elementos de v, usando aritmetica de
   ponteiro em vez de indice. Nunca e chamada com n igual a 0. */
int soma(const int *v, int n);

#endif
