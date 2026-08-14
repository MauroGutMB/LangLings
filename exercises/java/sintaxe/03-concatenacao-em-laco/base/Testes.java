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
        verificar("lista vazia", "", Exercicio.juntar(List.of(), ","));

        // A asserção que separa a versão ingênua da correta: com um único
        // item não deveria sobrar separador nenhum.
        verificar("um item só", "a", Exercicio.juntar(List.of("a"), ","));

        verificar("três itens", "a,b,c", Exercicio.juntar(List.of("a", "b", "c"), ","));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
