package engine

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"langlings/internal/domain"
)

// Paths reúne os diretórios do LangLings, seguindo XDG.
//
// A separação entre conteúdo e trabalho é deliberada: ContentRoot é o
// repositório versionado (manifestos, código-base, soluções) e Workspace é
// onde suas soluções vivem. Um reset destrutivo apaga o segundo e nunca toca
// no primeiro.
type Paths struct {
	ContentRoot string // repositório com exercises/ e languages/
	DataDir     string // ~/.local/share/langlings
	Workspace   string // DataDir/workspace
	DBPath      string // DataDir/langlings.db
	LogPath     string // ~/.local/state/langlings/langlings.log
}

// DefaultPaths monta os caminhos padrão a partir do ambiente.
func DefaultPaths(contentRoot string) (Paths, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Paths{}, err
	}

	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = filepath.Join(home, ".local", "share")
	}
	stateHome := os.Getenv("XDG_STATE_HOME")
	if stateHome == "" {
		stateHome = filepath.Join(home, ".local", "state")
	}

	dataDir := filepath.Join(dataHome, "langlings")
	return Paths{
		ContentRoot: contentRoot,
		DataDir:     dataDir,
		Workspace:   filepath.Join(dataDir, "workspace"),
		DBPath:      filepath.Join(dataDir, "langlings.db"),
		LogPath:     filepath.Join(stateHome, "langlings", "langlings.log"),
	}, nil
}

// EnsureDirs cria os diretórios necessários.
func (p Paths) EnsureDirs() error {
	for _, dir := range []string{p.DataDir, p.Workspace, filepath.Dir(p.LogPath)} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("criando %s: %w", dir, err)
		}
	}
	return nil
}

// ExerciseDir é onde o usuário edita um exercício, no host.
func (p Paths) ExerciseDir(ex domain.Exercise) string {
	return filepath.Join(p.Workspace, filepath.FromSlash(ex.Path))
}

// Materialize garante que o exercício exista no workspace do usuário.
//
// Se o diretório já existe, o conteúdo é preservado sem exceção: a solução em
// andamento do usuário nunca pode ser sobrescrita por um retorno à tela do
// exercício. Sobrescrever é operação explícita, feita por Restore.
func (p Paths) Materialize(ex domain.Exercise) (string, error) {
	dest := p.ExerciseDir(ex)

	if info, err := os.Stat(dest); err == nil && info.IsDir() {
		return dest, nil
	}

	if err := copyDir(ex.BaseDir(), dest); err != nil {
		return "", fmt.Errorf("materializando %s: %w", ex.Path, err)
	}
	return dest, nil
}

// Restore devolve o exercício ao código-base, descartando o que o usuário
// escreveu. É a metade destrutiva do reset e só deve ser chamada após
// confirmação explícita.
func (p Paths) Restore(ex domain.Exercise) (string, error) {
	dest := p.ExerciseDir(ex)

	if err := os.RemoveAll(dest); err != nil {
		return "", fmt.Errorf("limpando %s: %w", dest, err)
	}
	if err := copyDir(ex.BaseDir(), dest); err != nil {
		return "", fmt.Errorf("restaurando %s: %w", ex.Path, err)
	}
	return dest, nil
}

// ReadEditable lê os arquivos declarados como editáveis a partir de um
// workspace. Arquivos ausentes simplesmente não aparecem no resultado, o que
// já muda o hash — apagar um arquivo é uma mudança como outra qualquer.
func ReadEditable(dir string, ex domain.Exercise) (map[string][]byte, error) {
	out := make(map[string][]byte, len(ex.Editable))

	for _, rel := range ex.Editable {
		content, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(rel)))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		out[rel] = content
	}
	return out, nil
}

// copyDir copia uma árvore de diretórios preservando o modo dos arquivos.
func copyDir(src, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s não é um diretório", src)
	}

	return filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dest, rel)

		if fi.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if !fi.Mode().IsRegular() {
			// Links simbólicos e afins não têm lugar num código-base de
			// exercício; ignorar é mais seguro que reproduzir.
			return nil
		}
		return copyFile(path, target, fi.Mode())
	})
}

func copyFile(src, dest string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
