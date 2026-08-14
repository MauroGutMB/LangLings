import java.util.List;
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
        verificar("pares([1,2,3,4,5,6])", List.of(2, 4, 6),
                Exercicio.pares(List.of(1, 2, 3, 4, 5, 6)));
        verificar("pares([1,3,5])", List.of(), Exercicio.pares(List.of(1, 3, 5)));
        verificar("pares([])", List.of(), Exercicio.pares(List.of()));
        verificar("pares([2,4])", List.of(2, 4), Exercicio.pares(List.of(2, 4)));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
