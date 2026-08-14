import { test } from 'node:test';
import assert from 'node:assert/strict';
import { ordenarNumeros } from './exercicio.mjs';

test('ordena números de um dígito só', () => {
  assert.deepEqual(ordenarNumeros([3, 1, 2]), [1, 2, 3]);
});

test('não altera o array recebido', () => {
  const original = [3, 1, 2];
  ordenarNumeros(original);
  assert.deepEqual(original, [3, 1, 2]);
});

// A asserção que separa ordem de string de ordem numérica: comparado como
// texto, '10' vem antes de '2' porque '1' < '2' no primeiro caractere.
test('ordena números com quantidades de dígitos diferentes', () => {
  assert.deepEqual(ordenarNumeros([10, 2, 1]), [1, 2, 10]);
});
