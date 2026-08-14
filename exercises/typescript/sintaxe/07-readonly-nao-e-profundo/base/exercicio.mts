interface Opcoes {
  tema: string;
}

interface Config {
  readonly nome: string;
  readonly opcoes: Opcoes;
}

// TODO: readonly aqui é RASO — ele impede `config.opcoes = outraCoisa`, mas
// não protege os campos de dentro de opcoes. Mutar config.opcoes.tema
// compila sem erro e muda o objeto original, que o objective pede para
// preservar.
export function comTemaEscuro(config: Config): Config {
  config.opcoes.tema = "escuro";
  return config;
}
