class Exercicio {

    static void exemplos() {
        String nome = "ana";

        String maiuscula = nome.toUpperCase();
        System.out.println(maiuscula); // ANA
        System.out.println(nome);      // ana (inalterado)

        System.out.println(nome.length());      // 3
        System.out.println(nome.substring(1));  // na
        System.out.println(nome + " silva");    // ana silva

        String saudacao = "ola";
        saudacao += ", ";
        saudacao += nome;
        System.out.println(saudacao); // ola, ana

        StringBuilder sb = new StringBuilder();
        sb.append("ola");
        sb.append(", ");
        sb.append(nome);
        System.out.println(sb.toString()); // ola, ana
    }

    // Texto invertido.
    static String inverter(String texto) {
        return new StringBuilder(texto).reverse().toString();
    }
}
