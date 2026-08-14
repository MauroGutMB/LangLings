import test from 'node:test';
import assert from 'node:assert/strict';
import { comTemaEscuro } from './exercicio.mts';

test('devolve uma config com tema escuro', () => {
  const original = { nome: 'padrão', opcoes: { tema: 'claro' } };
  const resultado = comTemaEscuro(original);
  assert.equal(resultado.opcoes.tema, 'escuro');
});

test('não muda o objeto original', () => {
  const original = { nome: 'padrão', opcoes: { tema: 'claro' } };
  comTemaEscuro(original);
  assert.equal(original.opcoes.tema, 'claro');
});
