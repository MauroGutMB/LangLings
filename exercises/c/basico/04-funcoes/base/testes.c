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

    /* As tres posicoes do maior sao testadas de proposito: um retorno fixo
       como `return a` passaria na primeira e reprovaria nas outras duas. */
    verificar_int("maior_de_tres(9, 1, 2)", 9, maior_de_tres(9, 1, 2));
    verificar_int("maior_de_tres(1, 9, 2)", 9, maior_de_tres(1, 9, 2));
    verificar_int("maior_de_tres(1, 2, 9)", 9, maior_de_tres(1, 2, 9));
    verificar_int("maior_de_tres(5, 5, 5)", 5, maior_de_tres(5, 5, 5));
    verificar_int("maior_de_tres(-3, -1, -7)", -1, maior_de_tres(-3, -1, -7));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
