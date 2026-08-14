#ifndef EXERCICIO_H
#define EXERCICIO_H

/* Um ponto no plano: dois inteiros agrupados sob um nome só. */
struct Ponto {
    int x;
    int y;
};

/* Imprime os exemplos comentados. O harness de teste nao chama esta funcao:
   a saida dos exemplos ficaria misturada com o resultado das verificacoes. */
void exemplos(void);

/* SUA VEZ: um Ponto com os campos x e y preenchidos com os valores dados. */
struct Ponto criar_ponto(int x, int y);

#endif
