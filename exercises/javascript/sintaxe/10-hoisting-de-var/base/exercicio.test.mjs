import { test } from 'node:test';
import assert from 'node:assert/strict';
import { criarContadores } from './exercicio.mjs';

test('devolve um array com n funções', () => {
  const funcoes = criarContadores(3);
  assert.equal(funcoes.length, 3);
  for (const f of funcoes) {
    assert.equal(typeof f, 'function');
  }
});

test('a primeira função devolve 0', () => {
  const funcoes = criarContadores(3);
  assert.equal(funcoes[0](), 0);
});

// A asserção que separa var de let: com var, todas as closures compartilham
// a mesma variável, e todas devolvem o valor final do laço (n), não o índice
// de quando foram criadas.
test('cada função devolve o índice da sua posição', () => {
  const funcoes = criarContadores(4);
  assert.deepEqual(
    funcoes.map((f) => f()),
    [0, 1, 2, 3],
  );
});
