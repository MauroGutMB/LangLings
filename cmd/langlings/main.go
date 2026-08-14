// Comando langlings.
//
// Este arquivo é só fiação: descobre onde está o conteúdo, abre o banco, monta
// o engine e entrega tudo à TUI ou a um subcomando. Nenhuma regra de negócio
// mora aqui — se alguma aparecer, ela pertence a internal/.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"

	"langlings/internal/catalog"
	"langlings/internal/domain"
	"langlings/internal/engine"
	"langlings/internal/runner"
	"langlings/internal/store"
	"langlings/internal/tui"
)

const uso = `langlings — exercícios de programação multi-linguagem

uso:
  langlings                    abre a interface
  langlings verify [caminho…]  confere o conteúdo (sem argumento: tudo)
  langlings reset [opções]     zera o progresso
  langlings paths              mostra onde ficam os arquivos

opções globais:
  -content <dir>   raiz do conteúdo (exercises/ e languages/)
                   padrão: $LANGLINGS_CONTENT, ou o diretório atual

reset:
  -exercise <path> zera um exercício
  -language <slug> zera uma linguagem
  -all             zera tudo
  -files           também restaura os arquivos ao código-base (destrutivo)
`

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "langlings: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]

	// O subcomando, quando existe, é o primeiro argumento que não é flag.
	comando := ""
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		comando, args = args[0], args[1:]
	}

	switch comando {
	case "":
		return comandoTUI(args)
	case "verify":
		return comandoVerify(args)
	case "reset":
		return comandoReset(args)
	case "paths":
		return comandoPaths(args)
	case "help", "-h", "--help":
		fmt.Print(uso)
		return nil
	default:
		fmt.Print(uso)
		return fmt.Errorf("subcomando desconhecido: %q", comando)
	}
}

// ---------- montagem ----------

type app struct {
	engine *engine.Engine
	store  *store.Store
}

func (a *app) Close() {
	if a.store != nil {
		a.store.Close()
	}
}

// montar faz o trabalho comum a todos os subcomandos.
func montar(ctx context.Context, contentRoot string) (*app, error) {
	raiz, err := descobrirConteudo(contentRoot)
	if err != nil {
		return nil, err
	}

	cat, err := catalog.Load(raiz)
	if err != nil {
		// Erros de manifesto são acumulados: melhor ver os cinco problemas de
		// uma vez do que descobrir um por execução.
		return nil, fmt.Errorf("conteúdo inválido em %s:\n%s", raiz, indentar(err.Error()))
	}

	paths, err := engine.DefaultPaths(raiz)
	if err != nil {
		return nil, err
	}
	if err := paths.EnsureDirs(); err != nil {
		return nil, err
	}

	st, err := store.Open(paths.DBPath)
	if err != nil {
		return nil, err
	}

	d := runner.NewDocker()
	if err := d.Available(ctx); err != nil {
		st.Close()
		return nil, fmt.Errorf("%w\n\nO LangLings roda os exercícios em containers. Verifique se o Docker está ativo.", err)
	}

	e := engine.New(cat, st, d, paths)
	if err := e.Bootstrap(ctx); err != nil {
		st.Close()
		return nil, err
	}

	return &app{engine: e, store: st}, nil
}

// descobrirConteudo localiza a raiz com exercises/ e languages/.
func descobrirConteudo(flagValue string) (string, error) {
	candidatos := []string{flagValue, os.Getenv("LANGLINGS_CONTENT")}

	if wd, err := os.Getwd(); err == nil {
		candidatos = append(candidatos, wd)
	}
	// O binário costuma morar ao lado do conteúdo numa instalação pessoal.
	if exe, err := os.Executable(); err == nil {
		candidatos = append(candidatos, filepath.Dir(exe))
	}

	for _, c := range candidatos {
		if c == "" {
			continue
		}
		if raiz, ok := subirAte(c); ok {
			return raiz, nil
		}
	}

	return "", errors.New("não encontrei o conteúdo (exercises/ e languages/).\n" +
		"Rode a partir da raiz do projeto, use -content <dir> ou defina LANGLINGS_CONTENT")
}

// subirAte procura o conteúdo no diretório e nos ancestrais dele, para que
// rodar o comando de dentro de uma subpasta do projeto funcione.
func subirAte(dir string) (string, bool) {
	atual, err := filepath.Abs(dir)
	if err != nil {
		return "", false
	}

	for {
		if temConteudo(atual) {
			return atual, true
		}
		pai := filepath.Dir(atual)
		if pai == atual {
			return "", false
		}
		atual = pai
	}
}

func temConteudo(dir string) bool {
	for _, sub := range []string{catalog.ExercisesDir, catalog.LanguagesDir} {
		info, err := os.Stat(filepath.Join(dir, sub))
		if err != nil || !info.IsDir() {
			return false
		}
	}
	return true
}

// ---------- subcomandos ----------

func comandoTUI(args []string) error {
	fs := flag.NewFlagSet("langlings", flag.ContinueOnError)
	content := fs.String("content", "", "raiz do conteúdo")
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Ctrl+C e SIGTERM precisam derrubar os containers, não deixá-los para trás.
	ctx, parar := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer parar()

	a, err := montar(ctx, *content)
	if err != nil {
		return err
	}
	defer a.Close()
	defer a.engine.Shutdown(context.Background())

	modelo := tui.New(ctx, a.engine)
	p := tea.NewProgram(modelo, tea.WithAltScreen(), tea.WithContext(ctx))
	modelo.Attach(p)

	_, err = p.Run()
	if err != nil && !errors.Is(err, tea.ErrProgramKilled) && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// comandoVerify é o gate de conteúdo: transforma auditoria manual de dezenas
// de diretórios num comando com exit code.
func comandoVerify(args []string) error {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	content := fs.String("content", "", "raiz do conteúdo")
	if err := fs.Parse(args); err != nil {
		return err
	}

	ctx := context.Background()
	a, err := montar(ctx, *content)
	if err != nil {
		return err
	}
	defer a.Close()
	defer a.engine.Shutdown(context.Background())

	alvos := fs.Args()

	exercicios := a.engine.Catalog.Exercises
	if len(alvos) > 0 {
		exercicios = nil
		for _, alvo := range alvos {
			ex, ok := a.engine.Catalog.Exercise(alvo)
			if !ok {
				return fmt.Errorf("exercício %q não está no catálogo", alvo)
			}
			exercicios = append(exercicios, ex)
		}
	}

	if err := garantirImagens(ctx, a, exercicios); err != nil {
		return err
	}

	relatorios := make([]engine.VerifyReport, 0, len(exercicios))
	for _, ex := range exercicios {
		relatorios = append(relatorios, a.engine.Verify(ctx, ex))
	}

	reprovados := 0
	for _, r := range relatorios {
		if r.OK() {
			fmt.Printf("  ok    %s\n", r.Exercise.Path)
			continue
		}
		reprovados++
		fmt.Printf("  FALHA %s\n", r.Exercise.Path)
		for _, p := range r.Problems {
			fmt.Printf("        · %s\n", strings.ReplaceAll(p, "\n", "\n          "))
		}
	}

	fmt.Printf("\n%d exercício(s), %d com problema\n", len(relatorios), reprovados)
	if reprovados > 0 {
		return fmt.Errorf("o conteúdo não passou no gate")
	}
	return nil
}

// garantirImagens materializa a imagem de cada linguagem envolvida antes de
// verificar qualquer coisa.
//
// A TUI faz isso pelo fluxo de instalação, mas o verify vai direto ao container
// efêmero. Para uma imagem oficial o daemon puxa sozinho e ninguém percebe a
// falta; para uma linguagem que constrói a própria imagem, o docker tentaria
// baixar `langlings/<slug>` do Hub e falharia com "pull access denied" — um erro
// que não tem nada a ver com o problema real.
func garantirImagens(ctx context.Context, a *app, exercicios []domain.Exercise) error {
	vistos := map[string]bool{}

	for _, ex := range exercicios {
		if vistos[ex.Language] {
			continue
		}
		vistos[ex.Language] = true

		lang, ok := a.engine.Catalog.Language(ex.Language)
		if !ok {
			return fmt.Errorf("linguagem %q não encontrada", ex.Language)
		}
		// O log do build vai para stderr: num gate rodado em CI, a saída limpa
		// do stdout continua sendo só o relatório dos exercícios.
		if err := a.engine.Install(ctx, lang, os.Stderr); err != nil {
			return err
		}
	}
	return nil
}

func comandoReset(args []string) error {
	fs := flag.NewFlagSet("reset", flag.ContinueOnError)
	content := fs.String("content", "", "raiz do conteúdo")
	exercicio := fs.String("exercise", "", "caminho do exercício")
	linguagem := fs.String("language", "", "slug da linguagem")
	tudo := fs.Bool("all", false, "zerar tudo")
	arquivos := fs.Bool("files", false, "também restaurar os arquivos ao código-base")
	if err := fs.Parse(args); err != nil {
		return err
	}

	escolhas := 0
	for _, escolhido := range []bool{*exercicio != "", *linguagem != "", *tudo} {
		if escolhido {
			escolhas++
		}
	}
	if escolhas != 1 {
		return errors.New("escolha exatamente um: -exercise, -language ou -all")
	}

	ctx := context.Background()
	a, err := montar(ctx, *content)
	if err != nil {
		return err
	}
	defer a.Close()

	modo := engine.ResetEstado
	if *arquivos {
		modo = engine.ResetCompleto
		// Apagar o trabalho do usuário nunca acontece sem ele dizer sim.
		if !confirmar("Isto apaga o que você escreveu e restaura o código-base.") {
			fmt.Println("cancelado")
			return nil
		}
	}

	switch {
	case *exercicio != "":
		ex, ok := a.engine.Catalog.Exercise(*exercicio)
		if !ok {
			return fmt.Errorf("exercício %q não está no catálogo", *exercicio)
		}
		if err := a.engine.ResetExercise(ex, modo); err != nil {
			return err
		}
		fmt.Printf("%s zerado\n", ex.Path)

	case *linguagem != "":
		n, err := a.engine.ResetLanguage(*linguagem, modo)
		if err != nil {
			return err
		}
		fmt.Printf("%d exercício(s) de %s zerado(s)\n", n, *linguagem)

	default:
		n, err := a.engine.ResetAll(modo)
		if err != nil {
			return err
		}
		fmt.Printf("%d exercício(s) zerado(s)\n", n)
	}
	return nil
}

func comandoPaths(args []string) error {
	fs := flag.NewFlagSet("paths", flag.ContinueOnError)
	content := fs.String("content", "", "raiz do conteúdo")
	if err := fs.Parse(args); err != nil {
		return err
	}

	raiz, err := descobrirConteudo(*content)
	if err != nil {
		return err
	}
	p, err := engine.DefaultPaths(raiz)
	if err != nil {
		return err
	}

	fmt.Printf("conteúdo   %s\n", p.ContentRoot)
	fmt.Printf("workspace  %s\n", p.Workspace)
	fmt.Printf("banco      %s\n", p.DBPath)
	fmt.Printf("log        %s\n", p.LogPath)
	return nil
}

func confirmar(aviso string) bool {
	fmt.Printf("%s\nDigite \"sim\" para continuar: ", aviso)
	var resposta string
	fmt.Scanln(&resposta)
	return strings.TrimSpace(strings.ToLower(resposta)) == "sim"
}

func indentar(s string) string {
	linhas := strings.Split(s, "\n")
	for i, l := range linhas {
		linhas[i] = "  " + l
	}
	return strings.Join(linhas, "\n")
}
