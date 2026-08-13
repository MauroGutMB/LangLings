package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
)

// ContentHash resume o conteúdo dos arquivos editáveis de um exercício.
//
// É o árbitro entre as duas fontes do watcher: nem o fsnotify nem o polling
// decidem se algo mudou — ambos apenas sugerem "pode ter mudado", e o engine
// compara este hash com o da última validação. Se for igual, descarta. Isso
// torna eventos duplicados e rajadas de save irrelevantes.
//
// Os nomes entram ordenados para que a iteração aleatória de map não produza
// hashes diferentes para o mesmo conteúdo, e o tamanho de cada parte entra no
// resumo para que dois arquivos não possam ser confundidos por concatenação
// (ex: {"a": "xy", "b": ""} vs {"a": "x", "b": "y"}).
func ContentHash(files map[string][]byte) string {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)

	h := sha256.New()
	for _, name := range names {
		content := files[name]
		fmt.Fprintf(h, "%d:%s\n%d:", len(name), name, len(content))
		h.Write(content)
		h.Write([]byte("\n"))
	}
	return hex.EncodeToString(h.Sum(nil))
}
