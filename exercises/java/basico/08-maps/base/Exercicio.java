import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Exercicio {

    static void exemplos() {
        Map<String, Integer> idades = new HashMap<>();
        idades.put("ana", 30);
        idades.put("bruno", 25);

        System.out.println(idades.get("ana")); // 30

        // Ler uma chave ausente com get devolve null, não erro — e desempacotar
        // esse null como int explode em NullPointerException. containsKey
        // evita isso.
        System.out.println(idades.get("zoe"));           // null
        System.out.println(idades.containsKey("zoe"));   // false

        // getOrDefault é o atalho: "pegue o valor, ou use este padrão se a
        // chave não existir" — sem passar por null.
        System.out.println(idades.getOrDefault("zoe", 0)); // 0

        idades.remove("ana");
        System.out.println(idades.size()); // 1

        // O idioma de contagem: getOrDefault(chave, 0) + 1 funciona tanto na
        // primeira ocorrência quanto nas seguintes.
        Map<String, Integer> contagemDeLetras = new HashMap<>();
        for (char c : "banana".toCharArray()) {
            String letra = String.valueOf(c);
            contagemDeLetras.put(letra, contagemDeLetras.getOrDefault(letra, 0) + 1);
        }
        System.out.println(contagemDeLetras.get("a")); // 3
    }

    // SUA VEZ
    //
    // Devolva um map com quantas vezes cada palavra aparece em palavras.
    static Map<String, Integer> contagem(List<String> palavras) {
        return null; // <- troque isto
    }
}
