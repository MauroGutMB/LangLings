import java.util.Objects;
import java.util.Set;

class Exercicio {

    static class Ponto {
        int x;
        int y;

        Ponto(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public boolean equals(Object outro) {
            if (!(outro instanceof Ponto)) {
                return false;
            }
            Ponto p = (Ponto) outro;
            return x == p.x && y == p.y;
        }

        // hashCode consistente com equals: pontos iguais por equals()
        // precisam ter o mesmo hashCode(), senão o HashSet não os acha.
        @Override
        public int hashCode() {
            return Objects.hash(x, y);
        }
    }

    // Devolva true quando pontos já contém um ponto equivalente a alvo.
    static boolean contido(Set<Ponto> pontos, Ponto alvo) {
        return pontos.contains(alvo);
    }
}
