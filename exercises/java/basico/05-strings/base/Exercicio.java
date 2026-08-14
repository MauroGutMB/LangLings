class Exercicio {

    static void exemplos() {
        String nome = "ana";

        // toUpperCase devolve uma string NOVA — nome continua "ana". String
        // não tem método que altere o objeto original: não existe.
        String maiuscula = nome.toUpperCase();
        System.out.println(maiuscula); // ANA
        System.out.println(nome);      // ana (inalterado)

        System.out.println(nome.length());      // 3
        System.out.println(nome.substring(1));  // na
        System.out.println(nome + " silva");    // ana silva

        // Cada += aqui descarta a string velha e cria uma nova por baixo dos
        // panos. Para poucas concatenações não importa; num laço grande, é
        // desperdício — o que StringBuilder resolve.
        String saudacao = "ola";
        saudacao += ", ";
        saudacao += nome;
        System.out.println(saudacao); // ola, ana

        // StringBuilder é mutável: os métodos alteram o próprio objeto em vez
        // de criar um novo a cada chamada.
        StringBuilder sb = new StringBuilder();
        sb.append("ola");
        sb.append(", ");
        sb.append(nome);
        System.out.println(sb.toString()); // ola, ana
    }

    // SUA VEZ
    //
    // Devolva texto invertido.
    static String inverter(String texto) {
        return texto; // <- troque isto
    }
}
