#include "exercicio.h"

/* Parênteses ao redor de x, nos dois usos, isolam qualquer expressão que
 * chegue como argumento — e ao redor da macro inteira, isolam o resultado de
 * qualquer operador que venha depois dela (como o * fator de usa_dobro).
 */
#define DOBRO(x) ((x) + (x))

int usa_dobro(int a, int b, int fator)
{
    return DOBRO(a + b) * fator;
}
