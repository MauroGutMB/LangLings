import test from 'node:test';
import assert from 'node:assert/strict';
import { descrever } from './exercicio.mts';

test('sucesso devolve o dado', () => {
  assert.equal(descrever({ status: 'ok', dado: 'tudo certo' }), 'tudo certo');
});

test('falha devolve a mensagem prefixada', () => {
  assert.equal(descrever({ status: 'erro', mensagem: 'não encontrado' }), 'erro: não encontrado');
});
