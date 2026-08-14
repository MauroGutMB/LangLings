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

    verificar_int("eh_igual_a_cinco(5)", 1, eh_igual_a_cinco(5));
    /* A asserção que separa a solução ingênua da correta: um x diferente de
       5 precisa devolver 0, e a atribuição sempre entraria no if. */
    verificar_int("eh_igual_a_cinco(0)", 0, eh_igual_a_cinco(0));
    verificar_int("eh_igual_a_cinco(10)", 0, eh_igual_a_cinco(10));
    verificar_int("eh_igual_a_cinco(-5)", 0, eh_igual_a_cinco(-5));
    verificar_int("eh_igual_a_cinco(4)", 0, eh_igual_a_cinco(4));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
