#ifndef EXERCICIO_H
#define EXERCICIO_H

/* Um ponto no plano. */
struct Ponto {
    int x;
    int y;
};

/* Um retângulo é descrito pelo ponto do seu canto superior esquerdo, mais
   largura e altura — composição: um struct guardando outro struct. */
struct Retangulo {
    struct Ponto origem;
    int largura;
    int altura;
};

/* Imprime os exemplos comentados. O harness de teste nao chama esta funcao:
   a saida dos exemplos ficaria misturada com o resultado das verificacoes. */
void exemplos(void);

/* SUA VEZ: um Retangulo igual a r, com origem deslocada em (dx, dy).
   largura e altura nao mudam. */
struct Retangulo deslocar(struct Retangulo r, int dx, int dy);

#endif
