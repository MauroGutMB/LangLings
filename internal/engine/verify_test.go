package engine

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"langlings/internal/runner"
)

// julgarPelaSolucao faz o runner falso se comportar como um test runner de
// verdade: aprova quando o código no workspace é o da solução, reprova quando
// é o código-base. É o mínimo para exercitar o gate.
func (a *ambiente) julgarPelaSolucao() {
	a.fake.EphemeralFunc = func(spec runner.SessionSpec, cmd []string, _ runner.ExecOpts) runner.ExecResult {
		conteudo, err := os.ReadFile(filepath.Join(spec.Workspace, "go", "sintaxe", "slices", "main.go"))
		if err != nil {
			return runner.ExecResult{ExitCode: 1, Stderr: err.Error()}
		}
		if strings.Contains(string(conteudo), "solução") {
			return runner.ExecResult{ExitCode: 0, Stdout: "ok\n"}
		}
		return runner.ExecResult{ExitCode: 1, Stdout: "FAIL: Dobrar devolveu o slice original\n"}
	}
}

func TestVerify_ExercicioBemFormado(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.julgarPelaSolucao()

	report := a.engine.Verify(context.Background(), a.exercicio("go/sintaxe/slices"))

	require.True(t, report.OK(), "problemas: %v", report.Problems)
}

func TestVerify_NaoEncostaNoWorkspaceDoUsuario(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.julgarPelaSolucao()
	ex := a.exercicio("go/sintaxe/slices")

	dir, err := a.paths.Materialize(ex)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "main.go"), []byte("meu trabalho em andamento"), 0o644))

	a.engine.Verify(context.Background(), ex)

	conteudo, err := os.ReadFile(filepath.Join(dir, "main.go"))
	require.NoError(t, err)
	require.Equal(t, "meu trabalho em andamento", string(conteudo),
		"verificar conteúdo nunca pode mexer no que o usuário escreveu")
}

func TestVerify_BaseQueJaPassaEhReprovado(t *testing.T) {
	// O defeito mais provável em conteúdo escrito em lote: um exercício cujo
	// código inicial já satisfaz os testes. Não há nada a resolver, e a olho
	// nu ele parece perfeito.
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.fake.EphemeralFunc = func(runner.SessionSpec, []string, runner.ExecOpts) runner.ExecResult {
		return runner.ExecResult{ExitCode: 0}
	}

	report := a.engine.Verify(context.Background(), a.exercicio("go/sintaxe/slices"))

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), "base/ já passa")
}

func TestVerify_SolucaoQueNaoPassaEhReprovada(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.fake.EphemeralFunc = func(runner.SessionSpec, []string, runner.ExecOpts) runner.ExecResult {
		return runner.ExecResult{ExitCode: 1, Stdout: "FAIL\n"}
	}

	report := a.engine.Verify(context.Background(), a.exercicio("go/sintaxe/slices"))

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), "insolúvel")
}

func TestVerify_SolucaoForaDaAllowlistEhReprovada(t *testing.T) {
	// Se a solução precisa alterar um arquivo que o usuário não pode editar, o
	// exercício é impossível — e ninguém descobre até tentar resolvê-lo.
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	a.julgarPelaSolucao()

	ex := a.exercicio("go/sintaxe/slices")
	escrever(t, filepath.Join(ex.SolutionDir(), "helper.go"), "package main // arquivo escondido\n")

	report := a.engine.Verify(context.Background(), ex)

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), "helper.go não está em editable")
}

func TestVerify_SemHints(t *testing.T) {
	manifesto := `
title     = "Sem dica"
objective = "Resolva."
editable  = ["main.go"]

[validation]
mode    = "test"
command = ["go", "test", "./..."]
`
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifesto})
	a.julgarPelaSolucao()

	report := a.engine.Verify(context.Background(), a.exercicio("go/sintaxe/slices"))

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), "nenhum hint")
}

func TestVerify_SemDiretorioSolution(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/sintaxe/slices": manifestoTeste})
	ex := a.exercicio("go/sintaxe/slices")
	require.NoError(t, os.RemoveAll(ex.SolutionDir()))

	report := a.engine.Verify(context.Background(), ex)

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), "solution/ ausente")
}

func TestVerify_ModoCriteriaExigeRoteiroDeComandos(t *testing.T) {
	// Num exercício de Compilador não há "arquivo correto" a comparar: o que
	// se verifica é o efeito de um comando. A solução precisa dizer qual.
	a := novoAmbiente(t, map[string]string{"go/compilador/cross": manifestoCriteria})

	report := a.engine.Verify(context.Background(), a.exercicio("go/compilador/cross"))

	require.False(t, report.OK())
	require.Contains(t, strings.Join(report.Problems, "\n"), SolutionStepsFile)
}

func TestVerify_ModoCriteriaComRoteiro(t *testing.T) {
	a := novoAmbiente(t, map[string]string{"go/compilador/cross": manifestoCriteria})
	ex := a.exercicio("go/compilador/cross")

	escrever(t, filepath.Join(ex.SolutionDir(), SolutionStepsFile), "criar bin/hello.exe\n")

	// O runner falso "compila": cria o artefato no workspace temporário
	// quando o roteiro roda, exatamente como o usuário faria no shell.
	a.fake.EphemeralFunc = func(spec runner.SessionSpec, cmd []string, _ runner.ExecOpts) runner.ExecResult {
		if len(cmd) > 0 && cmd[0] == "sh" {
			artefato := filepath.Join(spec.Workspace, "go", "compilador", "cross", "bin", "hello.exe")
			require.NoError(t, os.MkdirAll(filepath.Dir(artefato), 0o755))
			require.NoError(t, os.WriteFile(artefato, []byte("MZ"), 0o644))
		}
		return runner.ExecResult{ExitCode: 0}
	}

	report := a.engine.Verify(context.Background(), ex)

	require.True(t, report.OK(), "problemas: %v", report.Problems)
}

func TestVerifyAll_PercorreOCatalogo(t *testing.T) {
	a := novoAmbiente(t, map[string]string{
		"go/sintaxe/slices": manifestoTeste,
		"go/sintaxe/maps":   manifestoTeste,
	})
	a.fake.EphemeralFunc = func(runner.SessionSpec, []string, runner.ExecOpts) runner.ExecResult {
		return runner.ExecResult{ExitCode: 0} // base sempre passa: os dois reprovam no gate
	}

	reports := a.engine.VerifyAll(context.Background())

	require.Len(t, reports, 2)
	for _, r := range reports {
		require.False(t, r.OK(), "%s deveria ter sido reprovado", r.Exercise.Path)
	}
}
