#include "exercicio.h"

/* Parece razoável: sizeof(v) sobre sizeof(v[0]) dá o número de elementos, do
   mesmo jeito que funcionou lá em exemplos() de 05-arrays. A diferença é que
   ali sizeof via a DECLARAÇÃO do array; aqui v é um PARÂMETRO, e todo array
   usado como parâmetro decai para um ponteiro ao primeiro elemento. sizeof(v)
   mede o ponteiro, não o array — e o resultado não muda com o array chamado. */
int numero_de_elementos(const int v[], int n)
{
    (void)n;
    return (int)(sizeof(v) / sizeof(v[0]));
}
