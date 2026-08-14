// TODO: dentro do if, `atual` está estreitado para string — mas `atual` é
// `let`, e mais abaixo o código reatribui a ela (para guardar uma versão
// normalizada). Só essa reatribuição já basta para TypeScript deixar de
// confiar no estreitamento dentro da closure devolvida, mesmo a closure
// tendo sido criada ANTES da reatribuição no código.
export function formatarMaisTarde(valor: string | number): () => string {
  let atual = valor;
  if (typeof atual === "string") {
    const formatarString = () => atual.toUpperCase();
    atual = atual.trim(); // guarda a versão normalizada para reaproveitar
    return formatarString;
  }
  return () => atual.toFixed(2);
}
