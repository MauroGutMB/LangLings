import test from 'node:test';
import assert from 'node:assert/strict';
import { somarSeNumero } from './exercicio.mts';

test('valor é number', () => {
  assert.equal(somarSeNumero(5, 10), 15);
});

test('valor não é number', () => {
  assert.equal(somarSeNumero('cinco', 10), 10);
});

test('valor é null', () => {
  assert.equal(somarSeNumero(null, 3), 3);
});
