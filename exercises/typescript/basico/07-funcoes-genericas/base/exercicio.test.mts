import test from 'node:test';
import assert from 'node:assert/strict';
import { ultimo } from './exercicio.mts';

test('array de números', () => {
  assert.equal(ultimo([1, 2, 3]), 3);
});

test('array de strings', () => {
  assert.equal(ultimo(['a', 'b', 'c']), 'c');
});

test('array vazio devolve undefined', () => {
  assert.equal(ultimo([]), undefined);
});

test('array com um único elemento', () => {
  assert.equal(ultimo([7]), 7);
});
