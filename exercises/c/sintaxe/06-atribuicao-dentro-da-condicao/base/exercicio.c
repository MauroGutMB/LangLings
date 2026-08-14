#include "exercicio.h"

/* Lido rápido, "if (x = 5)" parece um teste de igualdade — a diferença de um
 * caractere para == passa despercebida. Só que = é atribuição: x recebe 5, e
 * o valor da expressão inteira é esse 5, que o if enxerga como verdadeiro.
 * O teste entra sempre, para qualquer x que a função receba.
 */
int eh_igual_a_cinco(int x)
{
    if (x = 5) {
        return 1;
    }
    return 0;
}
