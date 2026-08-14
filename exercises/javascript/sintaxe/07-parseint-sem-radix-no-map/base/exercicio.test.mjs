import { test } from 'node:test';
import assert from 'node:assert/strict';
import { paraNumeros } from './exercicio.mjs';

test('array de um elemento só converte certo', () => {
  assert.deepEqual(paraNumeros(['42']), [42]);
});

test('array vazio devolve array vazio', () => {
  assert.deepEqual(paraNumeros([]), []);
});

// A asserção que separa map(parseInt) de map((s) => parseInt(s, 10)): o
// índice de cada posição vira a base da conversão quando o callback é
// passado direto.
test('array com vários elementos iguais converte todos igual', () => {
  assert.deepEqual(paraNumeros(['10', '10', '10', '10']), [10, 10, 10, 10]);
});

test('preserva a ordem dos valores convertidos', () => {
  assert.deepEqual(paraNumeros(['5', '20', '7']), [5, 20, 7]);
});
