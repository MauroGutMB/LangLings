#include <stdio.h>

#include "exercicio.h"

/* Um struct pode ter outro struct como campo. É a mesma ideia de agrupar
   dados sob um nome, aplicada em mais de um nível. */
void exemplos(void)
{
    struct Retangulo tela = {{0, 0}, 800, 600};

    /* Acessar um campo aninhado é encadear pontos: um para chegar em origem,
       outro para chegar em x dentro dela. */
    printf("%d %d\n", tela.origem.x, tela.origem.y); /* 0 0 */
    printf("%d %d\n", tela.largura, tela.altura);    /* 800 600 */

    /* Escrever num campo aninhado funciona igual: encadeia os pontos até o
       campo final e atribui. */
    tela.origem.x = 100;
    printf("%d\n", tela.origem.x); /* 100 */

    /* Copiar o struct de fora copia o struct de dentro junto — não sobra
       nenhum campo compartilhado entre a cópia e o original. */
    struct Retangulo copia = tela;
    copia.origem.y = 999;
    printf("%d %d\n", tela.origem.y, copia.origem.y); /* 0 999 */

    /* Um Ponto pode ser construído à parte e usado para inicializar o campo
       origem do Retangulo. */
    struct Ponto canto = {5, 5};
    struct Retangulo quadrado = {canto, 50, 50};
    printf("%d\n", quadrado.origem.x); /* 5 */
}

/* deslocar parte de uma cópia local de r (passar por valor já garante isso)
 * e soma dx e dy aos campos de origem, deixando largura e altura intactos.
 */
struct Retangulo deslocar(struct Retangulo r, int dx, int dy)
{
    struct Retangulo novo = r;
    novo.origem.x += dx;
    novo.origem.y += dy;
    return novo;
}
