import { test } from 'node:test';
import assert from 'node:assert/strict';
import { ehZero } from './exercicio.mjs';

test('o número 0 é zero', () => {
  assert.equal(ehZero(0), true);
});

test('o número -0 também é zero', () => {
  assert.equal(ehZero(-0), true);
});

test('outros números não são zero', () => {
  assert.equal(ehZero(5), false);
  assert.equal(ehZero(-1), false);
});

// A asserção que separa == de ===: nenhum destes é o número 0, mas == os
// converte para 0 antes de comparar e a versão ingênua diz que são.
test('string vazia, false e null não são zero', () => {
  assert.equal(ehZero(''), false);
  assert.equal(ehZero(false), false);
  assert.equal(ehZero(null), false);
});
