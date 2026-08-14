type Sucesso = { status: "ok"; dado: string };
type Falha = { status: "erro"; mensagem: string };
type Resultado = Sucesso | Falha;

// TODO: dado só existe em Sucesso — Falha não tem esse campo. Acessar
// resultado.dado sem antes checar resultado.status é erro de compilação.
export function descrever(resultado: Resultado): string {
  return resultado.dado;
}
