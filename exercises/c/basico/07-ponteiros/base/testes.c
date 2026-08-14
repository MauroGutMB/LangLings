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

    int x = 5;
    dobra(&x);
    verificar_int("dobra(&x) com x=5", 10, x);

    int y = 0;
    dobra(&y);
    verificar_int("dobra(&y) com y=0", 0, y);

    int z = -3;
    dobra(&z);
    verificar_int("dobra(&z) com z=-3", -6, z);

    int w = 100;
    dobra(&w);
    dobra(&w);
    verificar_int("dobra(&w) duas vezes, w=100", 400, w);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
