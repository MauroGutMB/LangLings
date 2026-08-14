#include <stdio.h>

#include "exercicio.h"

/* Um ponteiro guarda um endereço de memória, não um valor comum. & devolve o
   endereço de uma variável; * do outro lado lê (ou escreve) no que esse
   endereço aponta. */
void exemplos(void)
{
    int idade = 30;
    int *p = &idade; /* p guarda o ENDEREÇO de idade, não uma cópia do 30 */

    printf("%d\n", idade);  /* 30 */
    printf("%d\n", *p);     /* 30 — dereferenciar p lê o que está em idade */

    *p = 31; /* escrever através do ponteiro altera idade de verdade */
    printf("%d\n", idade); /* 31 */

    /* Toda variável tem um endereço, e ele geralmente muda a cada execução —
       por isso não há um valor fixo para comparar aqui, só o formato. */
    printf("%d\n", &idade == p); /* 1 — os dois apontam para o mesmo lugar */

    /* Passar por valor: a função recebe uma CÓPIA e não alcança o original. */
    int n = 10;
    int copia = n;
    copia = copia + 1;
    printf("%d %d\n", n, copia); /* 10 11 */

    /* Passar o endereço é o que permite uma função alterar a variável de
       quem chamou — é exatamente o que dobra(p) faz mais abaixo no arquivo. */
    int m = 10;
    dobra(&m);
    printf("%d\n", m); /* 20 */

    /* Um ponteiro nulo não aponta para lugar nenhum válido; dereferenciar um
       NULL é um erro grave. NULL serve para marcar "ainda sem endereço". */
    int *sem_destino = NULL;
    printf("%d\n", sem_destino == NULL); /* 1 */
}

/* dobra lê o valor apontado por p, dobra e escreve de volta no mesmo
 * endereço — é a mesma ideia do *p = 31 lá em cima em exemplos().
 */
void dobra(int *p)
{
    *p = *p * 2;
}
