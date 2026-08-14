#include <stdio.h>

#include "exercicio.h"

/* C não tem tipo booleano nativo no núcleo da linguagem: condição é um número.
   Zero é falso, qualquer outro valor é verdadeiro. */
void exemplos(void)
{
    int idade = 20;

    /* if / else if / else: as condições são testadas de cima para baixo e a
       execução para na primeira verdadeira. Por isso a ordem importa — se o
       primeiro teste fosse >= 0, nenhum dos outros seria alcançado. */
    if (idade >= 60) {
        printf("idoso\n");
    } else if (idade >= 18) {
        printf("adulto\n"); /* é este que sai */
    } else {
        printf("menor\n");
    }

    /* Os operadores de comparação devolvem 1 ou 0, e dá para imprimi-los. */
    printf("%d %d %d\n", idade == 20, idade != 20, idade > 100); /* 1 0 0 */

    /* && e || avaliam da esquerda para a direita e param assim que o resultado
       já está decidido. Isso é garantido pela linguagem, não um detalhe do
       compilador: é o que permite testar um divisor antes de dividir. */
    int divisor = 0;
    if (divisor != 0 && 10 / divisor > 1) {
        printf("nunca chega aqui\n");
    }

    /* switch compara um valor inteiro (ou char) contra rótulos constantes.
       Sem o break, a execução ESCORREGA para o próximo case — o que é útil
       para agrupar rótulos, como em 6 e 7 aqui embaixo. */
    int dia = 7;
    switch (dia) {
    case 6:
    case 7:
        printf("fim de semana\n"); /* é este que sai */
        break;
    case 1:
        printf("segunda\n");
        break;
    default:
        printf("meio da semana\n");
        break;
    }

    /* O operador ternário é um if que devolve valor, útil quando os dois
       ramos são apenas uma expressão. */
    int maior = idade > 30 ? idade : 30;
    printf("%d\n", maior); /* 30 */
}

/* SUA VEZ
 *
 * Devolva 'A' para nota >= 90, 'B' para >= 80, 'C' para >= 70 e 'F' abaixo.
 */
char conceito(int nota)
{
    return '?'; /* <- troque isto */
}
