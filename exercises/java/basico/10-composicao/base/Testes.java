import java.util.Map;
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
        verificar("contarLetras(\"banana\")",
                Map.of('b', 1, 'a', 3, 'n', 2),
                Exercicio.contarLetras("banana"));
        verificar("contarLetras(\"a b, a!\")",
                Map.of('a', 2, 'b', 1),
                Exercicio.contarLetras("a b, a!"));
        verificar("contarLetras(\"Aa\")",
                Map.of('A', 1, 'a', 1),
                Exercicio.contarLetras("Aa"));
        verificar("contarLetras(\"\")", Map.of(), Exercicio.contarLetras(""));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
