#include <stdio.h>

#include "exercicio.h"

/* Um struct agrupa vários campos, possivelmente de tipos diferentes, sob um
   nome só. Cada instância tem sua própria cópia de cada campo. */
void exemplos(void)
{
    /* Inicializar campo a campo, na ordem da declaração. */
    struct Ponto origem = {0, 0};
    printf("%d %d\n", origem.x, origem.y); /* 0 0 */

    /* O ponto acessa um campo para ler ou escrever — dos dois lados de um
       sinal de igual. */
    struct Ponto p = {3, 4};
    p.x = 10;
    printf("%d %d\n", p.x, p.y); /* 10 4 */

    /* Atribuir um struct a outra variável copia TODOS os campos. Depois
       disso as duas são independentes: mexer numa não afeta a outra. */
    struct Ponto copia = p;
    copia.y = 99;
    printf("%d %d\n", p.y, copia.y); /* 4 99 */

    /* Passar um struct para uma função também copia: a função recebe seus
       próprios x e y, e mexer neles não alcança o struct de quem chamou. */
    printf("%d\n", origem.x + p.x); /* 10 */

    /* criar_ponto (implementada mais abaixo) devolve um struct novo a cada
       chamada — os dois pontos abaixo não compartilham memória nenhuma. */
    struct Ponto a = criar_ponto(1, 1);
    struct Ponto b = criar_ponto(2, 2);
    printf("%d %d\n", a.x, b.x); /* 1 2 */
}

/* SUA VEZ
 *
 * Devolva um Ponto com os campos x e y preenchidos com os valores recebidos.
 */
struct Ponto criar_ponto(int x, int y)
{
    struct Ponto p = {0, 0}; /* <- troque isto */
    return p;
}
