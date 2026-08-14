package catalog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"langlings/internal/domain"
)

// Nomes fixos do layout de conteúdo.
const (
	LanguagesDir    = "languages"
	ExercisesDir    = "exercises"
	LanguageFile    = "language.toml"
	ExerciseFile    = "exercise.toml"
	BaseDirName     = "base"
	SolutionDirName = "solution"
)

// Catalog é o conteúdo descoberto em disco: as linguagens e seus exercícios.
type Catalog struct {
	Root      string
	Languages []domain.Language
	Exercises []domain.Exercise
}

// Load lê languages/ e exercises/ a partir de root.
//
// Os erros são acumulados em vez de interrompidos no primeiro: quando algo
// está errado no conteúdo, é melhor ver os cinco manifestos quebrados de uma
// vez do que descobrir um por execução.
func Load(root string) (*Catalog, error) {
	c := &Catalog{Root: root}
	var problems []error

	langs, errs := loadLanguages(filepath.Join(root, LanguagesDir))
	c.Languages = langs
	problems = append(problems, errs...)

	known := make(map[string]bool, len(langs))
	for _, l := range langs {
		known[l.Slug] = true
	}

	exercises, errs := loadExercises(filepath.Join(root, ExercisesDir), known)
	c.Exercises = exercises
	problems = append(problems, errs...)

	return c, errors.Join(problems...)
}

func loadLanguages(dir string) ([]domain.Language, []error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, []error{fmt.Errorf("diretório de linguagens não encontrado: %s", dir)}
		}
		return nil, []error{err}
	}

	var langs []domain.Language
	var problems []error

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		slug := e.Name()
		path := filepath.Join(dir, slug, LanguageFile)
		if !fileExists(path) {
			problems = append(problems, fmt.Errorf("%s: %s ausente", filepath.Join(dir, slug), LanguageFile))
			continue
		}

		var m languageManifest
		if err := decodeFile(path, &m); err != nil {
			problems = append(problems, fmt.Errorf("%s: %w", path, err))
			continue
		}
		lang, err := m.toLanguage(slug)
		if err != nil {
			problems = append(problems, fmt.Errorf("%s: %w", path, err))
			continue
		}
		langs = append(langs, lang)
	}

	sort.Slice(langs, func(i, j int) bool { return langs[i].Slug < langs[j].Slug })
	return langs, problems
}

// loadExercises espera o layout exercises/<linguagem>/<categoria>/<slug>/.
func loadExercises(dir string, knownLanguages map[string]bool) ([]domain.Exercise, []error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, []error{fmt.Errorf("diretório de exercícios não encontrado: %s", dir)}
		}
		return nil, []error{err}
	}

	var exercises []domain.Exercise
	var problems []error

	for _, langEntry := range entries {
		if !langEntry.IsDir() {
			continue
		}
		language := langEntry.Name()
		if !knownLanguages[language] {
			problems = append(problems, fmt.Errorf(
				"exercises/%s: nenhuma linguagem %q declarada em %s/", language, language, LanguagesDir))
			continue
		}

		catEntries, err := os.ReadDir(filepath.Join(dir, language))
		if err != nil {
			problems = append(problems, err)
			continue
		}

		for _, catEntry := range catEntries {
			if !catEntry.IsDir() {
				continue
			}
			category := domain.Category(catEntry.Name())
			if !category.Valid() {
				problems = append(problems, fmt.Errorf(
					"exercises/%s/%s: categoria desconhecida (use: %s)",
					language, catEntry.Name(), domain.CategoriesHint()))
				continue
			}

			slugEntries, err := os.ReadDir(filepath.Join(dir, language, catEntry.Name()))
			if err != nil {
				problems = append(problems, err)
				continue
			}

			for _, slugEntry := range slugEntries {
				if !slugEntry.IsDir() {
					continue
				}
				id := filepath.ToSlash(filepath.Join(language, catEntry.Name(), slugEntry.Name()))
				exDir := filepath.Join(dir, language, catEntry.Name(), slugEntry.Name())

				ex, err := loadExercise(id, exDir, language, category)
				if err != nil {
					problems = append(problems, fmt.Errorf("%s: %w", id, err))
					continue
				}
				exercises = append(exercises, ex)
			}
		}
	}

	sortExercises(exercises)
	return exercises, problems
}

func loadExercise(id, dir, language string, category domain.Category) (domain.Exercise, error) {
	manifestPath := filepath.Join(dir, ExerciseFile)
	if !fileExists(manifestPath) {
		return domain.Exercise{}, fmt.Errorf("%s ausente", ExerciseFile)
	}

	var m exerciseManifest
	if err := decodeFile(manifestPath, &m); err != nil {
		return domain.Exercise{}, fmt.Errorf("%s: %w", ExerciseFile, err)
	}

	ex, err := m.toExercise(id, dir, language, category)
	if err != nil {
		return domain.Exercise{}, err
	}

	// base/ é o que vai para o workspace do usuário: sem ele não há exercício.
	if !dirExists(ex.BaseDir()) {
		return domain.Exercise{}, fmt.Errorf("diretório %s/ ausente", BaseDirName)
	}
	for _, f := range ex.Editable {
		if !fileExists(filepath.Join(ex.BaseDir(), f)) {
			return domain.Exercise{}, fmt.Errorf("editable %q não existe em %s/", f, BaseDirName)
		}
	}

	return ex, nil
}

// sortExercises ordena por linguagem, depois pela ordem canônica das
// categorias (não alfabética), depois pelo slug. É a ordem exibida na TUI.
func sortExercises(exercises []domain.Exercise) {
	rank := make(map[domain.Category]int, len(domain.Categories))
	for i, c := range domain.Categories {
		rank[c] = i
	}
	sort.Slice(exercises, func(i, j int) bool {
		a, b := exercises[i], exercises[j]
		if a.Language != b.Language {
			return a.Language < b.Language
		}
		if a.Category != b.Category {
			return rank[a.Category] < rank[b.Category]
		}
		return a.Path < b.Path
	})
}

// Language devolve a linguagem de slug informado.
func (c *Catalog) Language(slug string) (domain.Language, bool) {
	for _, l := range c.Languages {
		if l.Slug == slug {
			return l, true
		}
	}
	return domain.Language{}, false
}

// Exercise devolve o exercício de path informado.
func (c *Catalog) Exercise(path string) (domain.Exercise, bool) {
	for _, e := range c.Exercises {
		if e.Path == path {
			return e, true
		}
	}
	return domain.Exercise{}, false
}

// ByLanguage devolve os exercícios de uma linguagem, na ordem de exibição.
func (c *Catalog) ByLanguage(slug string) []domain.Exercise {
	var out []domain.Exercise
	for _, e := range c.Exercises {
		if e.Language == slug {
			out = append(out, e)
		}
	}
	return out
}
