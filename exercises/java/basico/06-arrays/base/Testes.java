import java.util.Objects;

public class Testes {
    private static int falhas = 0;

    static void verificar(String oQue, Object esperado, Object obtido) {
        if (Objects.deepEquals(esperado, obtido)) {
            System.out.println("ok    " + oQue);
            return;
        }
        System.out.printf("FALHA %s%n      esperado: %s%n      obtido:   %s%n",
                oQue, esperado, obtido);
        falhas++;
    }

    public static void main(String[] args) {
        verificar("maximo({1,5,3})", 5, Exercicio.maximo(new int[] {1, 5, 3}));
        verificar("maximo({-1,-5,-3})", -1, Exercicio.maximo(new int[] {-1, -5, -3}));
        verificar("maximo({7})", 7, Exercicio.maximo(new int[] {7}));
        verificar("maximo({3,3,3})", 3, Exercicio.maximo(new int[] {3, 3, 3}));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
