import java.util.Arrays;

class Exercicio {

    static void exemplos() {
        int[] numeros = {10, 20, 30};

        System.out.println(numeros.length);   // 3
        System.out.println(numeros[0]);       // 10
        numeros[1] = 99;                       // escreve por índice
        System.out.println(numeros[1]);       // 99

        System.out.println(Arrays.toString(numeros)); // [10, 99, 30]

        int[] zerados = new int[3];
        System.out.println(Arrays.toString(zerados)); // [0, 0, 0]
    }

    // Maior elemento de numeros (não vazio).
    static int maximo(int[] numeros) {
        int maior = numeros[0];
        for (int n : numeros) {
            if (n > maior) {
                maior = n;
            }
        }
        return maior;
    }
}
