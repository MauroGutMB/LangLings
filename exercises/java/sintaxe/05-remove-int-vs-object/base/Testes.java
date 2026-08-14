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
        // valor 1 não está na lista — a lista deveria voltar inalterada.
        verificar("valor ausente não remove nada",
                List.of(10, 20, 30, 40, 50),
                Exercicio.removerValor(List.of(10, 20, 30, 40, 50), 1));

        // A asserção que separa a versão ingênua da correta: remover o VALOR
        // 20, não o item na posição 20 (que nem existe nesta lista de 3).
        verificar("remove o valor, não o índice",
                List.of(10, 30),
                Exercicio.removerValor(List.of(10, 20, 30), 20));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
