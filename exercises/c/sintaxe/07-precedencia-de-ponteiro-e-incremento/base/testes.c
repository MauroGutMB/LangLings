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

    int x = 5;
    incrementa_valor(&x);
    /* A asserção que separa a solução ingênua da correta: *p++ nunca escreve
       em *p, então x continuaria 5. */
    verificar_int("incrementa_valor(&x) com x=5", 6, x);

    int y = -1;
    incrementa_valor(&y);
    verificar_int("incrementa_valor(&y) com y=-1", 0, y);

    int z = 0;
    incrementa_valor(&z);
    incrementa_valor(&z);
    incrementa_valor(&z);
    verificar_int("incrementa_valor(&z) tres vezes, z=0", 3, z);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
