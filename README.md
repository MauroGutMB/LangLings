# LangLings

Uma CLI de exercícios de programação multi-linguagem, no espírito do Rustlings,
com interface de terminal e cada linguagem rodando isolada em seu próprio
container Docker.

## Descrição

Você escolhe uma atividade na interface. O código-base é copiado para um
diretório no seu computador e a tela mostra o caminho. Você abre o arquivo no
seu editor de sempre e trabalha nele normalmente. **Ao salvar, a validação
dispara sozinha** — não existe comando de "verificar" para lembrar. O código
roda dentro do container daquela linguagem, e o progresso vai para um banco
SQLite: 🔴 não iniciado, 🟡 em progresso, 🟢 completo.

Cada linguagem tem cinco categorias de atividade:

| Categoria | Como é validada |
|---|---|
| **Básico** | testes automatizados sobre o seu código |
| **Sintaxe** | testes automatizados sobre o seu código |
| **Compilador/Interpretador** | você compila à mão no shell do container, e o CLI confere o resultado |
| **Frameworks/Bibliotecas** | testes automatizados |
| **Exemplos reais** | testes automatizados sobre projetos multi-arquivo |

**Básico** é a porta de entrada de uma linguagem: o arquivo já vem com os
exemplos de uso comentados e rodando, e fecha com uma lacuna pequena, análoga ao
que você acabou de ler. Você aprende lendo código que funciona, não decorando
sintaxe no vazio. **Sintaxe** é o passo seguinte — a armadilha idiomática que só
aparece quando você já sabe escrever na linguagem.

A categoria **Compilador/Interpretador** é a que foge do padrão: não há testes.
Você aperta `[s]`, cai num shell dentro do container, compila do seu jeito, sai
com `exit`, e o LangLings verifica critérios — o binário existe? foi gerado
para a plataforma certa? o metadado que você injetou está lá?

Seu trabalho nunca é sobrescrito: reabrir uma atividade preserva o que você
escreveu. Restaurar o código-base é sempre uma ação explícita.

## Como executar

**Requisitos:** Linux ou WSL2, Go 1.24+, Docker acessível sem `sudo`
(seu usuário no grupo `docker`), e `gcc` — o driver de SQLite usa cgo.

```bash
go build -o langlings ./cmd/langlings
./langlings
```

Na primeira vez que você abrir uma linguagem ainda não instalada, o LangLings
baixa a imagem Docker dela e mostra o progresso na tela. Só as linguagens que
você usar ocupam espaço.

### Navegação

| Tecla | Ação |
|---|---|
| `↑` `↓` ou `k` `j` | navegar |
| `enter` | abrir (ou instalar a linguagem) |
| `s` | abrir shell no container, na atividade atual |
| `v` | validar agora |
| `h` | revelar uma dica |
| `esc` | voltar |
| `q` / `ctrl+c` | sair |

### Outros comandos

```bash
./langlings paths     # onde ficam workspace, banco e log
./langlings verify    # confere o catálogo de exercícios; exit != 0 se algo quebrou
./langlings reset -h  # zera o progresso (com -files, restaura os arquivos)
```

O comando procura `exercises/` e `languages/` a partir do diretório atual,
subindo pelos ancestrais. Use `-content <dir>` ou a variável `LANGLINGS_CONTENT`
para apontar outro lugar.

### Rodando os testes

```bash
go test -short ./...   # rápido: sem Docker, sem rede
go test -p 1 ./...     # completo: sobe containers de verdade
```

O `-p 1` importa: pacotes diferentes mexem no mesmo Docker.

## Linguagens

| Linguagem | Situação |
|---|---|
| Go | disponível |
| C | planejada |
| Rust | planejada |
| Java | planejada |
| Dart | planejada |
| JavaScript | planejada |
| TypeScript | planejada |
| Lua | planejada |
| Bash Script | planejada |

Adicionar uma linguagem é escrever um `languages/<slug>/language.toml` e as
atividades em `exercises/<slug>/<categoria>/<slug-da-atividade>/`. Cada
atividade tem um manifesto, um diretório `base/` (o que você recebe) e um
`solution/` (a referência). O comando `verify` confere que o `base/` realmente
reprova e que a `solution/` realmente passa — um exercício cujo código inicial
já está correto não é um exercício.

## Créditos

Projeto pessoal de **Mauro Gutemberg Magalhães Barros**.

Inspirado no [Rustlings](https://github.com/rust-lang/rustlings), que provou
que aprender uma linguagem corrigindo código quebrado funciona melhor que ler
sobre ela.

Construído com:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) e
  [Lip Gloss](https://github.com/charmbracelet/lipgloss) — interface de terminal
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — persistência do progresso
- [fsnotify](https://github.com/fsnotify/fsnotify) — detecção de alterações
- [BurntSushi/toml](https://github.com/BurntSushi/toml) — manifestos
- [testify](https://github.com/stretchr/testify) — testes
- [Docker](https://www.docker.com/) — isolamento por linguagem
