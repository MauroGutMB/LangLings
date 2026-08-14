#include "exercicio.h"

/* a / b parece bastar: para números positivos, arredondar para baixo é
 * exatamente o que a divisão de inteiros de C já faz. O que essa suposição
 * esconde é que a divisão de C trunca em direção a ZERO, não em direção ao
 * menos infinito — as duas coincidem só quando a divisão é exata ou quando
 * a e b têm o mesmo sinal.
 */
int divisao_para_baixo(int a, int b)
{
    return a / b;
}
