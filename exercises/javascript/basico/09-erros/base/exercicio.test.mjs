import test from 'node:test';
import assert from 'node:assert/strict';
import { dividir, ErroDeEntrada } from './exercicio.mjs';

test('divide normalmente', () => {
  assert.equal(dividir(6, 2), 3);
  assert.equal(dividir(1, 4), 0.25);
  assert.equal(dividir(0, 5), 0);
});

// Sem o throw, isto devolveria Infinity e o teste passaria batido.
test('divisão por zero lança em vez de devolver Infinity', () => {
  assert.throws(() => dividir(1, 0));
});

test('o que é lançado é um ErroDeEntrada com a mensagem certa', () => {
  assert.throws(() => dividir(1, 0), (erro) => {
    assert.ok(erro instanceof ErroDeEntrada, 'esperava um ErroDeEntrada');
    assert.equal(erro.message, 'divisão por zero');
    return true;
  });
});

test('zero dividido por zero também lança', () => {
  assert.throws(() => dividir(0, 0), ErroDeEntrada);
});
