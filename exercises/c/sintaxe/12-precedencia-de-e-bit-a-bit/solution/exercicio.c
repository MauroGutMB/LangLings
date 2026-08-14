#include "exercicio.h"

/* Os parênteses forçam o & a ser calculado antes do ==, na ordem que o
 * problema pede: primeiro isola os bits de mascara presentes em flags, depois
 * compara o resultado com mascara inteira.
 */
int tem_flag(int flags, int mascara)
{
    return (flags & mascara) == mascara;
}
