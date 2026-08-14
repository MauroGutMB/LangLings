import test from 'node:test';
import assert from 'node:assert/strict';
import { descreverStatus } from './exercicio.mts';

test('ativo', () => {
  assert.equal(descreverStatus('ativo'), 'em andamento');
});

test('pausado', () => {
  assert.equal(descreverStatus('pausado'), 'em espera');
});

test('cancelado', () => {
  assert.equal(descreverStatus('cancelado'), 'encerrado');
});
