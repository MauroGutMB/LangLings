class Exercicio {

    static void exemplos() {
        int nota = 7;

        if (nota >= 6) {
            System.out.println("aprovado"); // aprovado
        } else {
            System.out.println("reprovado");
        }

        int mes = 2;
        String estacao;
        switch (mes) {
            case 12:
            case 1:
            case 2:
                estacao = "verao";
                break;
            case 3:
            case 4:
            case 5:
                estacao = "outono";
                break;
            default:
                estacao = "outra";
                break;
        }
        System.out.println(estacao); // verao

        int diaNumero = 3;
        String tipo = switch (diaNumero) {
            case 1, 7 -> "fim de semana";
            case 2, 3, 4, 5, 6 -> "dia util";
            default -> "invalido";
        };
        System.out.println(tipo); // dia util
    }

    // Nome do dia da semana; qualquer valor fora de 1..7 é inválido.
    static String diaDaSemana(int n) {
        return switch (n) {
            case 1 -> "domingo";
            case 2 -> "segunda";
            case 3 -> "terca";
            case 4 -> "quarta";
            case 5 -> "quinta";
            case 6 -> "sexta";
            case 7 -> "sabado";
            default -> "invalido";
        };
    }
}
