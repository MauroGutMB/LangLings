#include <string.h>

#include "exercicio.h"

/* strlen(origem) diz quantos caracteres existem antes do terminador — e o
 * laço copia exatamente esses. Parece completo: todo caractere "de verdade"
 * foi copiado. O que fica de fora é o próprio terminador, que não é contado
 * por strlen mas que toda outra função de string (inclusive strcmp, usada
 * para conferir o resultado) exige para saber onde a string acaba.
 */
void copia_string(char *dest, const char *origem)
{
    size_t n = strlen(origem);
    for (size_t i = 0; i < n; i++) {
        dest[i] = origem[i];
    }
}
