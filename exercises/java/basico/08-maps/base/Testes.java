import java.util.List;
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
        verificar("contagem([a,b,a])",
                Map.of("a", 2, "b", 1),
                Exercicio.contagem(List.of("a", "b", "a")));
        verificar("contagem([])", Map.of(), Exercicio.contagem(List.of()));
        verificar("contagem([Ana,ana])",
                Map.of("Ana", 1, "ana", 1),
                Exercicio.contagem(List.of("Ana", "ana")));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
