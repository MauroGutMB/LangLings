import test from 'node:test';
import assert from 'node:assert/strict';
import { apresentarPessoa } from './exercicio.mts';

test('apresenta nome e idade corretamente', () => {
  assert.equal(apresentarPessoa({ nome: 'Ana', idade: 30 }), 'Olá, Ana, você tem 30 anos');
});

test('outro valor de nome e idade', () => {
  assert.equal(apresentarPessoa({ nome: 'Bruno', idade: 42 }), 'Olá, Bruno, você tem 42 anos');
});
