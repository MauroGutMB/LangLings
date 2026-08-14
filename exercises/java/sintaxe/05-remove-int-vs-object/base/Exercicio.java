import java.util.ArrayList;
import java.util.List;

// Devolva uma nova lista sem a primeira ocorrência do valor.
class Exercicio {
    static List<Integer> removerValor(List<Integer> numeros, int valor) {
        List<Integer> copia = new ArrayList<>(numeros);
        copia.remove(valor);
        return copia;
    }
}
