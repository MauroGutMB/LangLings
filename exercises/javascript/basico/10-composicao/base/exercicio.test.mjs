import test from 'node:test';
import assert from 'node:assert/strict';
import { resumir, ErroDeEstoque } from './exercicio.mjs';

const itens = [
  { nome: 'caneta', categoria: 'papelaria', preco: 250, qtd: 4 },
  { nome: 'caderno', categoria: 'papelaria', preco: 1200, qtd: 2 },
  { nome: 'café', categoria: 'cozinha', preco: 1850, qtd: 3 },
];

test('soma o valor de cada categoria', () => {
  const total = resumir(itens);

  assert.equal(total.size, 2);
  assert.equal(total.get('papelaria'), 3400);
  assert.equal(total.get('cozinha'), 5550);
});

test('um item só continua sendo uma categoria', () => {
  const total = resumir([{ nome: 'café', categoria: 'cozinha', preco: 100, qtd: 2 }]);

  assert.equal(total.size, 1);
  assert.equal(total.get('cozinha'), 200);
});

test('qtd zero soma zero, mas a categoria existe', () => {
  const total = resumir([{ nome: 'caneta', categoria: 'papelaria', preco: 250, qtd: 0 }]);

  assert.equal(total.get('papelaria'), 0);
});

test('array vazio lança ErroDeEstoque', () => {
  assert.throws(() => resumir([]), (erro) => {
    assert.ok(erro instanceof ErroDeEstoque, 'esperava um ErroDeEstoque');
    assert.equal(erro.message, 'nenhum item');
    return true;
  });
});

test('item sem categoria lança, citando o nome do item', () => {
  const invalidos = [
    { nome: 'caneta', categoria: 'papelaria', preco: 250, qtd: 4 },
    { nome: 'misterioso', categoria: '', preco: 100, qtd: 1 },
  ];

  assert.throws(() => resumir(invalidos), (erro) => {
    assert.ok(erro instanceof ErroDeEstoque, 'esperava um ErroDeEstoque');
    assert.ok(erro.message.includes('misterioso'), `mensagem sem o nome do item: ${erro.message}`);
    return true;
  });
});
