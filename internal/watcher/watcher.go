// Package watcher detecta que o usuário salvou um arquivo de exercício.
//
// É o componente mais frágil do projeto e o que falha do jeito mais cruel:
// quando ele para de funcionar, nada explode — o CLI apenas deixa de reagir, e
// você fica achando que o problema é o seu código de exercício. Por isso o
// desenho aqui é conservador de propósito.
//
// A regra central: nem o fsnotify nem o polling decidem coisa alguma. Ambos
// apenas dizem "pode ter mudado". Quem decide é um único ponto que compara um
// retrato dos arquivos com o retrato anterior. Isso torna eventos duplicados,
// rajadas e a ordem entre as duas fontes irrelevantes.
package watcher

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Origens possíveis de um sinal, úteis no log para diagnosticar "parou de
// reagir" em segundos em vez de horas.
const (
	SourceFsnotify = "fsnotify"
	SourcePolling  = "polling"
	SourceManual   = "manual"
)

// Event é a notificação de que os arquivos observados mudaram de verdade.
type Event struct {
	At     time.Time
	Source string
}

// Config ajusta os tempos. Os padrões valem para uso real; os testes injetam
// durações curtas para rodar em milissegundos.
type Config struct {
	// Debounce é o silêncio necessário antes de emitir. Um único save costuma
	// gerar de dois a quatro eventos de inotify em rajada.
	Debounce time.Duration

	// Poll é o intervalo da rede de segurança. Ele existe porque o inotify
	// morre em dois casos comuns: salvamento atômico (o editor escreve um
	// temporário e faz rename, deixando o watch preso ao inode antigo) e
	// sistemas de arquivo sem suporte, como o /mnt/c do WSL.
	Poll time.Duration
}

const (
	DefaultDebounce = 300 * time.Millisecond
	DefaultPoll     = 250 * time.Millisecond
)

func (c Config) withDefaults() Config {
	if c.Debounce <= 0 {
		c.Debounce = DefaultDebounce
	}
	if c.Poll <= 0 {
		c.Poll = DefaultPoll
	}
	return c
}

// fileState é o retrato de um arquivo observado.
type fileState struct {
	exists bool
	size   int64
	mod    time.Time
}

// Watch observa os arquivos de allowlist dentro de root e emite um Event a
// cada mudança real de conteúdo.
//
// allowlist contém caminhos relativos a root — é a lista `editable` do
// manifesto. Tudo o mais dentro de root é artefato de build e é ignorado: sem
// isso, o binário gerado por um `go build` dispararia revalidação em cascata.
//
// O canal é fechado quando ctx termina.
func Watch(ctx context.Context, root string, allowlist []string, cfg Config) (<-chan Event, error) {
	cfg = cfg.withDefaults()

	if len(allowlist) == 0 {
		return nil, fmt.Errorf("allowlist vazia: o watcher não saberia o que observar")
	}

	watched := make([]string, 0, len(allowlist))
	dirs := map[string]bool{}
	for _, rel := range allowlist {
		abs := filepath.Join(root, filepath.FromSlash(rel))
		watched = append(watched, abs)
		dirs[filepath.Dir(abs)] = true
	}
	sort.Strings(watched)

	// Sinais brutos das duas fontes. Buffer generoso: uma rajada de inotify
	// nunca deve bloquear a goroutine que a lê.
	raw := make(chan string, 64)
	events := make(chan Event)

	// O retrato inicial é tirado aqui, de forma síncrona, ANTES de qualquer
	// goroutine subir. Tirá-lo lá dentro abriria uma janela em que um save
	// chegado logo depois de Watch() retornar já apareceria no primeiro
	// retrato — e, como nada mais mudaria, o watcher ficaria mudo para sempre.
	// Salvar imediatamente após abrir o exercício é comportamento comum.
	inicial := snapshot(watched)

	fsw, err := startFsnotify(ctx, dirs, watched, raw)
	if err != nil {
		// Sem inotify o watcher continua funcionando só com polling. Degradar
		// é melhor que falhar: é exatamente o cenário de um workspace em
		// /mnt/c, onde o inotify não existe mas o arquivo muda do mesmo jeito.
		fsw = nil
	}

	go pollLoop(ctx, cfg.Poll, raw)
	go decide(ctx, watched, cfg.Debounce, raw, events, fsw, inicial)

	return events, nil
}

// startFsnotify observa os diretórios (não os arquivos).
//
// Observar o diretório é o que sobrevive ao salvamento atômico: quando o
// editor troca o arquivo por um novo inode, um watch no arquivo antigo fica
// órfão em silêncio, mas o watch no diretório continua vendo a criação.
func startFsnotify(ctx context.Context, dirs map[string]bool, watched []string, raw chan<- string) (*fsnotify.Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	interesse := make(map[string]bool, len(watched))
	for _, p := range watched {
		interesse[p] = true
	}

	added := 0
	for dir := range dirs {
		if err := w.Add(dir); err == nil {
			added++
		}
	}
	if added == 0 {
		w.Close()
		return nil, fmt.Errorf("nenhum diretório pôde ser observado")
	}

	go func() {
		defer w.Close()
		for {
			select {
			case <-ctx.Done():
				return
			case ev, ok := <-w.Events:
				if !ok {
					return
				}
				// Filtrar aqui evita que um `go build` gerando artefatos
				// acorde a comparação de retratos dezenas de vezes.
				if interesse[filepath.Clean(ev.Name)] {
					send(raw, SourceFsnotify)
				}
			case _, ok := <-w.Errors:
				if !ok {
					return
				}
				// Erro de inotify não derruba o watcher: o polling cobre.
			}
		}
	}()

	return w, nil
}

// pollLoop é a rede de segurança. Ele não compara nada — apenas pede, em
// intervalos fixos, que os retratos sejam conferidos.
func pollLoop(ctx context.Context, every time.Duration, raw chan<- string) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			send(raw, SourcePolling)
		}
	}
}

// decide é o único ponto com poder de decisão.
//
// Ele mantém o retrato dos arquivos observados. Um sinal bruto, venha de onde
// vier, só vira Event se o retrato realmente mudou — e mesmo assim depois de
// um período de silêncio, para que uma rajada de saves vire uma emissão só.
func decide(ctx context.Context, watched []string, debounce time.Duration, raw <-chan string, events chan<- Event, fsw *fsnotify.Watcher, inicial map[string]fileState) {
	defer close(events)

	last := inicial

	// Timer parado: só começa a contar quando há mudança pendente.
	timer := time.NewTimer(debounce)
	if !timer.Stop() {
		<-timer.C
	}

	var (
		pending bool
		origem  string
	)

	for {
		select {
		case <-ctx.Done():
			if fsw != nil {
				fsw.Close()
			}
			return

		case source := <-raw:
			current := snapshot(watched)
			if sameSnapshot(last, current) {
				continue
			}
			last = current

			// Cada mudança nova reinicia a contagem: o que interessa é o
			// silêncio depois do último save, não o tempo desde o primeiro.
			if pending && !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(debounce)
			pending = true
			origem = source

		case <-timer.C:
			if !pending {
				continue
			}
			pending = false

			select {
			case events <- Event{At: time.Now(), Source: origem}:
			case <-ctx.Done():
				if fsw != nil {
					fsw.Close()
				}
				return
			}
		}
	}
}

func snapshot(paths []string) map[string]fileState {
	out := make(map[string]fileState, len(paths))
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			out[p] = fileState{exists: false}
			continue
		}
		out[p] = fileState{exists: true, size: info.Size(), mod: info.ModTime()}
	}
	return out
}

func sameSnapshot(a, b map[string]fileState) bool {
	if len(a) != len(b) {
		return false
	}
	for path, sa := range a {
		sb, ok := b[path]
		if !ok || sa.exists != sb.exists || sa.size != sb.size || !sa.mod.Equal(sb.mod) {
			return false
		}
	}
	return true
}

// send nunca bloqueia: um sinal perdido numa rajada é irrelevante, porque o
// próximo tick do polling reconfere os mesmos retratos.
func send(ch chan<- string, source string) {
	select {
	case ch <- source:
	default:
	}
}
