package domain

import "time"

// Status é o estado de um exercício, exibido como 🔴 / 🟡 / 🟢.
type Status string

const (
	StatusNotStarted Status = "not_started"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

// Icon é o marcador exibido na lista de exercícios.
func (s Status) Icon() string {
	switch s {
	case StatusCompleted:
		return "🟢"
	case StatusInProgress:
		return "🟡"
	default:
		return "🔴"
	}
}

// Progress é o estado persistido de um exercício.
type Progress struct {
	Path     string
	Language string
	Category Category

	Status   Status
	Attempts int

	FirstCompletedAt *time.Time
	LastValidatedAt  *time.Time

	// LastContentHash é o hash dos arquivos editáveis na última validação.
	// É o que permite descartar um save que não mudou nada.
	LastContentHash string

	// LastPassed é o resultado da última validação. Combinado com Status,
	// dispensa uma coluna separada para "regrediu".
	LastPassed bool

	// Orphaned marca um exercício que existe no banco mas sumiu do disco.
	// Registros órfãos nunca são deletados em silêncio.
	Orphaned bool
}

// Regressed informa que o exercício já foi concluído mas a última validação
// falhou — o usuário quebrou o código depois de resolver.
func (p Progress) Regressed() bool {
	return p.Status == StatusCompleted && p.LastValidatedAt != nil && !p.LastPassed
}

// Apply devolve o progresso resultante de uma validação executada.
//
// A regra que merece atenção: um exercício `completed` que passa a falhar
// permanece `completed`. Progresso é histórico, não estado atual — se você
// concluiu e depois quebrou o código experimentando, a lista continua verde e
// a tela do exercício mostra o aviso de regressão. A alternativa (voltar para
// in_progress) faria você perder a memória de ter completado.
//
// É função pura: nenhuma leitura de relógio, nenhum I/O. `now` entra por
// parâmetro justamente para os testes serem determinísticos.
func Apply(cur Progress, passed bool, hash string, now time.Time) Progress {
	next := cur
	next.Attempts = cur.Attempts + 1
	next.LastValidatedAt = &now
	next.LastContentHash = hash
	next.LastPassed = passed

	switch {
	case passed && cur.Status != StatusCompleted:
		next.Status = StatusCompleted
		completedAt := now
		next.FirstCompletedAt = &completedAt

	case passed:
		// Já estava completo e continua passando: nada a mudar além dos
		// contadores. FirstCompletedAt é preservado.

	case cur.Status == StatusCompleted:
		// Regressão: mantém o histórico. Regressed() passa a devolver true.

	default:
		// Primeira tentativa que falhou, ou mais uma: em progresso.
		next.Status = StatusInProgress
	}

	return next
}

// Summary agrega o progresso de um conjunto de exercícios, para o resumo da
// tela de boas-vindas e os contadores por linguagem.
type Summary struct {
	Total      int
	Completed  int
	InProgress int
	NotStarted int
}

// Summarize conta os estados de uma lista de progressos, ignorando órfãos.
func Summarize(items []Progress) Summary {
	var s Summary
	for _, p := range items {
		if p.Orphaned {
			continue
		}
		s.Total++
		switch p.Status {
		case StatusCompleted:
			s.Completed++
		case StatusInProgress:
			s.InProgress++
		default:
			s.NotStarted++
		}
	}
	return s
}
