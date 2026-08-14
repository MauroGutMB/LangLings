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
        verificar("nada para remover", List.of(1, 2, 3),
                Exercicio.removerNegativos(List.of(1, 2, 3)));

        // A asserção que separa a versão ingênua da correta: removendo
        // enquanto o for-each anda sobre a lista lança
        // ConcurrentModificationException em vez de devolver a lista
        // filtrada. (Removido o penúltimo elemento não dispara a exceção por
        // uma peculiaridade do ArrayList — por isso os casos abaixo evitam
        // essa posição.)
        verificar("remove o primeiro (negativo)", List.of(2, 3),
                Exercicio.removerNegativos(List.of(-1, 2, 3)));
        verificar("remove negativo no meio, longe do fim", List.of(1, 3, 4),
                Exercicio.removerNegativos(List.of(1, -2, 3, 4)));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
