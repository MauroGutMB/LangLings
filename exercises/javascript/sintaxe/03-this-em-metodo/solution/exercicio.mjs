// contador.incrementar() deve somar 1 a contador.valor e devolver o valor novo.
export const contador = {
  valor: 0,
  incrementar() {
    this.valor += 1;
    return this.valor;
  },
};
