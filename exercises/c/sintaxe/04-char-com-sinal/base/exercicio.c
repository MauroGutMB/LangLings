#include "exercicio.h"

/* Parece direto: se o byte vale 128 ou mais, compare com 128. O detalhe é que
 * char, neste ambiente, é um tipo COM sinal — vai de -128 a 127. Um byte que
 * "deveria" valer 200 não existe como char positivo: ele é armazenado como
 * -56. Comparado com 128, ele nunca vence, porque nenhum char chega lá.
 */
int eh_byte_alto(char c)
{
    return c >= 128;
}
