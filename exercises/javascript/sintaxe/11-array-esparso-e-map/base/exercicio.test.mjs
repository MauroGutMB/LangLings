import { test } from 'node:test';
import assert from 'node:assert/strict';
import { criarZeros } from './exercicio.mjs';

test('o comprimento do array é n', () => {
  assert.equal(criarZeros(4).length, 4);
});

test('array de tamanho 0 devolve array vazio', () => {
  assert.deepEqual(criarZeros(0), []);
});

// A asserção que separa Array(n).map(...) de Array.from({length:n}, ...): num
// array esparso não existe nenhuma posição de verdade para map visitar, então
// o resultado não tem 0 nenhum lá dentro, mesmo tendo o length certo.
test('cada posição vale 0 de verdade', () => {
  assert.deepEqual(criarZeros(3), [0, 0, 0]);
});
