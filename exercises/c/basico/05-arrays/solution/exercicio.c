#include <stdio.h>

#include "exercicio.h"

/* Um array em C é um bloco contíguo de elementos do mesmo tipo, com tamanho
   fixado na declaração. Ele não guarda o próprio comprimento em lugar nenhum. */
void exemplos(void)
{
    int notas[5] = {7, 9, 4, 10, 6};

    /* Índices vão de 0 a n-1. notas[5] já está fora do array. */
    printf("%d %d\n", notas[0], notas[4]); /* 7 6 */

    notas[2] = 8; /* escrever é igual a ler, do outro lado do sinal de igual */
    printf("%d\n", notas[2]); /* 8 */

    /* Um inicializador mais curto que o array zera o resto. É a forma
       garantida de começar com tudo em zero: sem inicializador nenhum, os
       elementos de um array local ficam com lixo. */
    int contadores[4] = {0};
    printf("%d %d\n", contadores[0], contadores[3]); /* 0 0 */

    /* Aqui, onde o array foi declarado, sizeof enxerga o bloco inteiro, e a
       divisão pelo tamanho de um elemento dá o número de elementos. Guarde a
       ressalva: isso só funciona onde o compilador vê a declaração. */
    int quantos = (int)(sizeof(notas) / sizeof(notas[0]));
    printf("%d\n", quantos); /* 5 */

    int soma = 0;
    for (int i = 0; i < quantos; i++) {
        soma += notas[i];
    }
    printf("%d\n", soma); /* 40 */

    /* Uma função que recebe array recebe também o comprimento, como
       maior_valor(v, n) faz: do lado de dentro não há como perguntar quantos
       elementos vieram. */
    printf("%d\n", maior_valor(notas, quantos));

    /* Arrays de várias dimensões são arrays de arrays, guardados linha a
       linha na memória. */
    int matriz[2][3] = {{1, 2, 3}, {4, 5, 6}};
    printf("%d\n", matriz[1][2]); /* 6 */
}

/* maior_valor devolve o maior entre os n primeiros elementos de v.
 *
 * O candidato começa em v[0], não em 0: com um array só de negativos, um
 * candidato inicial zerado sobreviveria a todas as comparações e a função
 * devolveria um valor que nem está no array.
 */
int maior_valor(const int v[], int n)
{
    int maior = v[0];
    for (int i = 1; i < n; i++) {
        if (v[i] > maior) {
            maior = v[i];
        }
    }
    return maior;
}
