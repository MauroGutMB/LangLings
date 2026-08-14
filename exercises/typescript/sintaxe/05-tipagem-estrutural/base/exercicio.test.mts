import test from 'node:test';
import assert from 'node:assert/strict';
import { velocidadeMedia } from './exercicio.mts';

test('100 metros em 20 segundos é 5 m/s', () => {
  assert.equal(velocidadeMedia(100, 20), 5);
});

test('10 metros em 2 segundos é 5 m/s', () => {
  assert.equal(velocidadeMedia(10, 2), 5);
});

test('parado no tempo (distância 0) dá 0 m/s', () => {
  assert.equal(velocidadeMedia(0, 10), 0);
});
