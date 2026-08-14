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
        verificar("inverter(\"ana\")", "ana", Exercicio.inverter("ana"));
        verificar("inverter(\"langlings\")", "sgnilgnal", Exercicio.inverter("langlings"));
        verificar("inverter(\"\")", "", Exercicio.inverter(""));
        verificar("inverter(\"a\")", "a", Exercicio.inverter("a"));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
