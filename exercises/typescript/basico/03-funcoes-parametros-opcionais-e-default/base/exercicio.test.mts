import test from 'node:test';
import assert from 'node:assert/strict';
import { formatarNome } from './exercicio.mts';

test('sem sobrenome devolve só o nome', () => {
  assert.equal(formatarNome('Ana'), 'Ana');
});

test('com sobrenome devolve os dois separados por espaço', () => {
  assert.equal(formatarNome('Ana', 'Silva'), 'Ana Silva');
});

test('sobrenome explicitamente undefined se comporta como omitido', () => {
  assert.equal(formatarNome('Bruno', undefined), 'Bruno');
});
