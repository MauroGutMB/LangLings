import test from 'node:test';
import assert from 'node:assert/strict';
import { primeiroEUltimo } from './exercicio.mts';

test('array com vários elementos', () => {
  assert.deepEqual(primeiroEUltimo([1, 2, 3, 4]), [1, 4]);
});

test('array com um único elemento repete o valor', () => {
  assert.deepEqual(primeiroEUltimo([7]), [7, 7]);
});

test('array com dois elementos', () => {
  assert.deepEqual(primeiroEUltimo([5, 9]), [5, 9]);
});

test('a ordem do retorno é [primeiro, ultimo], não ordenado', () => {
  assert.deepEqual(primeiroEUltimo([30, 10, 20]), [30, 20]);
});
