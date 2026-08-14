interface Opcoes {
  tema: string;
}

interface Config {
  readonly nome: string;
  readonly opcoes: Opcoes;
}

export function comTemaEscuro(config: Config): Config {
  return { ...config, opcoes: { ...config.opcoes, tema: "escuro" } };
}
