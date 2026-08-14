-- Nada aqui é novo: é lista, dicionário, closure, string e erro dos
-- exercícios anteriores, trabalhando juntos num só pedaço de código.
local M = {}

function M.exemplos()
  local vendas = {
    {produto = "cafe", quantidade = 2},
    {produto = "cha", quantidade = 1},
    {produto = "cafe", quantidade = 3},
  }

  -- ipairs percorre a lista; cada elemento é uma tabela com seus próprios
  -- campos, acessados por chave.
  local total = {}
  for _, venda in ipairs(vendas) do
    -- "or 0" cobre o primeiro produto visto, quando total[produto] ainda é
    -- nil — o mesmo idioma de valor padrão que apareceu em funções.
    total[venda.produto] = (total[venda.produto] or 0) + venda.quantidade
  end

  -- pairs percorre o dicionário resultante, em ordem não garantida.
  local linhas = {}
  for produto, quantidade in pairs(total) do
    table.insert(linhas, string.format("%s: %d", produto, quantidade))
  end
  table.sort(linhas) -- só para a saída ficar estável de rodada em rodada
  print(table.concat(linhas, " | ")) -- cafe: 5 | cha: 1

  -- A closure guarda quantos relatórios já foram gerados, sem precisar de
  -- uma variável global nem de um parâmetro extra em toda chamada.
  local function fabricaContadorDeRelatorios()
    local n = 0
    return function()
      n = n + 1
      return n
    end
  end
  local proximoRelatorio = fabricaContadorDeRelatorios()
  print("relatório #" .. proximoRelatorio()) -- relatório #1
  print("relatório #" .. proximoRelatorio()) -- relatório #2

  -- pcall protege um acesso que pode dar error — aqui, chamar algo que não é
  -- função.
  local naoEhFuncao = 42
  local ok, erro = pcall(function() return naoEhFuncao() end)
  print(ok, erro ~= nil) -- false true
end

-- resumoDeEstoque é exatamente o padrão de acumulação usado em exemplos():
-- percorrer a lista com ipairs e escrever no dicionário com "or 0" para
-- cobrir a primeira ocorrência de cada produto.
function M.resumoDeEstoque(itens)
  local resultado = {}
  for _, item in ipairs(itens) do
    resultado[item.produto] = (resultado[item.produto] or 0) + item.quantidade
  end
  return resultado
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
