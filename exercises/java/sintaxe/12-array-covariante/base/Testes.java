import java.util.Arrays;
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
        Integer[] inteiros = {1, 2, 3};
        Exercicio.preencherComZero(inteiros);
        verificar("Integer[] vira zeros", "[0, 0, 0]", Arrays.toString(inteiros));

        // A asserção que separa a versão ingênua da correta: numeros é
        // declarado Number[], mas o array por baixo é um Double[] de
        // verdade — gravar um Integer nele lança ArrayStoreException.
        Double[] doubles = {1.5, 2.5, 3.5};
        Exercicio.preencherComZero(doubles);
        verificar("Double[] vira zeros", "[0.0, 0.0, 0.0]", Arrays.toString(doubles));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
