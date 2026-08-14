import test from 'node:test';
import assert from 'node:assert/strict';
import { contarOcorrencias } from './exercicio.mjs';

test('agrupa as repetições', () => {
  const contagem = contarOcorrencias(['go', 'rust', 'go', 'go']);

  assert.equal(contagem.size, 2);
  assert.equal(contagem.get('go'), 3);
  assert.equal(contagem.get('rust'), 1);
});

test('devolve um Map de verdade', () => {
  assert.ok(contarOcorrencias(['a']) instanceof Map);
});

test('maiúsculas fazem palavras diferentes', () => {
  const contagem = contarOcorrencias(['Go', 'go']);

  assert.equal(contagem.get('Go'), 1);
  assert.equal(contagem.get('go'), 1);
});

test('não inventa chaves', () => {
  assert.equal(contarOcorrencias(['a']).has('b'), false);
});

test('array vazio devolve um Map vazio', () => {
  assert.equal(contarOcorrencias([]).size, 0);
});
