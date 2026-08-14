#include <stdio.h>

#include "exercicio.h"

/* Os três laços de C fazem a mesma coisa com ergonomias diferentes. A escolha
   é sobre quando a condição é testada e onde o contador vive. */
void exemplos(void)
{
    /* for tem três partes separadas por ponto e vírgula: inicialização,
       condição testada ANTES de cada volta, e passo executado no fim de cada
       volta. O i declarado aqui só existe dentro do laço. */
    for (int i = 0; i < 3; i++) {
        printf("i=%d ", i); /* i=0 i=1 i=2 */
    }
    printf("\n");

    /* O acumulador precisa viver FORA do laço: declarado dentro, ele seria
       criado do zero a cada volta e a soma nunca sobreviveria. */
    int soma = 0;
    for (int i = 1; i <= 5; i++) {
        soma += i;
    }
    printf("%d\n", soma); /* 15 */

    /* while só tem a condição. Se ela já começa falsa, o corpo não roda
       nenhuma vez. */
    int restante = 0;
    while (restante > 0) {
        printf("nunca sai\n");
        restante--;
    }

    /* do/while testa DEPOIS, então o corpo roda pelo menos uma vez. É o
       formato certo quando você precisa ler algo antes de saber se continua. */
    int n = 0;
    do {
        printf("roda uma vez mesmo com n=%d\n", n);
        n++;
    } while (n < 0);

    /* break abandona o laço inteiro; continue pula só o resto desta volta.
       Repare que o passo do for (i++) continua acontecendo depois do continue —
       num while, esquecer de avançar o contador antes do continue trava. */
    for (int i = 0; i < 10; i++) {
        if (i % 2 == 0) {
            continue;
        }
        if (i > 5) {
            break;
        }
        printf("impar=%d ", i); /* impar=1 impar=3 impar=5 */
    }
    printf("\n");
}

/* soma_ate acumula os inteiros de 1 até n.
 *
 * Não há um caso especial para n <= 0: a condição do for é testada antes da
 * primeira volta, então o corpo simplesmente não roda e o acumulador sai com o
 * valor com que entrou.
 */
int soma_ate(int n)
{
    int soma = 0;
    for (int i = 1; i <= n; i++) {
        soma += i;
    }
    return soma;
}
