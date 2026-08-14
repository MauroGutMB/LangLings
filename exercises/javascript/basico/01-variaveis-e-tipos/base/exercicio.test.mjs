import test from 'node:test';
import assert from 'node:assert/strict';
import { temValor } from './exercicio.mjs';

test('reconhece valores comuns', () => {
  assert.equal(temValor(42), true);
  assert.equal(temValor('oi'), true);
  assert.equal(temValor(true), true);
  assert.equal(temValor({}), true);
});

test('undefined e null são ausência de valor', () => {
  assert.equal(temValor(undefined), false);
  assert.equal(temValor(null), false);
});

test('uma propriedade que não existe vale undefined', () => {
  const config = { tema: 'escuro' };
  assert.equal(temValor(config.tema), true);
  assert.equal(temValor(config.fonte), false);
});

test('0, string vazia e false são valores', () => {
  assert.equal(temValor(0), true);
  assert.equal(temValor(''), true);
  assert.equal(temValor(false), true);
});
