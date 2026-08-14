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

    int a[3] = {1, 2, 3};
    int b[5] = {1, 2, 3, 4, 5};
    int c[1] = {9};
    int d[4] = {1, 2, 3, 4};

    verificar_int("numero_de_elementos(a, 3)", 3, numero_de_elementos(a, 3));
    /* A asserção que separa a solução ingênua da correta: sizeof(v)/sizeof(v[0])
       dentro da função sempre dá 2 (16 bytes de ponteiro / 4 bytes de int, ou
       equivalente), então qualquer array com n diferente de 2 expõe o bug. */
    verificar_int("numero_de_elementos(b, 5)", 5, numero_de_elementos(b, 5));
    verificar_int("numero_de_elementos(c, 1)", 1, numero_de_elementos(c, 1));
    verificar_int("numero_de_elementos(d, 4)", 4, numero_de_elementos(d, 4));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
