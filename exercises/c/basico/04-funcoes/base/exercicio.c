#include <stdio.h>

#include "exercicio.h"

/* O compilador lê o arquivo de cima para baixo e precisa conhecer a assinatura
   de uma função ANTES da primeira chamada. Este é o protótipo de dobro: ele
   promete que a definição aparece mais adiante. As funções que outros arquivos
   chamam têm o protótipo em exercicio.h; as de uso interno, aqui mesmo. */
static int dobro(int x);

/* maior devolve o maior entre a e b. Sem static, porque exercicio.h a declara
   e testes.c pode chamá-la. */
int maior(int a, int b)
{
    return a > b ? a : b;
}

void exemplos(void)
{
    printf("%d\n", maior(3, 9)); /* 9 */
    printf("%d\n", dobro(21));   /* 42 — a definição vem depois, o protótipo
                                    lá em cima é o que permite a chamada */

    /* Todo parâmetro em C chega por CÓPIA. O que a função faz com o parâmetro
       não sai dela: quem chamou continua com o valor original. */
    int n = 10;
    printf("%d %d\n", dobro(n), n); /* 20 10 */

    /* Uma função pode chamar outra, inclusive aninhando os resultados. É a
       forma mais barata de reduzir um problema a um caso já resolvido. */
    printf("%d\n", maior(maior(1, 7), 4)); /* 7 */

    /* void na lista de parâmetros significa "nenhum parâmetro". Deixar os
       parênteses vazios em C não é a mesma coisa: seria uma declaração antiga,
       que aceita qualquer lista de argumentos sem conferir nada. */
}

/* static limita a função a este arquivo: nenhum outro .c consegue chamá-la,
   e o compilador avisa se ela nunca for usada. */
static int dobro(int x)
{
    return x * 2;
}

/* SUA VEZ
 *
 * Devolva o maior entre a, b e c.
 */
int maior_de_tres(int a, int b, int c)
{
    return 0; /* <- troque isto */
}
