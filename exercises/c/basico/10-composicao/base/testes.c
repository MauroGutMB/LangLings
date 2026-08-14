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

    struct Retangulo r = {{2, 3}, 10, 5};

    struct Retangulo movido = deslocar(r, 4, -1);
    verificar_int("deslocar(r, 4, -1).origem.x", 6, movido.origem.x);
    verificar_int("deslocar(r, 4, -1).origem.y", 2, movido.origem.y);
    verificar_int("deslocar(r, 4, -1).largura", 10, movido.largura);
    verificar_int("deslocar(r, 4, -1).altura", 5, movido.altura);

    /* r não deve ser alterado: passou por valor. */
    verificar_int("r.origem.x continua 2", 2, r.origem.x);
    verificar_int("r.origem.y continua 3", 3, r.origem.y);

    struct Retangulo r2 = {{0, 0}, 1, 1};
    struct Retangulo parado = deslocar(r2, 0, 0);
    verificar_int("deslocar(r2, 0, 0).origem.x", 0, parado.origem.x);
    verificar_int("deslocar(r2, 0, 0).origem.y", 0, parado.origem.y);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
