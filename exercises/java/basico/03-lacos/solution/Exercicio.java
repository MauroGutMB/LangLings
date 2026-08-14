class Exercicio {

    static void exemplos() {
        for (int i = 0; i < 3; i++) {
            System.out.println(i); // 0, 1, 2
        }

        int[] numeros = {10, 20, 30};

        for (int n : numeros) {
            System.out.println(n); // 10, 20, 30
        }

        int restante = numeros.length;
        while (restante > 0) {
            System.out.println(restante); // 3, 2, 1
            restante--;
        }
    }

    // Soma de todos os elementos; array vazio soma 0.
    static int soma(int[] numeros) {
        int total = 0;
        for (int n : numeros) {
            total += n;
        }
        return total;
    }
}
