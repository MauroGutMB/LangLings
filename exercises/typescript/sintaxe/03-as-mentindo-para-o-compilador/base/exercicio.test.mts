import test from 'node:test';
import assert from 'node:assert/strict';
import { comprimento } from './exercicio.mts';

test('string normal', () => {
  assert.equal(comprimento('abc'), 3);
});

test('string vazia', () => {
  assert.equal(comprimento(''), 0);
});

test('valor que não é string devolve 0', () => {
  assert.equal(comprimento(42), 0);
});

test('null devolve 0', () => {
  assert.equal(comprimento(null), 0);
});
