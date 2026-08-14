-- A mesma tabela também serve de dicionário: chaves quaisquer, não só
-- índices numéricos seguidos.
local M = {}

function M.exemplos()
  local pessoa = {nome = "Ana", idade = 30}

  -- t.chave é açúcar de sintaxe para t["chave"] — as duas formas acessam o
  -- mesmo par.
  print(pessoa.nome, pessoa["idade"]) -- Ana 30

  -- Ler uma chave que não existe devolve nil, sem erro.
  print(pessoa.profissao) -- nil

  -- Atribuir cria a chave se ela não existia, ou substitui o valor se já
  -- existia.
  pessoa.profissao = "engenheira"
  pessoa.idade = 31
  print(pessoa.profissao, pessoa.idade) -- engenheira 31

  -- Atribuir nil a uma chave APAGA o par — não sobra rastro dela na tabela.
  pessoa.profissao = nil
  print(pessoa.profissao) -- nil

  -- pairs percorre toda chave que existe, numérica ou string, em ordem NÃO
  -- garantida. ipairs, em contraste, só anda pelas posições 1, 2, 3...
  local contagem = {maca = 3, pera = 5, uva = 1}
  local total = 0
  for chave, valor in pairs(contagem) do
    total = total + valor
    -- chave e valor aqui são, por exemplo, "maca" e 3 — a ordem entre as
    -- três chaves não é previsível.
  end
  print(total) -- 9

  -- Uma tabela pode misturar as duas formas: posições numéricas e chaves
  -- string convivem na mesma tabela.
  local mista = {"primeiro", "segundo", extra = "terceiro"}
  print(mista[1], mista[2], mista.extra) -- primeiro segundo terceiro
end

-- SUA VEZ
--
-- Devolva quantos pares chave/valor existem em t. # não serve: ele só conta
-- posições numéricas seguidas a partir de 1, e t pode não ter nenhuma.
function M.contarChaves(t)
  return -1 -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
