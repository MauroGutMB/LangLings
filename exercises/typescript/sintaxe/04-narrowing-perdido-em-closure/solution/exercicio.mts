export function formatarMaisTarde(valor: string | number): () => string {
  const resultado =
    typeof valor === "string" ? valor.toUpperCase() : valor.toFixed(2);
  return () => resultado;
}
