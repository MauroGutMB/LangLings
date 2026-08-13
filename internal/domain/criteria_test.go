package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCriterion_Probe(t *testing.T) {
	t.Run("file_exists vira stat", func(t *testing.T) {
		c := Criterion{Kind: KindFileExists, Path: "bin/hello"}
		require.Equal(t, Probe{Kind: ProbeStat, Path: "bin/hello"}, c.Probe())
	})

	t.Run("demais kinds viram run", func(t *testing.T) {
		c := Criterion{Kind: KindStdoutEquals, Command: []string{"./bin/hello"}}
		require.Equal(t, Probe{Kind: ProbeRun, Command: []string{"./bin/hello"}}, c.Probe())
	})
}

func TestCriterion_Evaluate(t *testing.T) {
	tests := []struct {
		nome      string
		criterio  Criterion
		resultado ProbeResult
		querPasso bool
	}{
		{
			nome:      "file_exists aprova quando o artefato existe",
			criterio:  Criterion{Kind: KindFileExists, Path: "bin/hello"},
			resultado: ProbeResult{Exists: true},
			querPasso: true,
		},
		{
			nome:      "file_exists reprova quando não existe",
			criterio:  Criterion{Kind: KindFileExists, Path: "bin/hello"},
			resultado: ProbeResult{Exists: false},
			querPasso: false,
		},
		{
			nome:      "exit_code aprova no código esperado",
			criterio:  Criterion{Kind: KindExitCode, Command: []string{"true"}, Code: 0},
			resultado: ProbeResult{ExitCode: 0},
			querPasso: true,
		},
		{
			nome:      "exit_code reprova em código diferente",
			criterio:  Criterion{Kind: KindExitCode, Command: []string{"false"}, Code: 0},
			resultado: ProbeResult{ExitCode: 1},
			querPasso: false,
		},
		{
			nome:      "exit_code pode esperar falha",
			criterio:  Criterion{Kind: KindExitCode, Command: []string{"false"}, Code: 1},
			resultado: ProbeResult{ExitCode: 1},
			querPasso: true,
		},
		{
			nome:      "stdout_equals aprova na saída exata",
			criterio:  Criterion{Kind: KindStdoutEquals, Command: []string{"echo"}, Expected: "olá\n"},
			resultado: ProbeResult{Stdout: "olá\n"},
			querPasso: true,
		},
		{
			nome:      "stdout_equals é sensível a espaço em branco",
			criterio:  Criterion{Kind: KindStdoutEquals, Command: []string{"echo"}, Expected: "olá"},
			resultado: ProbeResult{Stdout: "olá\n"},
			querPasso: false,
		},
		{
			nome:      "stdout_matches aprova quando casa",
			criterio:  Criterion{Kind: KindStdoutMatches, Command: []string{"file"}, Pattern: `PE32\+.*x86-64`},
			resultado: ProbeResult{Stdout: "bin/hello.exe: PE32+ executable (console) x86-64, for MS Windows"},
			querPasso: true,
		},
		{
			nome:      "stdout_matches reprova quando não casa",
			criterio:  Criterion{Kind: KindStdoutMatches, Command: []string{"file"}, Pattern: `PE32\+`},
			resultado: ProbeResult{Stdout: "bin/hello: ELF 64-bit LSB executable"},
			querPasso: false,
		},
		{
			nome:      "kind desconhecido reprova em vez de entrar em pânico",
			criterio:  Criterion{Kind: "inventado"},
			resultado: ProbeResult{},
			querPasso: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			got := tt.criterio.Evaluate(tt.resultado)
			require.Equal(t, tt.querPasso, got.Passed)
			require.NotEmpty(t, got.Detail, "todo veredito precisa de motivo exibível")
		})
	}
}

func TestCriterion_Evaluate_FalhaDeInfraNaoEhReprovacaoDoExercicio(t *testing.T) {
	// Container morto ou timeout é diferente de "o critério não foi satisfeito".
	// Em ambos o resultado é falha, mas o motivo precisa dizer a verdade.
	c := Criterion{Kind: KindFileExists, Path: "bin/hello"}

	got := c.Evaluate(ProbeResult{Exists: true, Err: errors.New("container excedeu o limite de memória")})

	require.False(t, got.Passed)
	require.Contains(t, got.Detail, "não foi possível verificar")
	require.Contains(t, got.Detail, "memória")
}

func TestCriterion_Validate(t *testing.T) {
	tests := []struct {
		nome     string
		criterio Criterion
		querErro string // substring esperada; vazio = deve passar
	}{
		{
			nome:     "file_exists válido",
			criterio: Criterion{Kind: KindFileExists, Path: "bin/hello"},
		},
		{
			nome:     "file_exists sem path",
			criterio: Criterion{Kind: KindFileExists},
			querErro: "exige path",
		},
		{
			nome:     "file_exists não pode escapar do exercício",
			criterio: Criterion{Kind: KindFileExists, Path: "../../etc/passwd"},
			querErro: "escapar",
		},
		{
			nome:     "exit_code sem command",
			criterio: Criterion{Kind: KindExitCode},
			querErro: "exige command",
		},
		{
			nome:     "stdout_equals sem expected",
			criterio: Criterion{Kind: KindStdoutEquals, Command: []string{"echo"}},
			querErro: "exige expected",
		},
		{
			nome:     "stdout_matches sem pattern",
			criterio: Criterion{Kind: KindStdoutMatches, Command: []string{"file"}},
			querErro: "exige pattern",
		},
		{
			nome:     "stdout_matches com regex quebrada falha cedo",
			criterio: Criterion{Kind: KindStdoutMatches, Command: []string{"file"}, Pattern: "PE32(+"},
			querErro: "pattern inválido",
		},
		{
			nome:     "kind desconhecido",
			criterio: Criterion{Kind: "arquivo_existe"},
			querErro: "desconhecido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			err := tt.criterio.Validate()
			if tt.querErro == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.querErro)
		})
	}
}

func TestCriterion_Label(t *testing.T) {
	t.Run("usa describe quando presente", func(t *testing.T) {
		c := Criterion{Kind: KindFileExists, Path: "bin/hello", Describe: "o binário foi gerado"}
		require.Equal(t, "o binário foi gerado", c.Label())
	})

	t.Run("gera um texto útil sem describe", func(t *testing.T) {
		c := Criterion{Kind: KindFileExists, Path: "bin/hello"}
		require.Equal(t, "bin/hello existe", c.Label())
	})
}
