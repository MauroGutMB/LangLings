#include <string.h>

#include "exercicio.h"

/* Copia os n caracteres de origem e, na posição seguinte, escreve o
 * terminador — strlen nunca o conta, então ele precisa ser escrito à mão.
 */
void copia_string(char *dest, const char *origem)
{
    size_t n = strlen(origem);
    for (size_t i = 0; i < n; i++) {
        dest[i] = origem[i];
    }
    dest[n] = '\0';
}
