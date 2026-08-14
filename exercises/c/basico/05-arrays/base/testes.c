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

    int crescente[4] = {1, 2, 3, 9};
    int decrescente[4] = {9, 3, 2, 1};
    int meio[5] = {2, 4, 8, 4, 2};
    int um_so[1] = {42};

    /* Um array so de negativos separa quem percorreu de quem inicializou o
       candidato com 0 em vez de com o primeiro elemento. */
    int negativos[3] = {-8, -2, -5};

    verificar_int("maior_valor(crescente, 4)", 9, maior_valor(crescente, 4));
    verificar_int("maior_valor(decrescente, 4)", 9, maior_valor(decrescente, 4));
    verificar_int("maior_valor(meio, 5)", 8, maior_valor(meio, 5));
    verificar_int("maior_valor(um_so, 1)", 42, maior_valor(um_so, 1));
    verificar_int("maior_valor(negativos, 3)", -2, maior_valor(negativos, 3));

    /* n menor que o array: so os 2 primeiros contam, o 9 na posicao 3 nao. */
    verificar_int("maior_valor(crescente, 2)", 2, maior_valor(crescente, 2));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
