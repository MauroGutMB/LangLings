import test from 'node:test';
import assert from 'node:assert/strict';
import { maiorPorTamanho } from './exercicio.mts';

test('strings, a mais longa vence', () => {
  assert.equal(maiorPorTamanho('oi', 'alo'), 'alo');
});

test('arrays, o mais longo vence', () => {
  assert.deepEqual(maiorPorTamanho([1, 2], [1, 2, 3]), [1, 2, 3]);
});

test('mesmo tamanho, devolve o segundo', () => {
  assert.equal(maiorPorTamanho('abc', 'xyz'), 'xyz');
});
