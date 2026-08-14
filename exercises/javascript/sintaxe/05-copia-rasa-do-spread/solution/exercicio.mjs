// Devolva uma cópia de config totalmente independente do original, inclusive
// o objeto aninhado em config.tema.
export function clonarConfig(config) {
  return { ...config, tema: { ...config.tema } };
}
