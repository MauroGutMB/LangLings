#include "exercicio.h"

/* *p++ lê como "incrementa o que p aponta" para quem espera que * e ++ se
 * combinem naquele valor. Mas ++ posfixo tem precedência maior que *: a
 * expressão é *(p++), não (*p)++. Ela lê *p, descarta o resultado, e avança
 * p — que é uma cópia local do ponteiro, então nem esse avanço sobrevive à
 * função. O inteiro que p aponta nunca é escrito.
 */
void incrementa_valor(int *p)
{
    *p++;
}
