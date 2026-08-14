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
    /* Sem isto o stdout sai em bloco no fim do processo e as falhas apareceriam
       DEPOIS da mensagem de erro do make, que vem por stderr. */
    setvbuf(stdout, NULL, _IOLBF, 0);

    int a[5] = {1, 2, 3, 4, 5};
    int b[1] = {42};
    int c[4] = {-1, -2, -3, -4};

    verificar_int("soma(a, 5)", 15, soma(a, 5));
    verificar_int("soma(b, 1)", 42, soma(b, 1));
    verificar_int("soma(c, 4)", -10, soma(c, 4));
    verificar_int("soma(a, 2)", 3, soma(a, 2)); /* n menor que o array */

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
