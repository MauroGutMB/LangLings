import java.util.List;

// Devolva os itens concatenados com separador entre eles.
class Exercicio {
    static String juntar(List<String> itens, String separador) {
        String resultado = "";
        for (String item : itens) {
            resultado += item + separador;
        }
        return resultado;
    }
}
