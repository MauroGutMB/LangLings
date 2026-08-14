// Sobrescreva cada posição de numeros com o zero apropriado ao tipo real do
// array.
class Exercicio {
    static void preencherComZero(Number[] numeros) {
        for (int i = 0; i < numeros.length; i++) {
            numeros[i] = 0; // sempre grava um Integer, não importa o array real
        }
    }
}
