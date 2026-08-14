import test from 'node:test';
import assert from 'node:assert/strict';
import { somarPares } from './exercicio.mjs';

test('soma só os pares', () => {
  assert.equal(somarPares([1, 2, 3, 4]), 6);
});

test('ignora um array inteiro de ímpares', () => {
  assert.equal(somarPares([1, 3, 5]), 0);
});

test('o zero é par e não muda a soma', () => {
  assert.equal(somarPares([0, 5]), 0);
  assert.equal(somarPares([0, 2]), 2);
});

test('números negativos também são pares ou ímpares', () => {
  assert.equal(somarPares([-4, -3, 2]), -2);
});

test('array vazio soma 0', () => {
  assert.equal(somarPares([]), 0);
});
