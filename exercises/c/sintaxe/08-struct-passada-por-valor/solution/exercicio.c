#include "exercicio.h"

/* Escrever através do ponteiro, sem passar por uma cópia local, alcança o
 * Ponto de quem chamou.
 */
void mover(struct Ponto *p, int dx, int dy)
{
    p->x += dx;
    p->y += dy;
}
