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

    /* O buffer é preenchido com um byte sentinela diferente de zero antes de
       cada chamada. Se copia_string não escrever o terminador na posição
       certa, o byte logo depois do conteúdo copiado continua sendo o
       sentinela — e strcmp encontra a diferença aí, sem nunca ler além do
       buffer alocado (32 bytes é folga suficiente para todos os casos). */
    char buf1[32];
    memset(buf1, 'X', sizeof(buf1));
    copia_string(buf1, "ola");
    verificar_string("copia_string(_, \"ola\")", "ola", buf1);

    char buf2[32];
    memset(buf2, 'X', sizeof(buf2));
    copia_string(buf2, "");
    verificar_string("copia_string(_, \"\")", "", buf2);

    char buf3[32];
    memset(buf3, 'X', sizeof(buf3));
    copia_string(buf3, "abcdefgh");
    verificar_string("copia_string(_, \"abcdefgh\")", "abcdefgh", buf3);

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
