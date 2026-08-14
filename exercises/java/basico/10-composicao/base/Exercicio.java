import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Exercicio {

    static void exemplos() {
        String frase = "java e divertido";
        String[] palavras = frase.split(" ");
        System.out.println(palavras.length); // 3

        StringBuilder iniciais = new StringBuilder();
        for (String p : palavras) {
            iniciais.append(Character.toUpperCase(p.charAt(0)));
        }
        System.out.println(iniciais.toString()); // JED

        Map<String, Integer> tamanhos = new HashMap<>();
        for (String p : palavras) {
            tamanhos.put(p, p.length());
        }
        System.out.println(tamanhos.get("java")); // 4

        List<String> longas = new ArrayList<>();
        for (String p : palavras) {
            if (p.length() > 2) {
                longas.add(p);
            }
        }
        System.out.println(longas); // [java, divertido]
    }

    // SUA VEZ
    //
    // Devolva quantas vezes cada letra aparece em texto. Ignore espaços e
    // qualquer caractere que não seja letra; maiúscula e minúscula contam
    // como letras diferentes.
    static Map<Character, Integer> contarLetras(String texto) {
        return null; // <- troque isto
    }
}
