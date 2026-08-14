import java.util.Map;

// Devolva a pontuação associada a chave, ou 0 se a chave não existir.
class Exercicio {
    static int pontuacao(Map<String, Integer> pontuacoes, String chave) {
        return pontuacoes.get(chave);
    }
}
