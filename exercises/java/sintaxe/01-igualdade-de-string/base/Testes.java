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
        verificar("literais iguais", true, Exercicio.saoIguais("java", "java"));
        verificar("conteúdos diferentes", false, Exercicio.saoIguais("java", "kotlin"));

        // A asserção que separa == de .equals(): mesmo conteúdo, mas objetos
        // construídos em tempo de execução (não internados pelo compilador)
        // não são a mesma referência — == diz que são diferentes, errado.
        String a = new String("langlings");
        String b = new String("langlings");
        verificar("mesmo conteúdo, objetos distintos", true, Exercicio.saoIguais(a, b));

        String c = "lang" + "lings"; // concatenado de literais: constante em tempo de compilação
        String d = new StringBuilder("lang").append("lings").toString(); // construído em runtime
        verificar("conteúdo igual, um deles construído em runtime", true, Exercicio.saoIguais(c, d));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
