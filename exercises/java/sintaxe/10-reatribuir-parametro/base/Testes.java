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
        // A asserção que separa a versão ingênua da correta: depois de
        // chamar zerarPrimeiro, o ARRAY DO CHAMADOR precisa ter mudado.
        int[] valores = {5, 10, 15};
        Exercicio.zerarPrimeiro(valores);
        verificar("primeiro elemento zerado", 0, valores[0]);
        verificar("os demais elementos continuam", Arrays.toString(new int[] {0, 10, 15}),
                Arrays.toString(valores));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
