import test from 'node:test';
import assert from 'node:assert/strict';
import { paraNumeroSeguro } from './exercicio.mts';

test('já é number', () => {
  assert.equal(paraNumeroSeguro(42), 42);
});

test('é uma string', () => {
  assert.equal(paraNumeroSeguro('42'), 0);
});

test('é null', () => {
  assert.equal(paraNumeroSeguro(null), 0);
});

test('é undefined', () => {
  assert.equal(paraNumeroSeguro(undefined), 0);
});

test('zero continua sendo number', () => {
  assert.equal(paraNumeroSeguro(0), 0);
});
