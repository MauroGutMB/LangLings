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
        verificar("diaDaSemana(1)", "domingo", Exercicio.diaDaSemana(1));
        verificar("diaDaSemana(4)", "quarta", Exercicio.diaDaSemana(4));
        verificar("diaDaSemana(7)", "sabado", Exercicio.diaDaSemana(7));
        verificar("diaDaSemana(0)", "invalido", Exercicio.diaDaSemana(0));
        verificar("diaDaSemana(8)", "invalido", Exercicio.diaDaSemana(8));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
