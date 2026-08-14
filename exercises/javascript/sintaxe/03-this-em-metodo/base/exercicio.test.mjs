import { test } from 'node:test';
import assert from 'node:assert/strict';
import { contador } from './exercicio.mjs';

// A asserção que separa arrow function de método comum: chamar incrementar()
// não pode lançar. Numa arrow function declarada no nível do módulo, this é
// undefined, e this.valor explode antes de devolver qualquer coisa.
test('incrementar devolve o valor novo, sem lançar', () => {
  const antes = contador.valor;
  const depois = contador.incrementar();
  assert.equal(depois, antes + 1);
});

test('incrementar altera contador.valor de fato', () => {
  const antes = contador.valor;
  contador.incrementar();
  assert.equal(contador.valor, antes + 1);
});

test('chamadas repetidas acumulam', () => {
  const antes = contador.valor;
  contador.incrementar();
  contador.incrementar();
  contador.incrementar();
  assert.equal(contador.valor, antes + 3);
});
