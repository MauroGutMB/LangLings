interface Usuario {
  nome: string;
}

export function ehUsuario(valor: unknown): valor is Usuario {
  return (
    typeof valor === "object" &&
    valor !== null &&
    "nome" in valor &&
    typeof (valor as { nome: unknown }).nome === "string"
  );
}
