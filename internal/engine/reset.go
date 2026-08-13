package engine

import (
	"fmt"

	"langlings/internal/domain"
	"langlings/internal/store"
)

// ResetMode escolhe entre os dois níveis de reset.
type ResetMode int

const (
	// ResetEstado zera o progresso no banco e preserva os arquivos que o
	// usuário escreveu. É reversível na prática: o código continua lá.
	ResetEstado ResetMode = iota

	// ResetCompleto zera o progresso e devolve os arquivos ao código-base.
	// É destrutivo e exige confirmação explícita do usuário.
	ResetCompleto
)

// ResetExercise zera um exercício.
func (e *Engine) ResetExercise(ex domain.Exercise, mode ResetMode) error {
	if _, err := e.Store.Reset(store.ResetScope{Path: ex.Path}); err != nil {
		return err
	}
	if mode == ResetCompleto {
		if _, err := e.Paths.Restore(ex); err != nil {
			return err
		}
	}
	return nil
}

// ResetLanguage zera todos os exercícios de uma linguagem.
func (e *Engine) ResetLanguage(slug string, mode ResetMode) (int, error) {
	n, err := e.Store.Reset(store.ResetScope{Language: slug})
	if err != nil {
		return 0, err
	}

	if mode == ResetCompleto {
		for _, ex := range e.Catalog.ByLanguage(slug) {
			if _, err := e.Paths.Restore(ex); err != nil {
				return n, fmt.Errorf("restaurando %s: %w", ex.Path, err)
			}
		}
	}
	return n, nil
}

// ResetAll zera tudo.
func (e *Engine) ResetAll(mode ResetMode) (int, error) {
	n, err := e.Store.Reset(store.ResetScope{})
	if err != nil {
		return 0, err
	}

	if mode == ResetCompleto {
		for _, ex := range e.Catalog.Exercises {
			if _, err := e.Paths.Restore(ex); err != nil {
				return n, fmt.Errorf("restaurando %s: %w", ex.Path, err)
			}
		}
	}
	return n, nil
}
