class Exercicio {

    static void exemplos() {
        // for clássico: inicialização, condição, incremento.
        for (int i = 0; i < 3; i++) {
            System.out.println(i); // 0, 1, 2
        }

        int[] numeros = {10, 20, 30};

        // for-each: percorre cada elemento direto, sem índice. Prefira esta
        // forma quando você não precisa da posição.
        for (int n : numeros) {
            System.out.println(n); // 10, 20, 30
        }

        // while repete enquanto a condição for verdadeira; a condição é
        // checada antes de cada volta, então pode nunca executar o corpo.
        int restante = numeros.length;
        while (restante > 0) {
            System.out.println(restante); // 3, 2, 1
            restante--;
        }
    }

    // SUA VEZ
    //
    // Devolva a soma de todos os elementos de numeros. Array vazio soma 0.
    static int soma(int[] numeros) {
        return -1; // <- troque isto
    }
}
