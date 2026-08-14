import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

// Devolva uma lista mutável com os itens seguidos de novoItem.
class Exercicio {
    static List<String> adicionar(String[] itens, String novoItem) {
        List<String> lista = new ArrayList<>(Arrays.asList(itens));
        lista.add(novoItem);
        return lista;
    }
}
