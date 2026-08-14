class Exercicio {

    static class Pessoa {
        String nome;
        int idade;

        Pessoa(String nome, int idade) {
            this.nome = nome; // this distingue o campo do parâmetro de mesmo nome
            this.idade = idade;
        }

        void aniversario() {
            idade++;
        }
    }

    static void exemplos() {
        Pessoa ana = new Pessoa("ana", 30);
        System.out.println(ana.nome);   // ana
        System.out.println(ana.idade);  // 30

        ana.aniversario();
        System.out.println(ana.idade);  // 31

        Pessoa bruno = new Pessoa("bruno", 30);
        bruno.aniversario();
        System.out.println(ana.idade);   // 31 (inalterado)
        System.out.println(bruno.idade); // 31
    }

    static class Retangulo {
        double largura;
        double altura;

        Retangulo(double largura, double altura) {
            this.largura = largura;
            this.altura = altura;
        }

        double area() {
            return largura * altura;
        }

        // Perímetro do retângulo.
        double perimetro() {
            return 2 * (largura + altura);
        }
    }
}
