import test from 'node:test';
import assert from 'node:assert/strict';
import { apresentar } from './exercicio.mts';

test('com idade informada', () => {
  assert.equal(apresentar({ nome: 'Ana', idade: 30 }), 'Ana (30 anos)');
});

test('sem idade', () => {
  assert.equal(apresentar({ nome: 'Bruno' }), 'Bruno');
});

test('idade explicitamente undefined se comporta como ausente', () => {
  assert.equal(apresentar({ nome: 'Carla', idade: undefined }), 'Carla');
});
