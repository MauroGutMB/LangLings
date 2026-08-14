import { test } from 'node:test';
import assert from 'node:assert/strict';
import { comPadrao } from './exercicio.mjs';

test('valor ausente (undefined) usa o padrão', () => {
  assert.equal(comPadrao(undefined, 10), 10);
});

test('valor null usa o padrão', () => {
  assert.equal(comPadrao(null, 10), 10);
});

test('valor informado e verdadeiro vence o padrão', () => {
  assert.equal(comPadrao(5, 10), 5);
});

// A asserção que separa || de ??: 0 é um valor informado, não a ausência de
// um. || não enxerga essa diferença, porque 0 também é falsy.
test('o número 0 conta como informado', () => {
  assert.equal(comPadrao(0, 10), 0);
});

test('a string vazia conta como informada', () => {
  assert.equal(comPadrao('', 'padrao'), '');
});
