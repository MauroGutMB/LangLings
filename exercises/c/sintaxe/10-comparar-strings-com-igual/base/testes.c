#include <stdio.h>
#include <string.h>

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

    /* buf guarda uma cópia de "banana" num endereço diferente do literal:
       mesmo conteúdo, ponteiros diferentes. */
    char buf[16];
    strcpy(buf, "banana");

    /* A asserção que separa a solução ingênua da correta: mesmo conteúdo,
       endereços diferentes — == compararia ponteiros e diria que são
       diferentes. */
    verificar_int("strings_iguais(\"banana\", buf)", 1, strings_iguais("banana", buf));

    verificar_int("strings_iguais(\"banana\", \"maca\")", 0, strings_iguais("banana", "maca"));
    verificar_int("strings_iguais(\"\", \"\")", 1, strings_iguais("", ""));
    verificar_int("strings_iguais(\"ba\", \"banana\")", 0, strings_iguais("ba", "banana"));

    /* Mesmo ponteiro nos dois lados também precisa dar igual. */
    const char *p = "abc";
    verificar_int("strings_iguais(p, p)", 1, strings_iguais(p, p));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
