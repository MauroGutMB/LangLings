import test from 'node:test';
import assert from 'node:assert/strict';
import { precoComDesconto } from './exercicio.mts';

test('20% de desconto em 100', () => {
  assert.equal(precoComDesconto(100, 20), 80);
});

test('0% de desconto não muda o preço', () => {
  assert.equal(precoComDesconto(50, 0), 50);
});

test('100% de desconto zera o preço', () => {
  assert.equal(precoComDesconto(30, 100), 0);
});

test('10% de desconto em 250', () => {
  assert.equal(precoComDesconto(250, 10), 225);
});
