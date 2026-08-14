import test from 'node:test';
import assert from 'node:assert/strict';
import { media } from './exercicio.mjs';

test('média de valores inteiros', () => {
  assert.equal(media([8, 9, 10]), 9);
  assert.equal(media([5]), 5);
});

test('a média não precisa ser inteira', () => {
  assert.equal(media([1, 2]), 1.5);
});

test('média com negativos', () => {
  assert.equal(media([-2, 2]), 0);
});

// 0 / 0 é NaN, e NaN não é uma média.
test('array vazio devolve 0', () => {
  assert.equal(media([]), 0);
});

test('não altera o array recebido', () => {
  const notas = [1, 2, 3];
  media(notas);
  assert.deepEqual(notas, [1, 2, 3]);
});
