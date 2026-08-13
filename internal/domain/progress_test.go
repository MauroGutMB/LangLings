package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	t0 = time.Date(2026, 8, 13, 10, 0, 0, 0, time.UTC)
	t1 = t0.Add(time.Hour)
)

func TestApply_Transicoes(t *testing.T) {
	tests := []struct {
		nome       string
		atual      Progress
		passou     bool
		querStatus Status
		// querPrimeiraConclusao: nil = deve continuar sem data de conclusão
		querPrimeiraConclusao *time.Time
		querRegrediu          bool
	}{
		{
			nome:       "não iniciado + falha vira em progresso",
			atual:      Progress{Status: StatusNotStarted},
			passou:     false,
			querStatus: StatusInProgress,
		},
		{
			nome:                  "não iniciado + sucesso vai direto para completo",
			atual:                 Progress{Status: StatusNotStarted},
			passou:                true,
			querStatus:            StatusCompleted,
			querPrimeiraConclusao: &t1,
		},
		{
			nome:       "em progresso + falha continua em progresso",
			atual:      Progress{Status: StatusInProgress, Attempts: 3},
			passou:     false,
			querStatus: StatusInProgress,
		},
		{
			nome:                  "em progresso + sucesso vira completo",
			atual:                 Progress{Status: StatusInProgress, Attempts: 3},
			passou:                true,
			querStatus:            StatusCompleted,
			querPrimeiraConclusao: &t1,
		},
		{
			nome: "completo + falha PERMANECE completo, marcando regressão",
			atual: Progress{
				Status:           StatusCompleted,
				FirstCompletedAt: &t0,
				LastPassed:       true,
			},
			passou:                false,
			querStatus:            StatusCompleted,
			querPrimeiraConclusao: &t0, // preservado, não sobrescrito
			querRegrediu:          true,
		},
		{
			nome: "completo + sucesso continua completo sem regressão",
			atual: Progress{
				Status:           StatusCompleted,
				FirstCompletedAt: &t0,
				LastPassed:       true,
			},
			passou:                true,
			querStatus:            StatusCompleted,
			querPrimeiraConclusao: &t0,
		},
		{
			nome: "regressão é reversível: volta a passar e deixa de regredir",
			atual: Progress{
				Status:           StatusCompleted,
				FirstCompletedAt: &t0,
				LastValidatedAt:  &t0,
				LastPassed:       false,
			},
			passou:                true,
			querStatus:            StatusCompleted,
			querPrimeiraConclusao: &t0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			got := Apply(tt.atual, tt.passou, "hash-abc", t1)

			require.Equal(t, tt.querStatus, got.Status)
			require.Equal(t, tt.atual.Attempts+1, got.Attempts, "toda validação conta como tentativa")
			require.Equal(t, "hash-abc", got.LastContentHash)
			require.Equal(t, &t1, got.LastValidatedAt)
			require.Equal(t, tt.passou, got.LastPassed)
			require.Equal(t, tt.querRegrediu, got.Regressed())

			if tt.querPrimeiraConclusao == nil {
				require.Nil(t, got.FirstCompletedAt)
			} else {
				require.NotNil(t, got.FirstCompletedAt)
				require.Equal(t, *tt.querPrimeiraConclusao, *got.FirstCompletedAt)
			}
		})
	}
}

func TestApply_NaoMutaOriginal(t *testing.T) {
	original := Progress{Status: StatusNotStarted, Attempts: 0}
	Apply(original, true, "h", t1)

	require.Equal(t, StatusNotStarted, original.Status, "Apply deve devolver cópia, não mutar")
	require.Equal(t, 0, original.Attempts)
}

func TestRegressed_ExigeValidacaoPrevia(t *testing.T) {
	// Um exercício marcado como completo mas nunca validado nesta instalação
	// não é uma regressão — é ausência de dado.
	p := Progress{Status: StatusCompleted, LastPassed: false}
	require.False(t, p.Regressed())
}

func TestSummarize(t *testing.T) {
	items := []Progress{
		{Status: StatusCompleted},
		{Status: StatusCompleted},
		{Status: StatusInProgress},
		{Status: StatusNotStarted},
		{Status: StatusCompleted, Orphaned: true}, // órfão não conta
	}

	got := Summarize(items)

	require.Equal(t, Summary{Total: 4, Completed: 2, InProgress: 1, NotStarted: 1}, got)
}

func TestStatusIcon(t *testing.T) {
	require.Equal(t, "🔴", StatusNotStarted.Icon())
	require.Equal(t, "🟡", StatusInProgress.Icon())
	require.Equal(t, "🟢", StatusCompleted.Icon())
}
