import java.util.ArrayList;
import java.util.List;

class Exercicio {

    static void exemplos() {
        List<Integer> numeros = new ArrayList<>();
        numeros.add(10);
        numeros.add(20);
        numeros.add(30);

        System.out.println(numeros.size());   // 3
        System.out.println(numeros.get(1));   // 20

        numeros.remove(Integer.valueOf(20)); // remove o VALOR 20, não o índice 20
        System.out.println(numeros);          // [10, 30]

        numeros.add(0, 5); // insere 5 na posição 0, empurrando o resto
        System.out.println(numeros);          // [5, 10, 30]

        List<String> fixa = List.of("a", "b", "c");
        System.out.println(fixa); // [a, b, c]
    }

    // Só os números pares de numeros, na mesma ordem.
    static List<Integer> pares(List<Integer> numeros) {
        List<Integer> resultado = new ArrayList<>();
        for (int n : numeros) {
            if (n % 2 == 0) {
                resultado.add(n);
            }
        }
        return resultado;
    }
}
