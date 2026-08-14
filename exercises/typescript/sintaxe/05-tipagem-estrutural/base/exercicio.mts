type Metros = number;
type Segundos = number;

// TODO: Metros e Segundos são só apelidos de number — o compilador não
// distingue um do outro, então a divisão abaixo pode estar de cabeça para
// baixo sem que nenhum erro de tipo apareça.
export function velocidadeMedia(
  distanciaEmMetros: Metros,
  tempoEmSegundos: Segundos,
): number {
  return tempoEmSegundos / distanciaEmMetros;
}
