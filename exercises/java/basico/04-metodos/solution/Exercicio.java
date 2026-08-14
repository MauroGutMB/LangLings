class Exercicio {

    static void exemplos() {
        anunciar("turma cheia");

        int dobro = dobrar(21);
        System.out.println(dobro); // 42

        System.out.println(soma(2, 3));       // 5
        System.out.println(soma(2, 3, 4));    // 9
        System.out.println(soma(2.5, 1.5));   // 4.0
    }

    static void anunciar(String mensagem) {
        System.out.println("aviso: " + mensagem);
    }

    static int dobrar(int n) {
        return n * 2;
    }

    static int soma(int a, int b) {
        return a + b;
    }

    static int soma(int a, int b, int c) {
        return a + b + c;
    }

    static double soma(double a, double b) {
        return a + b;
    }

    // Maior entre os três valores.
    static int maiorDeTres(int a, int b, int c) {
        return Math.max(a, Math.max(b, c));
    }
}
