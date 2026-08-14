import test from 'node:test';
import assert from 'node:assert/strict';
import { celsiusParaFahrenheit } from './exercicio.mts';

test('0 graus Celsius é 32 Fahrenheit', () => {
  assert.equal(celsiusParaFahrenheit(0), 32);
});

test('100 graus Celsius é 212 Fahrenheit', () => {
  assert.equal(celsiusParaFahrenheit(100), 212);
});

test('-40 graus é o ponto em que as escalas se cruzam', () => {
  assert.equal(celsiusParaFahrenheit(-40), -40);
});
