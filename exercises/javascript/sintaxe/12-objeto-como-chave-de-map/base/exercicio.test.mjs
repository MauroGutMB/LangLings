import { test } from 'node:test';
import assert from 'node:assert/strict';
import { buscarValor } from './exercicio.mjs';

test('id ausente devolve undefined', () => {
  const mapa = new Map([[{ id: 1 }, 'a']]);
  assert.equal(buscarValor(mapa, 99), undefined);
});

test('mapa vazio devolve undefined para qualquer id', () => {
  const mapa = new Map();
  assert.equal(buscarValor(mapa, 1), undefined);
});

// A asserção que separa mapa.get({ id }) de percorrer as entradas: a chave
// que existe de verdade no mapa é uma referência criada antes desta chamada,
// e { id } aqui dentro cria uma referência nova a cada vez.
test('encontra o valor pela referência já existente no mapa', () => {
  const chaveA = { id: 1 };
  const chaveB = { id: 2 };
  const mapa = new Map([
    [chaveA, 'valor-a'],
    [chaveB, 'valor-b'],
  ]);

  assert.equal(buscarValor(mapa, 2), 'valor-b');
  assert.equal(buscarValor(mapa, 1), 'valor-a');
});
