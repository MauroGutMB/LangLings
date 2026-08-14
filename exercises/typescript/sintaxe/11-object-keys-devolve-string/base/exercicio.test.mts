import test from 'node:test';
import assert from 'node:assert/strict';
import { somarTodosOsPrecos } from './exercicio.mts';

test('soma os três valores', () => {
  assert.equal(somarTodosOsPrecos({ dolar: 5, euro: 6, real: 1 }), 12);
});

test('com zeros', () => {
  assert.equal(somarTodosOsPrecos({ dolar: 0, euro: 0, real: 0 }), 0);
});
