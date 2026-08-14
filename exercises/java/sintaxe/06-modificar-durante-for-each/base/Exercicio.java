import java.util.ArrayList;
import java.util.List;

// Devolva uma nova lista sem os números negativos.
class Exercicio {
    static List<Integer> removerNegativos(List<Integer> numeros) {
        List<Integer> copia = new ArrayList<>(numeros);
        for (Integer n : copia) {
            if (n < 0) {
                copia.remove(n);
            }
        }
        return copia;
    }
}
