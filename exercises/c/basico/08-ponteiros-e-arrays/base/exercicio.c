#include <stdio.h>

#include "exercicio.h"

/* Na maioria dos contextos, o nome de um array decai para um ponteiro ao seu
   primeiro elemento. É por isso que v[i] e *(v + i) fazem a mesma leitura. */
void exemplos(void)
{
    int notas[4] = {7, 9, 4, 10};

    printf("%d\n", notas[2]);     /* 4 */
    printf("%d\n", *(notas + 2)); /* 4 — mesma leitura, escrita diferente */

    /* notas sozinho, aqui, decai para &notas[0]. */
    int *p = notas;
    printf("%d\n", *p); /* 7 */

    /* Avançar o ponteiro anda um elemento inteiro por vez, não um byte:
       p + 1 aponta para notas[1], não para o byte seguinte de notas[0]. */
    p = p + 1;
    printf("%d\n", *p); /* 9 */

    /* ++ num ponteiro faz a mesma conta: avança um elemento. */
    p++;
    printf("%d\n", *p); /* 4 */

    /* Subtrair dois ponteiros do mesmo array dá a distância em elementos,
       não em bytes. */
    int *inicio = notas;
    int *fim = notas + 4;
    printf("%td\n", fim - inicio); /* 4 */

    /* sizeof, porém, só enxerga o array inteiro onde a declaração está
       visível — dentro de uma função que só recebe o ponteiro, essa conta
       não funciona mais (é a armadilha do exercício de sintaxe sobre isso). */
    printf("%zu\n", sizeof(notas) / sizeof(notas[0])); /* 4 */
}

/* SUA VEZ
 *
 * Devolva a soma dos n primeiros elementos de v, percorrendo com aritmética
 * de ponteiro (avançando um ponteiro) em vez de indexar com v[i].
 * Nunca é chamada com n igual a 0.
 */
int soma(const int *v, int n)
{
    return 0; /* <- troque isto */
}
