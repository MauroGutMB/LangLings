import java.util.HashSet;
import java.util.Objects;
import java.util.Set;

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
        Exercicio.Ponto p1 = new Exercicio.Ponto(1, 2);
        Set<Exercicio.Ponto> pontos = new HashSet<>();
        pontos.add(p1);

        verificar("o mesmo objeto está contido", true, Exercicio.contido(pontos, p1));

        // A asserção que separa a versão ingênua da correta: um objeto
        // DIFERENTE, mas com o mesmo x e y, deveria contar como equivalente.
        verificar("um objeto diferente com mesmo x e y",
                true, Exercicio.contido(pontos, new Exercicio.Ponto(1, 2)));

        verificar("um ponto realmente ausente",
                false, Exercicio.contido(pontos, new Exercicio.Ponto(9, 9)));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
