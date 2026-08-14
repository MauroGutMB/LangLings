// Sobrescreva cada posição de numeros com o zero apropriado ao tipo real do
// array.
class Exercicio {
    static void preencherComZero(Number[] numeros) {
        Class<?> tipoReal = numeros.getClass().getComponentType();
        for (int i = 0; i < numeros.length; i++) {
            if (tipoReal == Double.class) {
                numeros[i] = 0.0;
            } else if (tipoReal == Integer.class) {
                numeros[i] = 0;
            } else {
                numeros[i] = 0; // outros tipos de Number: melhor esforço
            }
        }
    }
}
