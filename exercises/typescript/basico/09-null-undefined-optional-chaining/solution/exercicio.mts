interface Endereco {
  cidade: string;
}

interface Usuario {
  nome: string;
  endereco?: Endereco;
}

// ?. (optional chaining) acessa um campo só se o lado esquerdo não for
// null/undefined — se for, a cadeia inteira "encurta" e devolve undefined,
// em vez de lançar TypeError como um acesso comum faria.
export function exemplos(): void {
  const comEndereco: Usuario = { nome: "Ana", endereco: { cidade: "Recife" } };
  const semEndereco: Usuario = { nome: "Bruno" };

  console.log(comEndereco.endereco?.cidade); // "Recife"
  console.log(semEndereco.endereco?.cidade); // undefined, sem lançar erro

  // ?? (nullish coalescing) troca por um default, mas SÓ quando o valor da
  // esquerda é null ou undefined — 0 e "" passam direto, diferente de ||.
  const cidadeOuDefault = semEndereco.endereco?.cidade ?? "sem cidade";
  console.log(cidadeOuDefault); // "sem cidade"

  // As duas combinam naturalmente: encurta com ?., substitui o undefined
  // resultante com ??.
  function nomeDaCidade(u: Usuario): string {
    return u.endereco?.cidade ?? "desconhecida";
  }
  console.log(nomeDaCidade(comEndereco)); // "Recife"
  console.log(nomeDaCidade(semEndereco)); // "desconhecida"

  // ?. também funciona em chamada de método: só chama se o método existir.
  console.log(comEndereco.endereco?.cidade.toUpperCase()); // "RECIFE"
  console.log(semEndereco.endereco?.cidade.toUpperCase()); // undefined, não lança
}
// Para ver a saída: abra o shell com [s] e rode
// `node --eval "import('./exercicio.mts').then(m => m.exemplos())"`

// SUA VEZ
//
// Devolva a cidade do endereço de usuario, ou "desconhecida" se não houver
// endereço ou o endereço não tiver cidade.
export function obterCidade(usuario: Usuario): string {
  return usuario.endereco?.cidade ?? "desconhecida";
}
