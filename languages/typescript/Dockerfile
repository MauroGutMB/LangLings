FROM node:24-alpine

# O Node executa .ts apagando os tipos, sem CHECAR nada: um erro de tipo em
# código que os testes não executam passaria com exit 0. Sem tsc na imagem, um
# exercício de TypeScript não estaria validando TypeScript — e a validação roda
# com --network none, então instalar na hora não é opção.
#
# 5.9.3 e não a 7.x: a 7 é a reescrita nativa do compilador e merece ser
# avaliada como mudança própria, não carregada de brinde numa fase de conteúdo.
RUN npm install -g typescript@5.9.3 \
 && npm install --prefix /opt/ts @types/node@24 \
 && npm cache clean --force

CMD ["/bin/sh"]
