import test from 'node:test';
import assert from 'node:assert/strict';
import { comPadroes } from './exercicio.mjs';

test('sem configuração, devolve só os padrões', () => {
  assert.deepEqual(comPadroes({}), { tema: 'claro', fonte: 14, animacoes: true });
});

test('o que a configuração diz vence o padrão', () => {
  assert.deepEqual(comPadroes({ tema: 'escuro' }), {
    tema: 'escuro',
    fonte: 14,
    animacoes: true,
  });
});

test('completa só o que faltou', () => {
  const resultado = comPadroes({ fonte: 20, animacoes: false });
  assert.equal(resultado.fonte, 20);
  assert.equal(resultado.animacoes, false);
  assert.equal(resultado.tema, 'claro');
});

test('uma chave que não é padrão continua no resultado', () => {
  assert.equal(comPadroes({ idioma: 'pt' }).idioma, 'pt');
});

test('não altera a configuração recebida', () => {
  const config = { tema: 'escuro' };
  comPadroes(config);
  assert.deepEqual(config, { tema: 'escuro' });
});
