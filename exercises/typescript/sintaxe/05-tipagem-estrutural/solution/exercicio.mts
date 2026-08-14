type Metros = number;
type Segundos = number;

export function velocidadeMedia(
  distanciaEmMetros: Metros,
  tempoEmSegundos: Segundos,
): number {
  return distanciaEmMetros / tempoEmSegundos;
}
