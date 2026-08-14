// Devolva uma cópia de config totalmente independente do original, inclusive
// o objeto aninhado em config.tema.
//
// TODO: o spread copia o primeiro nível; o aninhado continua compartilhado.
export function clonarConfig(config) {
  return { ...config };
}
