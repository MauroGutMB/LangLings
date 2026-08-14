type Sucesso = { status: "ok"; dado: string };
type Falha = { status: "erro"; mensagem: string };
type Resultado = Sucesso | Falha;

export function descrever(resultado: Resultado): string {
  if (resultado.status === "ok") {
    return resultado.dado;
  }
  return `erro: ${resultado.mensagem}`;
}
