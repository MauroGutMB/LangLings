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

    verificar_int("soma_ate(1)", 1, soma_ate(1));
    verificar_int("soma_ate(4)", 10, soma_ate(4));
    verificar_int("soma_ate(100)", 5050, soma_ate(100));
    verificar_int("soma_ate(0)", 0, soma_ate(0));
    verificar_int("soma_ate(-5)", 0, soma_ate(-5));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
