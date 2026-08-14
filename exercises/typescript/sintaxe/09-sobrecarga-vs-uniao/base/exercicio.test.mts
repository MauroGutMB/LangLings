import test from 'node:test';
import assert from 'node:assert/strict';
import { converterTipo } from './exercicio.mts';

test('string vira number', () => {
  const resultado = converterTipo('42');
  assert.equal(resultado.toFixed(0), '42');
});

test('number vira string', () => {
  const resultado = converterTipo(7);
  assert.equal(resultado.toUpperCase(), '7');
});
