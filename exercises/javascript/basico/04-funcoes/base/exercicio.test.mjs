import test from 'node:test';
import assert from 'node:assert/strict';
import { repetir } from './exercicio.mjs';

test('repete o número de vezes pedido', () => {
  assert.equal(repetir('ab', 3), 'ababab');
  assert.equal(repetir('x', 1), 'x');
});

test('sem o segundo argumento, repete 2 vezes', () => {
  assert.equal(repetir('ab'), 'abab');
  assert.equal(repetir('-'), '--');
});

// Passar 0 é passar um valor: o padrão não substitui.
test('zero vezes devolve string vazia', () => {
  assert.equal(repetir('x', 0), '');
});

test('repetir texto vazio devolve texto vazio', () => {
  assert.equal(repetir('', 5), '');
});
