import test from 'node:test';
import assert from 'node:assert/strict';
import { iniciais } from './exercicio.mjs';

test('uma letra por palavra, em maiúsculas', () => {
  assert.equal(iniciais('Ana Maria Silva'), 'AMS');
  assert.equal(iniciais('ana maria silva'), 'AMS');
});

test('um nome só devolve uma letra', () => {
  assert.equal(iniciais('Ana'), 'A');
});

test('espaço sobrando nas pontas não vira inicial', () => {
  assert.equal(iniciais('   Ana Silva  '), 'AS');
});

test('as iniciais saem sem separador', () => {
  assert.equal(iniciais('Joao Pedro de Souza'), 'JPDS');
});
