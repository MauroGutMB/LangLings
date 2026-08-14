#include "exercicio.h"

/* Os n primeiros elementos são os índices 0 até n-1: a condição de
 * continuação precisa ser i < n.
 */
int soma_primeiros(const int v[], int n)
{
    int total = 0;
    for (int i = 0; i < n; i++) {
        total += v[i];
    }
    return total;
}
