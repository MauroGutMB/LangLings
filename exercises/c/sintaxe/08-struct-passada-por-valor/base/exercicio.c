#include "exercicio.h"

/* p já é um ponteiro, então parece seguro tirar uma cópia local para
 * trabalhar nela — afinal "p aponta para o Ponto de verdade". O detalhe é
 * que *p copia o struct inteiro para copia, uma variável comum: alterar
 * copia.x e copia.y só altera essa cópia, que desaparece no fim da função
 * sem nunca ter escrito de volta no Ponto original.
 */
void mover(struct Ponto *p, int dx, int dy)
{
    struct Ponto copia = *p;
    copia.x += dx;
    copia.y += dy;
}
