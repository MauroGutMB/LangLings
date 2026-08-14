class Exercicio {

    static void exemplos() {
        int nota = 7;

        // if/else de sempre.
        if (nota >= 6) {
            System.out.println("aprovado"); // aprovado
        } else {
            System.out.println("reprovado");
        }

        // switch clássico: sem break, a execução cai para o case seguinte
        // (fall-through) — é o motivo de cada case aqui terminar com break.
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

        // switch em forma de seta: cada case devolve um valor, sem
        // fall-through e sem break. É a forma preferida quando dá.
        int diaNumero = 3;
        String tipo = switch (diaNumero) {
            case 1, 7 -> "fim de semana";
            case 2, 3, 4, 5, 6 -> "dia util";
            default -> "invalido";
        };
        System.out.println(tipo); // dia util
    }

    // SUA VEZ
    //
    // Devolva o nome do dia da semana (1 = domingo ... 7 = sabado); qualquer
    // outro valor devolve "invalido". Use switch em forma de seta.
    static String diaDaSemana(int n) {
        return ""; // <- troque isto
    }
}
