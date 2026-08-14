interface Usuario {
  nome: string;
}

// TODO: interface só existe em tempo de compilação — ela é apagada do
// JavaScript gerado. typeof nunca devolve "Usuario": ele só devolve um dos
// tipos primitivos de runtime ("string", "number", "object", "undefined"
// etc.), e comparar com um literal fora desse conjunto é erro de compilação.
export function ehUsuario(valor: unknown): valor is Usuario {
  return typeof valor === "Usuario";
}
