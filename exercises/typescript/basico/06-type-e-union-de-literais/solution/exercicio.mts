// type cria um alias para um tipo — um nome novo para algo que já existia.
type IdOuNome = number | string;

export function exemplos(): void {
  function buscar(chave: IdOuNome): string {
    return `buscando por ${chave}`;
  }
  console.log(buscar(42));
  console.log(buscar("ana"));

  // Uma union de LITERAIS restringe a um conjunto fechado de valores
  // específicos — não "qualquer string", só essas strings exatas.
  type Direcao = "norte" | "sul" | "leste" | "oeste";

  function mover(direcao: Direcao): void {
    console.log(`movendo para ${direcao}`);
  }
  mover("norte"); // ok: está na lista
  // mover("cima"); // erro de compilação: "cima" não é uma Direcao

  // Diferença de string: uma variável Direcao só aceita esses 4 valores,
  // enquanto uma variável string aceitaria qualquer texto, incluindo "cima".
  let d: Direcao = "sul";
  d = "leste"; // ok: também está na lista
  console.log(d);
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

type Status = "ativo" | "pausado" | "cancelado";

// SUA VEZ
//
// "ativo" -> "em andamento", "pausado" -> "em espera",
// "cancelado" -> "encerrado".
export function descreverStatus(status: Status): string {
  switch (status) {
    case "ativo":
      return "em andamento";
    case "pausado":
      return "em espera";
    case "cancelado":
      return "encerrado";
  }
}
