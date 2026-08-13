# Comandos que o usuário digita no shell do container ([s] na tela).
#
# Este arquivo não é código do exercício: é o roteiro que torna a solução
# verificável por máquina. Sem ele, `langlings verify` não teria como provar
# que os critérios são satisfazíveis.
set -e

mkdir -p bin
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.versao=1.0.0" -o bin/hello.exe .
