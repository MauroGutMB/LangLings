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

    struct Ponto p = {2, 3};
    mover(&p, 4, -1);
    /* A asserção que separa a solução ingênua da correta: mover para uma
       cópia local nunca altera p.x e p.y de quem chamou. */
    verificar_int("depois de mover(&p, 4, -1), p.x", 6, p.x);
    verificar_int("depois de mover(&p, 4, -1), p.y", 2, p.y);

    struct Ponto origem = {0, 0};
    mover(&origem, 0, 0);
    verificar_int("mover(&origem, 0, 0), origem.x", 0, origem.x);
    verificar_int("mover(&origem, 0, 0), origem.y", 0, origem.y);

    struct Ponto q = {-5, 10};
    mover(&q, -3, -3);
    mover(&q, -3, -3);
    verificar_int("mover(&q, ...) duas vezes, q.x", -11, q.x);
    verificar_int("mover(&q, ...) duas vezes, q.y", 4, q.y);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
