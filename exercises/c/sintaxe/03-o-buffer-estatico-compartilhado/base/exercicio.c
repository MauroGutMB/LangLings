#include <ctype.h>
#include <string.h>

#include "exercicio.h"

/* static dentro de uma função faz o buffer existir uma única vez, criado na
 * primeira chamada e reaproveitado em todas as seguintes — é isso que permite
 * devolver um ponteiro para ele sem ele "morrer" quando a função retorna. O
 * problema é justamente esse reaproveitamento: toda chamada escreve por cima
 * do que a chamada anterior deixou lá, então dois ponteiros devolvidos em
 * momentos diferentes acabam apontando para o mesmo conteúdo — o mais
 * recente.
 */
char *para_maiusculas(const char *s)
{
    static char buffer[64];
    size_t i;
    for (i = 0; s[i] != '\0' && i < sizeof(buffer) - 1; i++) {
        buffer[i] = (char)toupper((unsigned char)s[i]);
    }
    buffer[i] = '\0';
    return buffer;
}
