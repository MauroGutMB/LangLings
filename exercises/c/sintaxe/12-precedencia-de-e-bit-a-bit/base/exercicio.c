#include "exercicio.h"

/* Lida da esquerda para a direita, "flags & mascara == mascara" parece dizer
 * "o E dos dois é igual a mascara". Mas & tem precedência MENOR que ==: o
 * compilador calcula (mascara == mascara) primeiro — que vale 1 sempre — e
 * só depois faz flags & 1. A expressão inteira testa um único bit de flags,
 * não os bits de mascara.
 */
int tem_flag(int flags, int mascara)
{
    return flags & mascara == mascara;
}
