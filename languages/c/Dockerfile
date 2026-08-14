FROM alpine:3.22

# build-base = gcc + make + musl-dev + binutils, que é tudo o que os exercícios
# de basico/sintaxe precisam. A imagem oficial gcc:16 entrega o mesmo compilador
# em 561 MB, mais um toolchain de bootstrap que ninguém usa aqui.
RUN apk add --no-cache build-base

CMD ["/bin/sh"]
