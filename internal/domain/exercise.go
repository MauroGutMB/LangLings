// Package domain contém as entidades e as regras do LangLings.
//
// Este pacote não importa nada de infraestrutura: sem os/exec, sem database/sql,
// sem bubbletea. Tudo aqui é testável sem Docker, sem banco e sem terminal. Se
// algum dia um mock for necessário para testar este pacote, é sinal de que
// infraestrutura vazou para dentro dele.
package domain

import (
	"fmt"
	"path"
	"strings"
	"time"
)

// Category é uma das quatro categorias de atividade de uma linguagem.
type Category string

const (
	CategorySintaxe    Category = "sintaxe"
	CategoryCompilador Category = "compilador"
	CategoryFrameworks Category = "frameworks"
	CategoryExemplos   Category = "exemplos"
)

// Categories lista as categorias na ordem em que devem aparecer na TUI.
var Categories = []Category{
	CategorySintaxe,
	CategoryCompilador,
	CategoryFrameworks,
	CategoryExemplos,
}

// Label devolve o nome de exibição da categoria.
func (c Category) Label() string {
	switch c {
	case CategorySintaxe:
		return "Sintaxe"
	case CategoryCompilador:
		return "Compilador/Interpretador"
	case CategoryFrameworks:
		return "Frameworks/Bibliotecas"
	case CategoryExemplos:
		return "Exemplos reais"
	default:
		return string(c)
	}
}

// Valid informa se a categoria é uma das quatro conhecidas.
func (c Category) Valid() bool {
	for _, known := range Categories {
		if c == known {
			return true
		}
	}
	return false
}

// Language é uma linguagem-alvo, descrita por languages/<slug>/language.toml.
//
// Note que não há campo "instalada": isso é derivado ao vivo de
// `docker image inspect`, nunca persistido, para não divergir da realidade.
type Language struct {
	Slug       string   // "go", derivado do nome do diretório
	Name       string   // "Go"
	Image      string   // imagem pronta, ex: golang:1.26-alpine
	Dockerfile string   // alternativa a Image: caminho relativo de um Dockerfile
	Workdir    string   // diretório de trabalho no container
	Shell      []string // comando do shell interativo

	// Env são variáveis passadas a toda execução desta linguagem. É por aqui
	// que o cache do toolchain é apontado para dentro de CacheDir, sem que o
	// runner precise saber o que é GOCACHE, CARGO_HOME ou MAVEN_OPTS.
	Env map[string]string

	// CacheDir é um diretório dentro do container montado como volume nomeado
	// e persistente entre execuções.
	//
	// Não é otimização de conforto: sem ele, cada container novo recompila a
	// biblioteca padrão do zero. Medido em Go, isso é a diferença entre 12s e
	// 400ms por validação — o que decide se o loop de save é utilizável.
	CacheDir string
}

// CacheVolume é o nome do volume Docker que guarda o cache desta linguagem.
func (l Language) CacheVolume() string { return "langlings-cache-" + l.Slug }

// ImageRef é o nome da imagem usada para esta linguagem. Quando a linguagem
// tem Dockerfile próprio, a imagem construída recebe um nome derivado do slug.
func (l Language) ImageRef() string {
	if l.Image != "" {
		return l.Image
	}
	return "langlings/" + l.Slug + ":latest"
}

// BuildsOwnImage informa se a imagem precisa ser construída em vez de baixada.
func (l Language) BuildsOwnImage() bool { return l.Image == "" && l.Dockerfile != "" }

// Exercise é uma atividade, descrita por um exercise.toml.
type Exercise struct {
	// Path é a identidade do exercício e a chave primária no banco:
	// o caminho relativo à raiz de exercises/, ex: "go/sintaxe/slices-append".
	Path string

	// Dir é o diretório absoluto do exercício no repositório de conteúdo.
	// Não faz parte da identidade — é onde o conteúdo mora agora.
	Dir string

	Language  string
	Category  Category
	Title     string
	Objective string

	// Editable é a allowlist de arquivos que o usuário edita, relativa à raiz
	// do exercício. Só mudanças nestes arquivos disparam validação; tudo o
	// mais no workspace é artefato de build.
	Editable []string

	Hints      []string
	Validation ValidationSpec
}

// BaseDir e SolutionDir são o layout fixo de um exercício em disco.
func (e Exercise) BaseDir() string     { return path.Join(e.Dir, "base") }
func (e Exercise) SolutionDir() string { return path.Join(e.Dir, "solution") }

// ValidationMode distingue os dois modos de validação, que é a bifurcação
// central do projeto.
type ValidationMode string

const (
	// ModeTest roda um comando de teste sobre o código. Exit 0 é aprovação.
	// Usado por Sintaxe, Frameworks e Exemplos reais.
	ModeTest ValidationMode = "test"

	// ModeCriteria não roda testes: observa efeitos colaterais que o usuário
	// produziu manualmente no shell do container. Usado por Compilador.
	ModeCriteria ValidationMode = "criteria"
)

// ValidationSpec descreve como decidir se um exercício foi resolvido.
type ValidationSpec struct {
	Mode     ValidationMode
	Command  []string    // ModeTest
	Criteria []Criterion // ModeCriteria
	Timeout  time.Duration

	// Network libera a rede para este exercício. Como limites de rede são
	// fixados na criação do container, um exercício com Network=true não pode
	// ser servido pela sessão isolada e roda num container efêmero à parte.
	Network bool
}

// DefaultTimeout é usado quando o manifesto não declara um.
const DefaultTimeout = 30 * time.Second

// Validate confere a coerência interna do exercício. É chamado no parsing do
// manifesto, para que um exercício malformado falhe no boot com mensagem clara
// em vez de falhar de modo confuso durante uma validação.
func (e Exercise) Validate() error {
	if strings.TrimSpace(e.Title) == "" {
		return fmt.Errorf("title é obrigatório")
	}
	if strings.TrimSpace(e.Objective) == "" {
		return fmt.Errorf("objective é obrigatório")
	}
	if !e.Category.Valid() {
		return fmt.Errorf("category %q desconhecida (use: sintaxe, compilador, frameworks, exemplos)", e.Category)
	}
	if len(e.Editable) == 0 {
		return fmt.Errorf("editable é obrigatório: sem ele o watcher não sabe o que observar")
	}
	for _, f := range e.Editable {
		if err := validateRelPath(f); err != nil {
			return fmt.Errorf("editable %q: %w", f, err)
		}
	}
	return e.Validation.Validate()
}

// Validate confere a coerência da especificação de validação.
func (v ValidationSpec) Validate() error {
	switch v.Mode {
	case ModeTest:
		if len(v.Command) == 0 {
			return fmt.Errorf("validation.command é obrigatório quando mode = \"test\"")
		}
		if len(v.Criteria) > 0 {
			return fmt.Errorf("validation.criteria não se aplica quando mode = \"test\"")
		}
	case ModeCriteria:
		if len(v.Criteria) == 0 {
			return fmt.Errorf("mode = \"criteria\" exige pelo menos um [[validation.criteria]]")
		}
		if len(v.Command) > 0 {
			return fmt.Errorf("validation.command não se aplica quando mode = \"criteria\"")
		}
		for i, c := range v.Criteria {
			if err := c.Validate(); err != nil {
				return fmt.Errorf("criteria[%d]: %w", i, err)
			}
		}
	default:
		return fmt.Errorf("validation.mode %q desconhecido (use \"test\" ou \"criteria\")", v.Mode)
	}
	if v.Timeout <= 0 {
		return fmt.Errorf("validation.timeout deve ser positivo")
	}
	return nil
}

// validateRelPath rejeita caminhos que escapariam do diretório do exercício.
// Um "../" aqui daria a um manifesto acesso de leitura e escrita ao resto do
// workspace, então a checagem não é cosmética.
func validateRelPath(p string) error {
	if p == "" {
		return fmt.Errorf("caminho vazio")
	}
	if path.IsAbs(p) || strings.HasPrefix(p, "/") {
		return fmt.Errorf("deve ser relativo ao diretório do exercício")
	}
	clean := path.Clean(p)
	if clean == ".." || strings.HasPrefix(clean, "../") {
		return fmt.Errorf("não pode escapar do diretório do exercício")
	}
	return nil
}
