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

    verificar_int("comeca_com(\"banana\", \"ban\")", 1, comeca_com("banana", "ban"));
    verificar_int("comeca_com(\"banana\", \"ana\")", 0, comeca_com("banana", "ana"));
    verificar_int("comeca_com(\"banana\", \"\")", 1, comeca_com("banana", ""));
    verificar_int("comeca_com(\"\", \"\")", 1, comeca_com("", ""));
    verificar_int("comeca_com(\"ba\", \"banana\")", 0, comeca_com("ba", "banana"));
    verificar_int("comeca_com(\"banana\", \"banana\")", 1, comeca_com("banana", "banana"));
    verificar_int("comeca_com(\"banana\", \"banX\")", 0, comeca_com("banana", "banX"));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
