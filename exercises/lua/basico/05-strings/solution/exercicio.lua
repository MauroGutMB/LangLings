-- Toda string em Lua é imutável: as funções de texto nunca alteram o original,
-- sempre devolvem um novo.
local M = {}

function M.exemplos()
  local nome = "Ana"

  -- Concatenação é .., não +. O + tentaria converter os dois lados em número.
  print("olá, " .. nome) -- olá, Ana

  -- # mede o comprimento em bytes.
  print(#nome, #"") -- 3 0

  -- string.format monta o texto de uma vez, com marcadores no lugar dos
  -- valores: %s recebe texto, %d um inteiro, %.2f um decimal arredondado.
  print(string.format("%s tem %d anos", nome, 30))   -- Ana tem 30 anos
  print(string.format("total: R$ %.2f", 12.5))       -- total: R$ 12.50

  -- Toda função de string também é um método do próprio texto: a chamada com
  -- dois pontos passa o texto como primeiro argumento automaticamente.
  print(nome:upper(), string.upper(nome)) -- ANA ANA
  print(nome:lower())                     -- ana
  print(nome:len())                       -- 3

  -- sub recorta por posição, contando a partir de 1. Índice negativo conta do
  -- fim para o começo, o que dá "os n últimos" de graça.
  local frase = "linguagem lua"
  print(frase:sub(1, 9))  -- linguagem
  print(frase:sub(-3))    -- lua

  -- rep repete, find procura e devolve onde começou (ou nil), gsub troca.
  print(("ab"):rep(3))            -- ababab
  print(frase:find("lua"))        -- 11 13
  print(frase:gsub("lua", "Lua")) -- linguagem Lua  1

  -- O original continua o mesmo depois de tudo isso.
  print(frase) -- linguagem lua
end

-- etiqueta monta o texto com um molde único.
--
-- Com .. seriam quatro pedaços e três operadores, e a forma final do texto só
-- apareceria depois de ler todos eles; com format o molde é legível de uma vez
-- e os valores ficam separados dele.
function M.etiqueta(nome, preco)
  return string.format("%s custa R$ %d", nome:upper(), preco)
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
