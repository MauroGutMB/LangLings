import test from 'node:test';
import assert from 'node:assert/strict';
import { ehUsuario } from './exercicio.mts';

test('objeto com campo nome string é um Usuario', () => {
  assert.equal(ehUsuario({ nome: 'Ana' }), true);
});

test('objeto sem campo nome não é', () => {
  assert.equal(ehUsuario({ idade: 30 }), false);
});

test('nome que não é string não conta', () => {
  assert.equal(ehUsuario({ nome: 42 }), false);
});

test('null não é um Usuario', () => {
  assert.equal(ehUsuario(null), false);
});

test('string não é um Usuario', () => {
  assert.equal(ehUsuario('Ana'), false);
});
