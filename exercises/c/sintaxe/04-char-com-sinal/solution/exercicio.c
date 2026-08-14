#include "exercicio.h"

/* (unsigned char)c reinterpreta o mesmo byte na faixa 0..255, sem trocar o
 * conteúdo — só a leitura do sinal.
 */
int eh_byte_alto(char c)
{
    return (unsigned char)c >= 128;
}
