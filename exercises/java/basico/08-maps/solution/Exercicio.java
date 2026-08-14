import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Exercicio {

    static void exemplos() {
        Map<String, Integer> idades = new HashMap<>();
        idades.put("ana", 30);
        idades.put("bruno", 25);

        System.out.println(idades.get("ana")); // 30

        System.out.println(idades.get("zoe"));           // null
        System.out.println(idades.containsKey("zoe"));   // false

        System.out.println(idades.getOrDefault("zoe", 0)); // 0

        idades.remove("ana");
        System.out.println(idades.size()); // 1

        Map<String, Integer> contagemDeLetras = new HashMap<>();
        for (char c : "banana".toCharArray()) {
            String letra = String.valueOf(c);
            contagemDeLetras.put(letra, contagemDeLetras.getOrDefault(letra, 0) + 1);
        }
        System.out.println(contagemDeLetras.get("a")); // 3
    }

    // Map com quantas vezes cada palavra aparece.
    static Map<String, Integer> contagem(List<String> palavras) {
        Map<String, Integer> resultado = new HashMap<>();
        for (String p : palavras) {
            resultado.put(p, resultado.getOrDefault(p, 0) + 1);
        }
        return resultado;
    }
}
