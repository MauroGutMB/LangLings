#include "exercicio.h"

/* "os n primeiros elementos" soa como "até chegar em n" — e i <= n parece
 * traduzir isso direto. Só que os n primeiros são os índices 0 até n-1: com
 * i <= n o laço roda uma vez de mais, somando v[n], que já é o elemento
 * seguinte ao último que deveria entrar na conta.
 */
int soma_primeiros(const int v[], int n)
{
    int total = 0;
    for (int i = 0; i <= n; i++) {
        total += v[i];
    }
    return total;
}
