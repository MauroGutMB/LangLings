package store

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"langlings/internal/domain"
)

// abrir cria um banco em arquivo real (não :memory:) para exercitar o caminho
// de WAL, que é o que roda em produção.
func abrir(t *testing.T) *Store {
	t.Helper()
	s, err := Open(filepath.Join(t.TempDir(), "langlings.db"))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, s.Close()) })
	return s
}

func exercicios(paths ...string) []domain.Exercise {
	out := make([]domain.Exercise, 0, len(paths))
	for _, p := range paths {
		out = append(out, domain.Exercise{Path: p, Language: "go", Category: domain.CategorySintaxe})
	}
	return out
}

func TestOpen_AplicaMigrationsEEhIdempotente(t *testing.T) {
	path := filepath.Join(t.TempDir(), "langlings.db")

	s1, err := Open(path)
	require.NoError(t, err)
	require.NoError(t, s1.Reconcile(exercicios("go/sintaxe/a")))
	require.NoError(t, s1.Close())

	// Reabrir não deve reaplicar migrations nem perder dados.
	s2, err := Open(path)
	require.NoError(t, err)
	defer s2.Close()

	all, err := s2.All()
	require.NoError(t, err)
	require.Len(t, all, 1)
}

func TestGet_ExercicioDesconhecidoNaoEhErro(t *testing.T) {
	s := abrir(t)

	p, err := s.Get("go/sintaxe/nunca-visto")

	require.NoError(t, err)
	require.Equal(t, domain.StatusNotStarted, p.Status)
	require.Equal(t, "go/sintaxe/nunca-visto", p.Path)
}

func TestReconcile_InsereNovos(t *testing.T) {
	s := abrir(t)

	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a", "go/sintaxe/b")))

	all, err := s.All()
	require.NoError(t, err)
	require.Len(t, all, 2)
	for _, p := range all {
		require.Equal(t, domain.StatusNotStarted, p.Status)
		require.False(t, p.Orphaned)
	}
}

func TestReconcile_PreservaProgressoExistente(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	p, err := s.Get("go/sintaxe/a")
	require.NoError(t, err)
	require.NoError(t, s.Save(domain.Apply(p, true, "hash1", time.Now())))

	// Um segundo boot não pode zerar o que já foi conquistado.
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	got, err := s.Get("go/sintaxe/a")
	require.NoError(t, err)
	require.Equal(t, domain.StatusCompleted, got.Status)
	require.Equal(t, 1, got.Attempts)
}

func TestReconcile_MarcaOrfaosSemDeletar(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/antigo")))

	p, _ := s.Get("go/sintaxe/antigo")
	require.NoError(t, s.Save(domain.Apply(p, true, "h", time.Now())))

	// O exercício foi renomeado: some do disco e aparece outro no lugar.
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/novo")))

	antigo, err := s.Get("go/sintaxe/antigo")
	require.NoError(t, err)
	require.True(t, antigo.Orphaned, "sumiu do disco, então é órfão")
	require.Equal(t, domain.StatusCompleted, antigo.Status, "mas o progresso continua lá")

	novo, err := s.Get("go/sintaxe/novo")
	require.NoError(t, err)
	require.Equal(t, domain.StatusNotStarted, novo.Status)
}

func TestReconcile_OrfaoQueVoltaDeixaDeSerOrfao(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))
	require.NoError(t, s.Reconcile(exercicios()))

	orfao, _ := s.Get("go/sintaxe/a")
	require.True(t, orfao.Orphaned)

	// Renomear e desfazer é um caso real; o órfão precisa voltar ao normal.
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	voltou, _ := s.Get("go/sintaxe/a")
	require.False(t, voltou.Orphaned)
}

func TestSave_RoundTripPreservaTudo(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	quando := time.Date(2026, 8, 13, 15, 30, 45, 123456789, time.UTC)
	original := domain.Progress{
		Path:             "go/sintaxe/a",
		Language:         "go",
		Category:         domain.CategorySintaxe,
		Status:           domain.StatusCompleted,
		Attempts:         7,
		FirstCompletedAt: &quando,
		LastValidatedAt:  &quando,
		LastContentHash:  "abc123",
		LastPassed:       true,
	}
	require.NoError(t, s.Save(original))

	got, err := s.Get("go/sintaxe/a")
	require.NoError(t, err)

	require.Equal(t, original.Status, got.Status)
	require.Equal(t, original.Attempts, got.Attempts)
	require.Equal(t, original.LastContentHash, got.LastContentHash)
	require.Equal(t, original.LastPassed, got.LastPassed)
	require.NotNil(t, got.FirstCompletedAt)
	require.True(t, quando.Equal(*got.FirstCompletedAt), "timestamp precisa sobreviver ao round-trip")
	require.True(t, quando.Equal(*got.LastValidatedAt))
}

func TestSave_TimestampsNulos(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	require.NoError(t, s.Save(domain.Progress{
		Path: "go/sintaxe/a", Language: "go", Category: domain.CategorySintaxe,
		Status: domain.StatusInProgress,
	}))

	got, err := s.Get("go/sintaxe/a")
	require.NoError(t, err)
	require.Nil(t, got.FirstCompletedAt)
	require.Nil(t, got.LastValidatedAt)
}

func TestSave_PersisteRegressao(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	agora := time.Now()
	p, _ := s.Get("go/sintaxe/a")
	p = domain.Apply(p, true, "h1", agora)                   // completa
	p = domain.Apply(p, false, "h2", agora.Add(time.Minute)) // depois quebra
	require.NoError(t, s.Save(p))

	got, err := s.Get("go/sintaxe/a")
	require.NoError(t, err)
	require.Equal(t, domain.StatusCompleted, got.Status)
	require.True(t, got.Regressed(), "a regressão precisa sobreviver ao reinício do CLI")
}

func TestByLanguage(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile([]domain.Exercise{
		{Path: "go/sintaxe/a", Language: "go", Category: domain.CategorySintaxe},
		{Path: "rust/sintaxe/b", Language: "rust", Category: domain.CategorySintaxe},
	}))

	got, err := s.ByLanguage("go")
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, "go/sintaxe/a", got[0].Path)
}

func TestReset(t *testing.T) {
	preparar := func(t *testing.T) *Store {
		s := abrir(t)
		require.NoError(t, s.Reconcile([]domain.Exercise{
			{Path: "go/sintaxe/a", Language: "go", Category: domain.CategorySintaxe},
			{Path: "go/sintaxe/b", Language: "go", Category: domain.CategorySintaxe},
			{Path: "rust/sintaxe/c", Language: "rust", Category: domain.CategorySintaxe},
		}))
		for _, path := range []string{"go/sintaxe/a", "go/sintaxe/b", "rust/sintaxe/c"} {
			p, _ := s.Get(path)
			require.NoError(t, s.Save(domain.Apply(p, true, "h", time.Now())))
		}
		return s
	}

	t.Run("por exercício", func(t *testing.T) {
		s := preparar(t)

		n, err := s.Reset(ResetScope{Path: "go/sintaxe/a"})
		require.NoError(t, err)
		require.Equal(t, 1, n)

		a, _ := s.Get("go/sintaxe/a")
		require.Equal(t, domain.StatusNotStarted, a.Status)
		require.Zero(t, a.Attempts)
		require.Nil(t, a.FirstCompletedAt)
		require.Empty(t, a.LastContentHash, "sem zerar o hash, o próximo save seria descartado como 'nada mudou'")

		b, _ := s.Get("go/sintaxe/b")
		require.Equal(t, domain.StatusCompleted, b.Status, "reset de um não pode afetar o vizinho")
	})

	t.Run("por linguagem", func(t *testing.T) {
		s := preparar(t)

		n, err := s.Reset(ResetScope{Language: "go"})
		require.NoError(t, err)
		require.Equal(t, 2, n)

		c, _ := s.Get("rust/sintaxe/c")
		require.Equal(t, domain.StatusCompleted, c.Status)
	})

	t.Run("global", func(t *testing.T) {
		s := preparar(t)

		n, err := s.Reset(ResetScope{})
		require.NoError(t, err)
		require.Equal(t, 3, n)

		all, _ := s.All()
		for _, p := range all {
			require.Equal(t, domain.StatusNotStarted, p.Status)
		}
	})
}

func TestForgetOrphans(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a", "go/sintaxe/b")))
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/a")))

	n, err := s.ForgetOrphans()
	require.NoError(t, err)
	require.Equal(t, 1, n)

	all, _ := s.All()
	require.Len(t, all, 1)
}

func TestRemap(t *testing.T) {
	s := abrir(t)
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/antigo")))
	p, _ := s.Get("go/sintaxe/antigo")
	require.NoError(t, s.Save(domain.Apply(p, true, "h", time.Now())))

	// O exercício foi renomeado; o usuário aponta o progresso para o novo path.
	require.NoError(t, s.Reconcile(exercicios("go/sintaxe/novo")))
	require.NoError(t, s.Remap("go/sintaxe/antigo", "go/sintaxe/novo"))

	novo, err := s.Get("go/sintaxe/novo")
	require.NoError(t, err)
	require.Equal(t, domain.StatusCompleted, novo.Status)
	require.False(t, novo.Orphaned)

	all, _ := s.All()
	require.Len(t, all, 1, "o registro antigo some depois do remap")
}

func TestRemap_OrigemInexistente(t *testing.T) {
	s := abrir(t)
	err := s.Remap("go/sintaxe/fantasma", "go/sintaxe/novo")
	require.Error(t, err)
	require.Contains(t, err.Error(), "fantasma")
}

func TestVersionOf(t *testing.T) {
	v, err := versionOf("0001_init.sql")
	require.NoError(t, err)
	require.Equal(t, 1, v)

	_, err = versionOf("init.sql")
	require.Error(t, err)
}
