#include "exercicio.h"

/* a e b são ponteiros (char *), e == entre ponteiros compara ENDEREÇOS, não
 * o que está neles. Duas strings com o mesmo texto guardadas em lugares
 * diferentes da memória são "diferentes" para ==, mesmo sendo idênticas
 * caractere a caractere.
 */
int strings_iguais(const char *a, const char *b)
{
    return a == b;
}
