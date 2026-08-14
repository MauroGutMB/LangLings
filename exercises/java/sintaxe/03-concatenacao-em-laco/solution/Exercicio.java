import java.util.List;

// Devolva os itens concatenados com separador entre eles.
class Exercicio {
    static String juntar(List<String> itens, String separador) {
        return String.join(separador, itens);
    }
}
