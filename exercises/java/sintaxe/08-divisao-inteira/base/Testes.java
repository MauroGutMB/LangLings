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
        verificar("media(4,2)", 2.0, Exercicio.media(4, 2));

        // A asserção que separa a versão ingênua da correta: 5/2 é 2.5, mas
        // a divisão inteira trunca para 2 antes de virar double.
        verificar("media(5,2)", 2.5, Exercicio.media(5, 2));
        verificar("media(1,3)", 1.0 / 3.0, Exercicio.media(1, 3));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
