#include "exercicio.h"

/* Uma macro de função não avalia o argumento: o pré-processador cola o texto
 * dele em cada x, ao pé da letra. Para uma variável isolada, DOBRO(n) vira
 * n + n e funciona. Mas usa_dobro chama DOBRO(a + b) — o texto "a + b" entra
 * em cada x, e a macro expande para "a + b + b", não para "(a + b) + (a + b)".
 * O primeiro operando perdeu o b que deveria ter dobrado junto.
 */
#define DOBRO(x) x + x

int usa_dobro(int a, int b, int fator)
{
    return DOBRO(a + b) * fator;
}
