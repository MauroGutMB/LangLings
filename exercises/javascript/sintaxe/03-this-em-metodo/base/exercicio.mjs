// contador.incrementar() deve somar 1 a contador.valor e devolver o valor novo.
//
// TODO: o método está escrito como arrow function. Uma arrow function não tem
// this próprio — ela usa o this de onde foi ESCRITA, e aqui isso é o nível do
// módulo, onde this é undefined.
export const contador = {
  valor: 0,
  incrementar: () => {
    this.valor += 1;
    return this.valor;
  },
};
