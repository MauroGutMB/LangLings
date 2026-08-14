// Package catalog descobre e interpreta o conteúdo do LangLings: os
// manifestos TOML que descrevem linguagens e exercícios.
//
// Toda a validação acontece aqui, no boot. A intenção é que um manifesto
// malformado produza uma mensagem que cita o arquivo e o campo, em vez de uma
// falha confusa no meio de uma validação — o que importa especialmente para
// conteúdo escrito em lote.
package catalog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"

	"langlings/internal/domain"
)

// exerciseManifest espelha o exercise.toml. É um tipo separado das entidades
// de domínio de propósito: o formato do arquivo pode mudar sem arrastar o
// domínio junto, e a conversão entre os dois é onde a validação mora.
type exerciseManifest struct {
	Title      string             `toml:"title"`
	Category   string             `toml:"category"`
	Objective  string             `toml:"objective"`
	Editable   []string           `toml:"editable"`
	Hints      []string           `toml:"hints"`
	Validation validationManifest `toml:"validation"`
}

type validationManifest struct {
	Mode     string              `toml:"mode"`
	Command  []string            `toml:"command"`
	Timeout  string              `toml:"timeout"`
	Network  bool                `toml:"network"`
	Criteria []criterionManifest `toml:"criteria"`
}

type criterionManifest struct {
	Kind     string   `toml:"kind"`
	Describe string   `toml:"describe"`
	Path     string   `toml:"path"`
	Command  []string `toml:"command"`
	Code     int      `toml:"code"`
	Expected string   `toml:"expected"`
	Pattern  string   `toml:"pattern"`
}

// languageManifest espelha o language.toml.
type languageManifest struct {
	Name       string            `toml:"name"`
	Image      string            `toml:"image"`
	Dockerfile string            `toml:"dockerfile"`
	Workdir    string            `toml:"workdir"`
	Shell      []string          `toml:"shell"`
	CacheDir   string            `toml:"cache_dir"`
	Env        map[string]string `toml:"env"`
}

// DefaultWorkdir é o diretório de trabalho dentro do container quando o
// manifesto da linguagem não declara um.
const DefaultWorkdir = "/workspace"

// decodeFile lê um TOML para dest e rejeita chaves desconhecidas.
//
// A rejeição não é preciosismo: sem ela, um `timeouts = "30s"` com "s" a mais
// seria silenciosamente ignorado e o exercício usaria o timeout padrão sem que
// ninguém percebesse. Erro de digitação em manifesto tem que doer no boot.
func decodeFile(path string, dest any) error {
	md, err := toml.DecodeFile(path, dest)
	if err != nil {
		return err
	}
	if undecoded := md.Undecoded(); len(undecoded) > 0 {
		keys := make([]string, 0, len(undecoded))
		for _, k := range undecoded {
			keys = append(keys, k.String())
		}
		sort.Strings(keys)
		return fmt.Errorf("chave(s) desconhecida(s): %s", strings.Join(keys, ", "))
	}
	return nil
}

// toExercise converte o manifesto em entidade de domínio e valida o resultado.
//
// id é o caminho relativo à raiz de exercises/ (a identidade do exercício), e
// dir é onde ele mora agora. A categoria vem do diretório, não do manifesto:
// o disco é a fonte da verdade. Se o manifesto também declarar uma categoria,
// ela precisa bater — divergência aí é quase sempre um copiar-e-colar mal
// terminado, exatamente o defeito que conteúdo gerado em lote produz.
func (m exerciseManifest) toExercise(id, dir, language string, category domain.Category) (domain.Exercise, error) {
	if m.Category != "" && domain.Category(m.Category) != category {
		return domain.Exercise{}, fmt.Errorf(
			"category = %q no manifesto, mas o exercício está no diretório %q — o diretório manda",
			m.Category, category)
	}

	timeout := domain.DefaultTimeout
	if m.Validation.Timeout != "" {
		parsed, err := time.ParseDuration(m.Validation.Timeout)
		if err != nil {
			return domain.Exercise{}, fmt.Errorf("validation.timeout %q inválido: %w", m.Validation.Timeout, err)
		}
		timeout = parsed
	}

	criteria := make([]domain.Criterion, 0, len(m.Validation.Criteria))
	for _, c := range m.Validation.Criteria {
		criteria = append(criteria, domain.Criterion{
			Kind:     domain.CriterionKind(c.Kind),
			Describe: c.Describe,
			Path:     c.Path,
			Command:  c.Command,
			Code:     c.Code,
			Expected: c.Expected,
			Pattern:  c.Pattern,
		})
	}

	ex := domain.Exercise{
		Path:      id,
		Dir:       dir,
		Language:  language,
		Category:  category,
		Title:     m.Title,
		Objective: strings.TrimSpace(m.Objective),
		Editable:  m.Editable,
		Hints:     m.Hints,
		Validation: domain.ValidationSpec{
			Mode:     domain.ValidationMode(m.Validation.Mode),
			Command:  m.Validation.Command,
			Criteria: criteria,
			Timeout:  timeout,
			Network:  m.Validation.Network,
		},
	}

	if err := ex.Validate(); err != nil {
		return domain.Exercise{}, err
	}
	return ex, nil
}

// toLanguage converte o manifesto em entidade de domínio e valida o resultado.
func (m languageManifest) toLanguage(slug string) (domain.Language, error) {
	if strings.TrimSpace(m.Name) == "" {
		return domain.Language{}, fmt.Errorf("name é obrigatório")
	}
	switch {
	case m.Image == "" && m.Dockerfile == "":
		return domain.Language{}, fmt.Errorf("declare image (imagem pronta) ou dockerfile (build local)")
	case m.Image != "" && m.Dockerfile != "":
		return domain.Language{}, fmt.Errorf("declare image ou dockerfile, não os dois")
	}

	// O dockerfile é resolvido contra a raiz do conteúdo na hora do build. Um
	// caminho absoluto ou com ".." apontaria para fora dela — e um manifesto
	// não deveria conseguir mandar o daemon construir a partir de qualquer
	// lugar do disco.
	if m.Dockerfile != "" {
		limpo := path.Clean(filepath.ToSlash(m.Dockerfile))
		switch {
		case path.IsAbs(limpo):
			return domain.Language{}, fmt.Errorf("dockerfile %q deve ser relativo à raiz do conteúdo", m.Dockerfile)
		case limpo == ".." || strings.HasPrefix(limpo, "../"):
			return domain.Language{}, fmt.Errorf("dockerfile %q não pode escapar da raiz do conteúdo", m.Dockerfile)
		}
	}

	workdir := m.Workdir
	if workdir == "" {
		workdir = DefaultWorkdir
	}
	if !strings.HasPrefix(workdir, "/") {
		return domain.Language{}, fmt.Errorf("workdir %q deve ser absoluto dentro do container", workdir)
	}

	shell := m.Shell
	if len(shell) == 0 {
		shell = []string{"/bin/sh"}
	}

	if m.CacheDir != "" && !strings.HasPrefix(m.CacheDir, "/") {
		return domain.Language{}, fmt.Errorf("cache_dir %q deve ser absoluto dentro do container", m.CacheDir)
	}
	if m.CacheDir != "" && m.CacheDir == workdir {
		return domain.Language{}, fmt.Errorf("cache_dir não pode ser o workdir: o volume de cache esconderia o código do usuário")
	}

	return domain.Language{
		Slug:       slug,
		Name:       m.Name,
		Image:      m.Image,
		Dockerfile: m.Dockerfile,
		Workdir:    workdir,
		Shell:      shell,
		CacheDir:   m.CacheDir,
		Env:        m.Env,
	}, nil
}

// fileExists é usado pelas checagens de layout do catálogo.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
