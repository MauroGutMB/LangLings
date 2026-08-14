// Tipos primitivos não são objetos: int, double, boolean e char vivem direto
// na pilha, sem alocação. Java também tem uma versão "objeto" de cada um
// (Integer, Double, Boolean, Character) — os wrappers — usada onde um objeto
// é exigido, como em List<Integer>.
class Exercicio {

    static void exemplos() {
        int idade = 30;
        double altura = 1.75; // double é o tipo padrão para números com casa decimal
        boolean maiorDeIdade = idade >= 18;
        char inicial = 'a';

        System.out.println(idade);          // 30
        System.out.println(altura);         // 1.75
        System.out.println(maiorDeIdade);   // true
        System.out.println(inicial);        // a

        // Autoboxing: um int é embrulhado num Integer sem cast explícito. O
        // caminho inverso (unboxing) acontece igual de silencioso — e é o que
        // torna perigoso desempacotar um Integer nulo (isso vira NPE, mas é
        // assunto de outro exercício).
        Integer idadeEmbrulhada = idade;
        System.out.println(idadeEmbrulhada.equals(30)); // true

        // int -> double é automático (widening): nenhuma casa é perdida.
        double idadeComoDouble = idade;
        System.out.println(idadeComoDouble); // 30.0

        // double -> int exige cast explícito e trunca, não arredonda.
        double preco = 9.99;
        int precoTruncado = (int) preco;
        System.out.println(precoTruncado); // 9
    }

    // SUA VEZ
    //
    // Converta celsius para Fahrenheit: F = C * 9 / 5 + 32.
    static double celsiusParaFahrenheit(double celsius) {
        return 0.0; // <- troque isto
    }
}
