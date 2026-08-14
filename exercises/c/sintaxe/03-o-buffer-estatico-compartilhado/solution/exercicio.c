#include <ctype.h>
#include <stdlib.h>
#include <string.h>

#include "exercicio.h"

/* Cada chamada aloca seu próprio buffer com malloc, então o resultado de uma
 * chamada não compartilha memória com o de nenhuma outra — diferente do
 * buffer static, que era um único espaço reaproveitado a cada chamada.
 * A troca é um vazamento de memória que o exercício aceita: o processo é
 * curto e o ponto aqui é a independência entre os resultados, não liberação.
 */
char *para_maiusculas(const char *s)
{
    size_t n = strlen(s);
    char *buffer = malloc(n + 1);
    size_t i;
    for (i = 0; i < n; i++) {
        buffer[i] = (char)toupper((unsigned char)s[i]);
    }
    buffer[i] = '\0';
    return buffer;
}
