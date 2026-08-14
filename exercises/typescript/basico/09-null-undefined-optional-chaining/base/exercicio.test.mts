import test from 'node:test';
import assert from 'node:assert/strict';
import { obterCidade } from './exercicio.mts';

test('usuário com endereço e cidade', () => {
  assert.equal(obterCidade({ nome: 'Ana', endereco: { cidade: 'Recife' } }), 'Recife');
});

test('usuário sem endereço', () => {
  assert.equal(obterCidade({ nome: 'Bruno' }), 'desconhecida');
});
