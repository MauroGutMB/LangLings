import { test } from 'node:test';
import assert from 'node:assert/strict';
import { temNaN } from './exercicio.mjs';

test('array sem NaN devolve false', () => {
  assert.equal(temNaN([1, 2, 3]), false);
});

test('array vazio devolve false', () => {
  assert.equal(temNaN([]), false);
});

// A asserção que separa indexOf de Number.isNaN: NaN !== NaN faz
// arr.indexOf(NaN) nunca encontrar nada, mesmo com o valor lá dentro.
test('array com NaN devolve true', () => {
  assert.equal(temNaN([1, NaN, 3]), true);
});

test('NaN no início ou no fim também conta', () => {
  assert.equal(temNaN([NaN, 1, 2]), true);
  assert.equal(temNaN([1, 2, NaN]), true);
});
