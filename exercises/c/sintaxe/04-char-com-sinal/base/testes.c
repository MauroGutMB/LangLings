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

    verificar_int("eh_byte_alto(0)", 0, eh_byte_alto((char)0));
    verificar_int("eh_byte_alto(127)", 0, eh_byte_alto((char)127));
    /* A asserção que separa a solução ingênua da correta: 0xFF é o byte 255,
       bem acima de 128 — mas como char com sinal ele vale -1. */
    verificar_int("eh_byte_alto(0xFF)", 1, eh_byte_alto((char)0xFF));
    verificar_int("eh_byte_alto(0x80)", 1, eh_byte_alto((char)0x80)); /* byte 128 */
    verificar_int("eh_byte_alto(0xC8)", 1, eh_byte_alto((char)0xC8)); /* byte 200 */
    verificar_int("eh_byte_alto(1)", 0, eh_byte_alto((char)1));

    if (falhas > 0) { printf("\n%d verificacao(oes) falharam\n", falhas); return 1; }
    printf("\ntodas as verificacoes passaram\n");
    return 0;
}
