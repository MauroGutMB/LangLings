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
        verificar("100 == 100 (dentro do cache)", true, Exercicio.saoIguais(100, 100));
        verificar("valores diferentes", false, Exercicio.saoIguais(1, 2));

        // A asserção que separa == de .equals()/intValue(): fora da faixa
        // cacheada (-128..127), cada autoboxing cria um objeto Integer novo,
        // e == passa a comparar referências diferentes mesmo com o mesmo
        // valor.
        verificar("200 == 200 (fora do cache)", true, Exercicio.saoIguais(200, 200));
        verificar("-500 == -500 (fora do cache)", true, Exercicio.saoIguais(-500, -500));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
