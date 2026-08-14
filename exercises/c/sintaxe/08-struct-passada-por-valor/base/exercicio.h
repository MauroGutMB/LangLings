#ifndef EXERCICIO_H
#define EXERCICIO_H

struct Ponto {
    int x;
    int y;
};

/* Soma dx a p->x e dy a p->y, alterando o Ponto apontado por p. */
void mover(struct Ponto *p, int dx, int dy);

#endif
