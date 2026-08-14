import test from 'node:test';
import assert from 'node:assert/strict';
import { classificar } from './exercicio.mjs';

test('classifica notas no meio de cada faixa', () => {
  assert.equal(classificar(95), 'A');
  assert.equal(classificar(85), 'B');
  assert.equal(classificar(75), 'C');
  assert.equal(classificar(42), 'F');
});

test('os limites de cada faixa são inclusivos', () => {
  assert.equal(classificar(90), 'A');
  assert.equal(classificar(80), 'B');
  assert.equal(classificar(70), 'C');
});

test('um ponto abaixo do limite cai na faixa de baixo', () => {
  assert.equal(classificar(89), 'B');
  assert.equal(classificar(79), 'C');
  assert.equal(classificar(69), 'F');
});

test('os extremos da escala', () => {
  assert.equal(classificar(100), 'A');
  assert.equal(classificar(0), 'F');
});
