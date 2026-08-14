// Devolva true quando valor for um objeto de verdade. null não conta.
//
// TODO: typeof null é 'object' — um dos casos especiais mais antigos da
// linguagem — e o código abaixo não trata isso.
export function ehObjeto(valor) {
  return typeof valor === 'object';
}
