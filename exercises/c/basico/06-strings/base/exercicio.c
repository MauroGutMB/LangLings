#include <stdio.h>
#include <string.h>

#include "exercicio.h"

/* C não tem tipo string. O que existe é um array de char cujo fim é marcado
   pelo byte '\0'. Toda função de <string.h> procura esse byte para saber onde
   a string acaba — sem ele, ela continua lendo memória adiante. */
void exemplos(void)
{
    char nome[] = "Ana"; /* 4 bytes na memória: 'A' 'n' 'a' '\0' */

    printf("%s\n", nome);         /* Ana */
    printf("%zu\n", strlen(nome)); /* 3 — strlen conta até o '\0', sem contá-lo */
    printf("%zu\n", sizeof(nome)); /* 4 — o array inclui o terminador */

    printf("%c %d\n", nome[0], nome[3] == '\0'); /* A 1 */

    /* strcpy copia a origem inteira, terminador incluído. O destino precisa
       ter espaço para os caracteres MAIS o '\0'. */
    char destino[16];
    strcpy(destino, "Bruno");
    printf("%s %zu\n", destino, strlen(destino)); /* Bruno 5 */

    /* strcat concatena a partir do '\0' do destino, sobrescrevendo-o. */
    strcat(destino, " Silva");
    printf("%s\n", destino); /* Bruno Silva */

    /* strcmp devolve 0 quando as strings têm o mesmo conteúdo, e um valor
       negativo ou positivo conforme a ordem. Repare que 0 significa IGUAL:
       usar o resultado direto num if inverte o sentido. */
    printf("%d\n", strcmp("abc", "abc") == 0); /* 1 */
    printf("%d\n", strcmp("abc", "abd") < 0);  /* 1 */

    /* strncmp limita a comparação aos n primeiros caracteres. É o que permite
       perguntar sobre o começo de uma string sem olhar o resto dela. */
    printf("%d\n", strncmp("prefixado", "pre", 3) == 0); /* 1 */

    /* Percorrer caractere a caractere é o laço mais comum em C: a condição é
       o próprio terminador. */
    int vogais = 0;
    for (int i = 0; nome[i] != '\0'; i++) {
        if (strchr("aeiouAEIOU", nome[i]) != NULL) {
            vogais++;
        }
    }
    printf("%d\n", vogais); /* 2 */
}

/* SUA VEZ
 *
 * Devolva 1 se s começa com prefixo, e 0 se não.
 * O prefixo vazio começa qualquer string.
 */
int comeca_com(const char *s, const char *prefixo)
{
    return 0; /* <- troque isto */
}
