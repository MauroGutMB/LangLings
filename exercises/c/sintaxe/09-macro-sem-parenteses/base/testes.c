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

    /* usa_dobro(a, b, fator) = DOBRO(a + b) * fator = (2*(a+b)) * fator */
    verificar_int("usa_dobro(3, 4, 1)", 14, usa_dobro(3, 4, 1));
    /* A asserção que separa a solução ingênua da correta: com fator diferente
       de 1, a expansão sem parênteses espalha o * fator só sobre o último
       termo, em vez de sobre o dobro inteiro. */
    verificar_int("usa_dobro(3, 4, 2)", 28, usa_dobro(3, 4, 2));
    verificar_int("usa_dobro(1, 1, 5)", 20, usa_dobro(1, 1, 5));
    verificar_int("usa_dobro(0, 0, 3)", 0, usa_dobro(0, 0, 3));
    verificar_int("usa_dobro(-2, 5, 2)", 12, usa_dobro(-2, 5, 2));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
