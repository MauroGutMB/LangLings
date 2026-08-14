import { test } from 'node:test';
import assert from 'node:assert/strict';
import { clonarConfig } from './exercicio.mjs';

test('a cópia tem os mesmos campos do primeiro nível', () => {
  const config = { nome: 'app', tema: { cor: 'azul' } };
  const clone = clonarConfig(config);
  assert.equal(clone.nome, 'app');
  assert.equal(clone.tema.cor, 'azul');
});

test('a cópia é um objeto diferente do original', () => {
  const config = { nome: 'app', tema: { cor: 'azul' } };
  const clone = clonarConfig(config);
  assert.notEqual(clone, config);
});

// A asserção que separa cópia rasa de cópia funda: alterar o aninhado da
// cópia não pode vazar para o original.
test('alterar o tema da cópia não afeta o original', () => {
  const config = { nome: 'app', tema: { cor: 'azul' } };
  const clone = clonarConfig(config);

  clone.tema.cor = 'vermelho';

  assert.equal(config.tema.cor, 'azul');
});
