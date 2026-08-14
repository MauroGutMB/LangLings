class Exercicio {

    static void exemplos() {
        // void: não devolve valor, só produz efeito (aqui, imprimir).
        anunciar("turma cheia");

        // Devolve um valor para quem chamou usar.
        int dobro = dobrar(21);
        System.out.println(dobro); // 42

        // Sobrecarga: mesmo nome, assinaturas diferentes. O compilador
        // escolhe qual chamar pela quantidade e pelo tipo dos argumentos.
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

    // SUA VEZ
    //
    // Devolva o maior entre a, b e c.
    static int maiorDeTres(int a, int b, int c) {
        return a; // <- troque isto
    }
}
