#include <stdio.h>

#include "exercicio.h"

/* Em C toda variável tem um tipo fixo, escolhido por você na declaração. O
   tipo decide o que cabe na variável e quantos bytes ela ocupa na memória. */
void exemplos(void)
{
    int idade = 30;      /* inteiro com sinal */
    char inicial = 'A';  /* um único byte; 'A' é o número 65 */
    double altura = 1.75;

    printf("%d %c %f\n", idade, inicial, altura); /* 30 A 1.750000 */

    /* char É um tipo numérico: imprimir com %d mostra o código do caractere. */
    printf("%d\n", inicial); /* 65 */

    /* sizeof responde, em bytes, quanto um tipo ou uma variável ocupa. É um
       operador resolvido pelo compilador, não uma chamada em tempo de execução.
       O tipo do resultado é size_t, por isso o %zu. */
    printf("%zu %zu %zu\n", sizeof(char), sizeof(int), sizeof(double)); /* 1 4 8 */

    /* Divisão entre dois int é divisão inteira: a fração é descartada antes de
       qualquer atribuição. Converter um dos lados muda a conta inteira. */
    printf("%d\n", idade / 4);            /* 7 */
    printf("%f\n", (double)idade / 4);    /* 7.500000 */

    /* int estoura silenciosamente; long tem mais espaço. A conversão precisa
       vir ANTES da multiplicação, senão a conta já aconteceu em int. */
    int grande = 100000;
    printf("%ld\n", (long)grande * grande); /* 10000000000 */

    /* const marca o que não muda; o compilador recusa uma atribuição a ela. */
    const int maioridade = 18;
    printf("%d\n", idade >= maioridade); /* 1 — C não tem bool nativo antes de
                                            <stdbool.h>: verdadeiro é 1 */
}

/* SUA VEZ
 *
 * Devolva quantos bytes ocupam `quantidade` valores do tipo double.
 * bytes_de_doubles(0) é 0.
 */
long bytes_de_doubles(int quantidade)
{
    return 0; /* <- troque isto */
}
