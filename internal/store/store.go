// Package store persiste o progresso em SQLite.
//
// Nada de ORM nem de auto-migrate por reflexão: o SQL que roda está em
// migrations/*.sql e é aplicado em ordem, dentro de uma transação. Você lê o
// esquema exato que existe no seu disco.
package store

import (
	"database/sql"
	"embed"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"langlings/internal/domain"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// Store é o acesso ao banco de progresso.
type Store struct {
	db *sql.DB
}

// Open abre (criando se necessário) o banco em path e aplica as migrations.
//
// Os pragmas entram na DSN em vez de num PRAGMA solto depois da conexão porque
// o database/sql mantém um pool: um PRAGMA executado uma vez valeria só para a
// conexão que o executou.
func Open(path string) (*Store, error) {
	dsn := "file:" + url.PathEscape(path) + "?" + strings.Join([]string{
		"_busy_timeout=5000", // duas instâncias do CLI esperam em vez de falhar
		"_journal_mode=WAL",
		"_foreign_keys=on",
	}, "&")

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("abrindo %s: %w", path, err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("conectando em %s: %w", path, err)
	}

	s := &Store{db: db}
	if err := s.migrate(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) Close() error { return s.db.Close() }

// migrate aplica as migrations pendentes, cada uma em sua transação.
func (s *Store) migrate() error {
	if _, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version    INTEGER PRIMARY KEY,
		applied_at TEXT NOT NULL
	)`); err != nil {
		return fmt.Errorf("criando schema_migrations: %w", err)
	}

	applied := map[int]bool{}
	rows, err := s.db.Query(`SELECT version FROM schema_migrations`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var v int
		if err := rows.Scan(&v); err != nil {
			return err
		}
		applied[v] = true
	}
	if err := rows.Err(); err != nil {
		return err
	}

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return err
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}
	sort.Strings(names)

	for _, name := range names {
		version, err := versionOf(name)
		if err != nil {
			return err
		}
		if applied[version] {
			continue
		}

		content, err := migrationsFS.ReadFile("migrations/" + name)
		if err != nil {
			return err
		}

		tx, err := s.db.Begin()
		if err != nil {
			return err
		}
		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %s: %w", name, err)
		}
		if _, err := tx.Exec(
			`INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)`,
			version, formatTime(time.Now().UTC()),
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("registrando migration %s: %w", name, err)
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

// versionOf extrai o número de "0001_init.sql".
func versionOf(filename string) (int, error) {
	prefix, _, found := strings.Cut(filename, "_")
	if !found {
		return 0, fmt.Errorf("migration %q: esperado <versão>_<nome>.sql", filename)
	}
	v, err := strconv.Atoi(prefix)
	if err != nil {
		return 0, fmt.Errorf("migration %q: versão %q não é número", filename, prefix)
	}
	return v, nil
}

// Reconcile alinha o banco com o catálogo em disco.
//
// Exercícios novos entram como não iniciados; exercícios que sumiram do disco
// são marcados como órfãos, nunca deletados — o progresso de um exercício
// renomeado precisa poder ser recuperado. Órfãos que reaparecem voltam ao
// normal, o que cobre o caso de renomear e desfazer.
func (s *Store) Reconcile(exercises []domain.Exercise) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	present := make(map[string]bool, len(exercises))
	for _, ex := range exercises {
		present[ex.Path] = true

		if _, err := tx.Exec(`
			INSERT INTO progress (path, language, category, status)
			VALUES (?, ?, ?, 'not_started')
			ON CONFLICT (path) DO UPDATE SET
				language = excluded.language,
				category = excluded.category,
				orphaned = 0
		`, ex.Path, ex.Language, string(ex.Category)); err != nil {
			return fmt.Errorf("reconciliando %s: %w", ex.Path, err)
		}
	}

	rows, err := tx.Query(`SELECT path FROM progress WHERE orphaned = 0`)
	if err != nil {
		return err
	}
	var missing []string
	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			rows.Close()
			return err
		}
		if !present[p] {
			missing = append(missing, p)
		}
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}

	for _, p := range missing {
		if _, err := tx.Exec(`UPDATE progress SET orphaned = 1 WHERE path = ?`, p); err != nil {
			return err
		}
	}

	return tx.Commit()
}

const selectColumns = `
	path, language, category, status, attempts,
	first_completed_at, last_validated_at, last_content_hash, last_passed, orphaned
`

// Get devolve o progresso de um exercício. Um exercício ainda não reconciliado
// devolve o zero-value com status não iniciado, em vez de erro: a ausência de
// registro e "nunca comecei" são a mesma coisa para quem chama.
func (s *Store) Get(path string) (domain.Progress, error) {
	row := s.db.QueryRow(`SELECT`+selectColumns+`FROM progress WHERE path = ?`, path)

	p, err := scanProgress(row)
	if err == sql.ErrNoRows {
		return domain.Progress{Path: path, Status: domain.StatusNotStarted}, nil
	}
	return p, err
}

// All devolve todo o progresso conhecido, inclusive órfãos.
func (s *Store) All() ([]domain.Progress, error) {
	return s.query(`SELECT` + selectColumns + `FROM progress ORDER BY path`)
}

// ByLanguage devolve o progresso dos exercícios de uma linguagem.
func (s *Store) ByLanguage(language string) ([]domain.Progress, error) {
	return s.query(`SELECT`+selectColumns+`FROM progress WHERE language = ? ORDER BY path`, language)
}

// Save grava o progresso de um exercício.
func (s *Store) Save(p domain.Progress) error {
	_, err := s.db.Exec(`
		INSERT INTO progress (
			path, language, category, status, attempts,
			first_completed_at, last_validated_at, last_content_hash, last_passed, orphaned
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT (path) DO UPDATE SET
			language           = excluded.language,
			category           = excluded.category,
			status             = excluded.status,
			attempts           = excluded.attempts,
			first_completed_at = excluded.first_completed_at,
			last_validated_at  = excluded.last_validated_at,
			last_content_hash  = excluded.last_content_hash,
			last_passed        = excluded.last_passed,
			orphaned           = excluded.orphaned
	`,
		p.Path, p.Language, string(p.Category), string(p.Status), p.Attempts,
		nullableTime(p.FirstCompletedAt), nullableTime(p.LastValidatedAt),
		p.LastContentHash, boolToInt(p.LastPassed), boolToInt(p.Orphaned),
	)
	if err != nil {
		return fmt.Errorf("salvando %s: %w", p.Path, err)
	}
	return nil
}

// ResetScope descreve o alcance de um reset de estado.
type ResetScope struct {
	// Exatamente um dos dois pode estar preenchido. Ambos vazios = tudo.
	Path     string
	Language string
}

// Reset zera o estado no banco dentro do escopo informado.
//
// Só mexe no banco: apagar ou restaurar os arquivos do workspace é decisão do
// engine, porque é a metade destrutiva da operação.
func (s *Store) Reset(scope ResetScope) (int, error) {
	const clause = `
		UPDATE progress SET
			status             = 'not_started',
			attempts           = 0,
			first_completed_at = NULL,
			last_validated_at  = NULL,
			last_content_hash  = '',
			last_passed        = 0
	`

	var (
		res sql.Result
		err error
	)
	switch {
	case scope.Path != "":
		res, err = s.db.Exec(clause+` WHERE path = ?`, scope.Path)
	case scope.Language != "":
		res, err = s.db.Exec(clause+` WHERE language = ?`, scope.Language)
	default:
		res, err = s.db.Exec(clause)
	}
	if err != nil {
		return 0, fmt.Errorf("resetando: %w", err)
	}

	n, err := res.RowsAffected()
	return int(n), err
}

// ForgetOrphans remove definitivamente os registros órfãos. Só é chamado por
// ação explícita do usuário — a reconciliação sozinha nunca deleta.
func (s *Store) ForgetOrphans() (int, error) {
	res, err := s.db.Exec(`DELETE FROM progress WHERE orphaned = 1`)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	return int(n), err
}

// Remap transfere o progresso de um path antigo para um novo, que é a saída
// para um exercício renomeado. O registro antigo deixa de existir.
func (s *Store) Remap(from, to string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	old, err := scanProgress(tx.QueryRow(`SELECT`+selectColumns+`FROM progress WHERE path = ?`, from))
	if err == sql.ErrNoRows {
		return fmt.Errorf("não há progresso registrado em %q", from)
	}
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
		UPDATE progress SET
			status = ?, attempts = ?, first_completed_at = ?, last_validated_at = ?,
			last_content_hash = ?, last_passed = ?, orphaned = 0
		WHERE path = ?
	`,
		string(old.Status), old.Attempts, nullableTime(old.FirstCompletedAt),
		nullableTime(old.LastValidatedAt), old.LastContentHash, boolToInt(old.LastPassed), to,
	); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM progress WHERE path = ?`, from); err != nil {
		return err
	}
	return tx.Commit()
}

// ---------- helpers ----------

func (s *Store) query(sqlText string, args ...any) ([]domain.Progress, error) {
	rows, err := s.db.Query(sqlText, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Progress
	for rows.Next() {
		p, err := scanProgress(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// scanner cobre *sql.Row e *sql.Rows com a mesma função de leitura.
type scanner interface {
	Scan(dest ...any) error
}

func scanProgress(sc scanner) (domain.Progress, error) {
	var (
		p                  domain.Progress
		category, status   string
		firstAt, lastAt    sql.NullString
		lastPassed, orphan int
	)

	err := sc.Scan(
		&p.Path, &p.Language, &category, &status, &p.Attempts,
		&firstAt, &lastAt, &p.LastContentHash, &lastPassed, &orphan,
	)
	if err != nil {
		return domain.Progress{}, err
	}

	p.Category = domain.Category(category)
	p.Status = domain.Status(status)
	p.LastPassed = lastPassed == 1
	p.Orphaned = orphan == 1

	if p.FirstCompletedAt, err = parseNullTime(firstAt); err != nil {
		return domain.Progress{}, fmt.Errorf("%s: first_completed_at: %w", p.Path, err)
	}
	if p.LastValidatedAt, err = parseNullTime(lastAt); err != nil {
		return domain.Progress{}, fmt.Errorf("%s: last_validated_at: %w", p.Path, err)
	}
	return p, nil
}

const timeLayout = time.RFC3339Nano

func formatTime(t time.Time) string { return t.UTC().Format(timeLayout) }

func nullableTime(t *time.Time) any {
	if t == nil {
		return nil
	}
	return formatTime(*t)
}

func parseNullTime(s sql.NullString) (*time.Time, error) {
	if !s.Valid || s.String == "" {
		return nil, nil
	}
	parsed, err := time.Parse(timeLayout, s.String)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
