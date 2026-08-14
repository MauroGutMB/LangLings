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

    /* bits: 0b0001=1  0b0010=2  0b0100=4  0b1000=8 */
    verificar_int("tem_flag(0b0011, 0b0001)", 1, tem_flag(3, 1));
    /* A asserção que separa a solução ingênua da correta: mascara com mais de
       um bit ligado. flags & 1 (a expressão que a versão ingênua calcula de
       verdade) dá 1 aqui, mas flags=0b0101 não tem o bit 0b0010 de mascara
       ligado — a resposta certa é 0. */
    verificar_int("tem_flag(0b0101, 0b0011)", 0, tem_flag(5, 3));
    verificar_int("tem_flag(0b0111, 0b0011)", 1, tem_flag(7, 3));
    verificar_int("tem_flag(0b1000, 0b0001)", 0, tem_flag(8, 1));
    verificar_int("tem_flag(0b0000, 0b0000)", 1, tem_flag(0, 0));
    verificar_int("tem_flag(0b1111, 0b1010)", 1, tem_flag(15, 10));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
