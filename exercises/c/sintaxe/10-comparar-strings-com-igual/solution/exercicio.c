#include <string.h>

#include "exercicio.h"

/* strcmp compara conteúdo, não endereço, e devolve 0 quando as strings são
 * iguais.
 */
int strings_iguais(const char *a, const char *b)
{
    return strcmp(a, b) == 0;
}
