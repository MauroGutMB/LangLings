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

    verificar_int("divisao_para_baixo(7, 2)", 3, divisao_para_baixo(7, 2));
    verificar_int("divisao_para_baixo(6, 3)", 2, divisao_para_baixo(6, 3));
    /* A asserção que separa a solução ingênua da correta: -7/2 trunca para
       -3 em C, mas arredondado para baixo o resultado é -4. */
    verificar_int("divisao_para_baixo(-7, 2)", -4, divisao_para_baixo(-7, 2));
    verificar_int("divisao_para_baixo(7, -2)", -4, divisao_para_baixo(7, -2));
    verificar_int("divisao_para_baixo(-7, -2)", 3, divisao_para_baixo(-7, -2));
    verificar_int("divisao_para_baixo(-6, 3)", -2, divisao_para_baixo(-6, 3));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
