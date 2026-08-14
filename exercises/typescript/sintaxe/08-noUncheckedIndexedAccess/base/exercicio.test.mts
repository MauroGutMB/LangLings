import test from 'node:test';
import assert from 'node:assert/strict';
import { primeiraLetra } from './exercicio.mts';

test('índice válido', () => {
  assert.equal(primeiraLetra(['abacaxi', 'banana'], 1), 'B');
});

test('índice fora dos limites devolve string vazia', () => {
  assert.equal(primeiraLetra(['abacaxi'], 5), '');
});

test('array vazio devolve string vazia', () => {
  assert.equal(primeiraLetra([], 0), '');
});
