#include "exercicio.h"

/* Os parênteses forçam a dereferência primeiro: (*p)++ incrementa o inteiro
 * apontado, não o ponteiro.
 */
void incrementa_valor(int *p)
{
    (*p)++;
}
