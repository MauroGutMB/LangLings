import { test } from 'node:test';
import assert from 'node:assert/strict';
import { ehObjeto } from './exercicio.mjs';

test('um objeto literal é objeto', () => {
  assert.equal(ehObjeto({}), true);
});

test('um array é objeto', () => {
  assert.equal(ehObjeto([1, 2, 3]), true);
});

test('número, string e booleano não são objeto', () => {
  assert.equal(ehObjeto(5), false);
  assert.equal(ehObjeto('texto'), false);
  assert.equal(ehObjeto(true), false);
});

// A asserção que separa typeof puro de typeof + checagem de null: typeof null
// devolve 'object', mas null não é um objeto de verdade.
test('null não é objeto', () => {
  assert.equal(ehObjeto(null), false);
});
