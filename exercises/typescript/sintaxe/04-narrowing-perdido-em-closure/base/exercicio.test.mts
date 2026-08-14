import test from 'node:test';
import assert from 'node:assert/strict';
import { formatarMaisTarde } from './exercicio.mts';

test('string vira maiúsculas', () => {
  const f = formatarMaisTarde('abc');
  assert.equal(f(), 'ABC');
});

test('number vira string com duas casas decimais', () => {
  const f = formatarMaisTarde(3.14159);
  assert.equal(f(), '3.14');
});

test('number inteiro também recebe duas casas', () => {
  const f = formatarMaisTarde(7);
  assert.equal(f(), '7.00');
});
