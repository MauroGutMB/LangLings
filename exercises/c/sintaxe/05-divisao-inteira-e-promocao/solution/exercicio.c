#include "exercicio.h"

/* A divisão trunca em direção a zero. Quando o resto não é zero e a e b têm
 * sinais diferentes, o quociente truncado ficou uma unidade acima do
 * arredondado para baixo — subtrair 1 corrige.
 */
int divisao_para_baixo(int a, int b)
{
    int quociente = a / b;
    int resto = a % b;
    if (resto != 0 && ((a < 0) != (b < 0))) {
        quociente -= 1;
    }
    return quociente;
}
