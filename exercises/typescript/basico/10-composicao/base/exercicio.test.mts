import test from 'node:test';
import assert from 'node:assert/strict';
import { calcularTotal } from './exercicio.mts';

const produtos = [
  { nome: 'Caneta', preco: 2 },
  { nome: 'Caderno', preco: 10 },
];

test('sem cupom, soma direta', () => {
  assert.equal(calcularTotal(produtos), 12);
});

test('cupom percentual', () => {
  assert.equal(calcularTotal(produtos, { tipo: 'percentual', valor: 50 }), 6);
});

test('cupom fixo', () => {
  assert.equal(calcularTotal(produtos, { tipo: 'fixo', valor: 5 }), 7);
});

test('cupom fixo maior que o total não deixa negativo', () => {
  assert.equal(calcularTotal(produtos, { tipo: 'fixo', valor: 100 }), 0);
});

test('lista de produtos vazia', () => {
  assert.equal(calcularTotal([]), 0);
});
