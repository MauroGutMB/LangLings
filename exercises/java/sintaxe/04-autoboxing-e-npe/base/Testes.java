import java.util.Map;
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
        Map<String, Integer> pontuacoes = Map.of("ana", 10, "bruno", 20);

        verificar("chave existente", 10, Exercicio.pontuacao(pontuacoes, "ana"));

        // A asserção que separa a versão ingênua da correta: chave ausente
        // devolveria um Integer nulo, e o unboxing automático para int
        // lançaria NullPointerException aqui.
        verificar("chave ausente", 0, Exercicio.pontuacao(pontuacoes, "carla"));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
