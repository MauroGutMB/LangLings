# LangLings

Uma CLI de exercícios de programação multi-linguagem, no espírito do Rustlings,
com interface de terminal e cada linguagem rodando isolada em seu próprio
container Docker. São 199 atividades em 9 linguagens: Go, C, Rust, Java, Dart,
JavaScript, TypeScript, Lua e Bash.

## Descrição

Você escolhe uma atividade na interface. O código-base é copiado para um
diretório no seu computador e a tela mostra o caminho. Você abre o arquivo no
seu editor de sempre e trabalha nele normalmente. **Ao salvar, a validação
dispara sozinha** — não existe comando de "verificar" para lembrar. O código
roda dentro do container daquela linguagem, e o progresso vai para um banco
SQLite: 🔴 não iniciado, 🟡 em progresso, 🟢 completo.

As atividades se dividem em cinco categorias:

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

**Requisitos:** Linux ou WSL2, Go 1.26+, Docker acessível sem `sudo`
(seu usuário no grupo `docker`), e `gcc` — o driver de SQLite usa cgo.

```bash
go build -o langlings ./cmd/langlings
./langlings
```

Na primeira vez que você abrir uma linguagem ainda não instalada, o LangLings
baixa a imagem Docker dela e mostra o progresso na tela. Só as linguagens que
você usar ocupam espaço.

### Onde você edita

O LangLings não é um editor e não roda dentro do container — só a validação
roda. Você trabalha nos arquivos com o seu editor de sempre, em:

```
~/.local/share/langlings/workspace/<linguagem>/<categoria>/<atividade>/
```

A tela da atividade mostra esse caminho, e `./langlings paths` também. Abra o
arquivo em outra janela, deixe a TUI visível ao lado, e salve: a validação
dispara e o estado muda na hora.

```bash
vim ~/.local/share/langlings/workspace/go/basico/07-maps/main.go
```

A tecla `s` abre um shell dentro do container, mas ele existe para *inspecionar*
o ambiente — rodar o compilador à mão, examinar um binário —, não para editar.
Enquanto o shell está aberto a interface fica suspensa, então você não vê a
validação acontecendo.

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

São **199 atividades**, distribuídas igualmente entre as nove linguagens.

| Linguagem | Básico | Sintaxe | Compilador | Frameworks | Exemplos | Imagem |
|---|---|---|---|---|---|---|
| Go | 10 | 12 | 1 | — | — | `golang:1.26-alpine` |
| C | 10 | 12 | — | — | — | Dockerfile próprio |
| Rust | 10 | 12 | — | — | — | `rust:1.97-slim-trixie` |
| Java | 10 | 12 | — | — | — | `eclipse-temurin:25-jdk-alpine` |
| Dart | 10 | 12 | — | — | — | `dart:3.13-sdk` |
| JavaScript | 10 | 12 | — | — | — | `node:24-alpine` |
| TypeScript | 10 | 12 | — | — | — | Dockerfile próprio |
| Lua | 10 | 12 | — | — | — | Dockerfile próprio |
| Bash | 10 | 12 | — | — | — | `bash:5.2-alpine3.22` |

Nenhuma atividade usa dependência externa: todas rodam com a rede desligada,
usando só o que vem no toolchain de cada linguagem. As três que constroem a
própria imagem têm motivo específico — Lua porque não existe imagem oficial no
Docker Hub, TypeScript porque o Node executa `.ts` apagando os tipos sem
checá-los (sem o `tsc` na imagem, um exercício de TypeScript aprovaria código
com erro de tipo), e C para não carregar os 561 MB da `gcc` oficial.

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
