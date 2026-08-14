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
        verificar("Retangulo(3,4).perimetro()", 14.0,
                new Exercicio.Retangulo(3, 4).perimetro());
        verificar("Retangulo(1,1).perimetro()", 4.0,
                new Exercicio.Retangulo(1, 1).perimetro());
        verificar("Retangulo(2.5,2).perimetro()", 9.0,
                new Exercicio.Retangulo(2.5, 2).perimetro());

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
