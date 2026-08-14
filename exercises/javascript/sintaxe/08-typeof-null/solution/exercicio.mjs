// Devolva true quando valor for um objeto de verdade. null não conta.
export function ehObjeto(valor) {
  return valor !== null && typeof valor === 'object';
}
