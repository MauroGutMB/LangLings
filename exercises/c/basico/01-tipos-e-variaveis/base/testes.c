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

    /* O esperado sai do proprio sizeof: o exercicio e perguntar ao compilador,
       nao decorar que um double tem 8 bytes nesta maquina. */
    verificar_int("bytes_de_doubles(1)", (long)sizeof(double), bytes_de_doubles(1));
    verificar_int("bytes_de_doubles(3)", 3 * (long)sizeof(double), bytes_de_doubles(3));
    verificar_int("bytes_de_doubles(0)", 0, bytes_de_doubles(0));
    verificar_int("bytes_de_doubles(1000)", 1000 * (long)sizeof(double), bytes_de_doubles(1000));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
