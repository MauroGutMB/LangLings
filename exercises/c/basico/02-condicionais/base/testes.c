#include <stdio.h>

#include "exercicio.h"

static int falhas = 0;

static void verificar_char(const char *o_que, char esperado, char obtido)
{
    if (esperado == obtido) { printf("ok    %s\n", o_que); return; }
    printf("FALHA %s\n      esperado: %c\n      obtido:   %c\n", o_que, esperado, obtido);
    falhas++;
}

int main(void)
{
    /* Sem isto o stdout sai em bloco no fim do processo e as falhas apareceriam
       DEPOIS da mensagem de erro do make, que vem por stderr. */
    setvbuf(stdout, NULL, _IOLBF, 0);

    verificar_char("conceito(100)", 'A', conceito(100));
    verificar_char("conceito(90)", 'A', conceito(90));
    verificar_char("conceito(89)", 'B', conceito(89));
    verificar_char("conceito(80)", 'B', conceito(80));
    verificar_char("conceito(79)", 'C', conceito(79));
    verificar_char("conceito(70)", 'C', conceito(70));
    verificar_char("conceito(69)", 'F', conceito(69));
    verificar_char("conceito(0)", 'F', conceito(0));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
