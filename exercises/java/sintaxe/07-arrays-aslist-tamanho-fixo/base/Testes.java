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
        // A asserção que separa a versão ingênua da correta: adicionar um
        // item à lista devolvida por Arrays.asList lança
        // UnsupportedOperationException — a lista tem tamanho fixo.
        verificar("adiciona a um array de duas posições",
                List.of("a", "b", "c"),
                Exercicio.adicionar(new String[] {"a", "b"}, "c"));

        verificar("adiciona a um array vazio",
                List.of("primeiro"),
                Exercicio.adicionar(new String[] {}, "primeiro"));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
