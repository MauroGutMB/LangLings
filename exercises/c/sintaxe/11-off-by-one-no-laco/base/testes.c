#include <stdio.h>

#include "exercicio.h"

static int falhas = 0;

static void verificar_int(const char *o_que, long esperado, long obtido)
{
    if (esperado == obtido) { printf("ok    %s\n", o_que); return; }
    printf("FALHA %s\n      esperado: %ld\n      obtido:   %ld\n", o_que, esperado, obtido);
    falhas++;
}

int main(void)
{
    setvbuf(stdout, NULL, _IOLBF, 0);

    /* Cada array tem UM elemento extra logo depois da posição n-1, com um
       valor bem distante dos outros — se o laço avançar um passo além do
       que devia, esse elemento entra na soma e o resultado desvia bastante
       do esperado. O elemento extra está dentro do array de verdade: não
       há leitura fora dos limites em nenhum caso. */
    int a[4] = {1, 2, 3, 1000};      /* soma_primeiros(a, 3) usa só 1,2,3 */
    int b[3] = {10, 20, 1000};       /* soma_primeiros(b, 2) usa só 10,20 */
    int c[6] = {5, 5, 5, 5, 5, 1000}; /* soma_primeiros(c, 5) usa os 5 cincos */

    verificar_int("soma_primeiros(a, 3)", 6, soma_primeiros(a, 3));
    verificar_int("soma_primeiros(b, 2)", 30, soma_primeiros(b, 2));
    verificar_int("soma_primeiros(c, 5)", 25, soma_primeiros(c, 5));

    int d[2] = {7, 3};
    verificar_int("soma_primeiros(d, 1)", 7, soma_primeiros(d, 1));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
