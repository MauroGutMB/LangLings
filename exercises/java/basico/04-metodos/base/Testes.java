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
        verificar("maiorDeTres(1,2,3)", 3, Exercicio.maiorDeTres(1, 2, 3));
        verificar("maiorDeTres(3,2,1)", 3, Exercicio.maiorDeTres(3, 2, 1));
        verificar("maiorDeTres(2,3,1)", 3, Exercicio.maiorDeTres(2, 3, 1));
        verificar("maiorDeTres(-5,-1,-9)", -1, Exercicio.maiorDeTres(-5, -1, -9));
        verificar("maiorDeTres(4,4,4)", 4, Exercicio.maiorDeTres(4, 4, 4));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
