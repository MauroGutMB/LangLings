// Este go.mod não existe para ser compilado: ele existe para NÃO ser.
//
// Sem ele, os diretórios solution/ dos exercícios entram no módulo principal e
// passam a ser compilados por `go build ./...` e `go vet ./...` do LangLings.
// Uma solução de exercício que não compile isoladamente — ou dois exercícios
// com o mesmo nome de pacote no mesmo caminho — quebrariam a build do projeto
// inteiro por um motivo que não tem nada a ver com o projeto.
//
// Declarar exercises/ como um módulo à parte faz o toolchain ignorar tudo o
// que está aqui dentro. Cada exercício continua trazendo o próprio go.mod em
// base/, que é o que vale quando ele é copiado para o workspace.
module langlings-content

go 1.24
