import java.util.Objects;

// Este harness não é editável. exemplos() não é chamado aqui de propósito: a
// saída de demonstração não deve se misturar com o resultado dos testes. Para
// ver exemplos() rodar, abra o shell do exercício ([s]) e use jshell:
//   jshell> new Exercicio().exemplos()  // ou: jshell Exercicio.java, depois exemplos()
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
        verificar("celsiusParaFahrenheit(0)", 32.0, Exercicio.celsiusParaFahrenheit(0));
        verificar("celsiusParaFahrenheit(100)", 212.0, Exercicio.celsiusParaFahrenheit(100));
        verificar("celsiusParaFahrenheit(37)", 98.6, Exercicio.celsiusParaFahrenheit(37));
        verificar("celsiusParaFahrenheit(-40)", -40.0, Exercicio.celsiusParaFahrenheit(-40));

        if (falhas > 0) {
            System.out.printf("%n%d verificação(ões) falharam%n", falhas);
            System.exit(1);
        }
        System.out.println("\ntodas as verificações passaram");
    }
}
