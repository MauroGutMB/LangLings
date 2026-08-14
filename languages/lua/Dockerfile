FROM alpine:3.22

# Não existe imagem oficial de Lua no Docker Hub (library/lua = 404). As
# alternativas eram depender de um repositório de terceiro ou duas linhas de
# apk; ficamos com as duas linhas. Resultado: 14,6 MB.
RUN apk add --no-cache lua5.4 \
 && ln -s /usr/bin/lua5.4 /usr/local/bin/lua

CMD ["/bin/sh"]
