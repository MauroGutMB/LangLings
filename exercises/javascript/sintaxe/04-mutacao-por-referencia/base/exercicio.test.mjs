import { test } from 'node:test';
import assert from 'node:assert/strict';
import { comNomeMaiusculo } from './exercicio.mjs';

test('o retorno tem o nome em maiúsculas', () => {
  const pessoa = { nome: 'ana', idade: 30 };
  const resultado = comNomeMaiusculo(pessoa);
  assert.equal(resultado.nome, 'ANA');
  assert.equal(resultado.idade, 30);
});

test('o retorno é um objeto diferente do recebido', () => {
  const pessoa = { nome: 'ana' };
  const resultado = comNomeMaiusculo(pessoa);
  assert.notEqual(resultado, pessoa);
});

// A asserção que separa mutar de copiar: depois de chamar a função, o objeto
// original precisa continuar exatamente como estava.
test('o objeto recebido não é alterado', () => {
  const pessoa = { nome: 'ana', idade: 30 };
  comNomeMaiusculo(pessoa);
  assert.equal(pessoa.nome, 'ana');
});
