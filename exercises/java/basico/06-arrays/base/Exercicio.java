import java.util.Arrays;

class Exercicio {

    static void exemplos() {
        // Tamanho fixo desde a criação: um array não cresce nem encolhe.
        int[] numeros = {10, 20, 30};

        System.out.println(numeros.length);   // 3
        System.out.println(numeros[0]);       // 10
        numeros[1] = 99;                       // escreve por índice
        System.out.println(numeros[1]);       // 99

        // println num array não imprime o conteúdo — imprime algo como
        // "[I@1b6d3586", um identificador interno do tipo. Arrays.toString
        // existe exatamente para isso.
        System.out.println(Arrays.toString(numeros)); // [10, 99, 30]

        // new int[n] cria um array de tamanho n com todo mundo zerado.
        int[] zerados = new int[3];
        System.out.println(Arrays.toString(zerados)); // [0, 0, 0]
    }

    // SUA VEZ
    //
    // Devolva o maior elemento de numeros (não vazio).
    static int maximo(int[] numeros) {
        return 0; // <- troque isto
    }
}
