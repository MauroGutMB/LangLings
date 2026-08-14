#include <stdio.h>
#include <string.h>

#include "exercicio.h"

static int falhas = 0;

static void verificar_string(const char *o_que, const char *esperado, const char *obtido)
{
    if (strcmp(esperado, obtido) == 0) { printf("ok    %s\n", o_que); return; }
    printf("FALHA %s\n      esperado: %s\n      obtido:   %s\n", o_que, esperado, obtido);
    falhas++;
}

int main(void)
{
    setvbuf(stdout, NULL, _IOLBF, 0);

    char *a = para_maiusculas("ola");
    verificar_string("para_maiusculas(\"ola\")", "OLA", a);

    /* A asserção que separa a solução ingênua da correta: obter um segundo
       resultado não pode apagar o primeiro que ainda está guardado em a. */
    char *b = para_maiusculas("mundo");
    verificar_string("para_maiusculas(\"mundo\")", "MUNDO", b);
    verificar_string("a ainda e \"OLA\" depois da segunda chamada", "OLA", a);

    char *c = para_maiusculas("");
    verificar_string("para_maiusculas(\"\")", "", c);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
