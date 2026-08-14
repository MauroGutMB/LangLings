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

    struct Ponto p1 = criar_ponto(3, 4);
    verificar_int("criar_ponto(3, 4).x", 3, p1.x);
    verificar_int("criar_ponto(3, 4).y", 4, p1.y);

    struct Ponto p2 = criar_ponto(-1, 0);
    verificar_int("criar_ponto(-1, 0).x", -1, p2.x);
    verificar_int("criar_ponto(-1, 0).y", 0, p2.y);

    struct Ponto p3 = criar_ponto(0, 0);
    verificar_int("criar_ponto(0, 0).x", 0, p3.x);
    verificar_int("criar_ponto(0, 0).y", 0, p3.y);

    /* criar_ponto duas vezes precisa devolver structs independentes. */
    verificar_int("p1.x continua 3", 3, p1.x);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
