-- Tabela é a única estrutura de dados de Lua. Usada com índices numéricos
-- seguidos a partir de 1, ela é a lista da linguagem.
local M = {}

function M.exemplos()
  local linguagens = {"lua", "go", "rust"}

  -- A contagem começa em 1, não em 0. Ler fora do intervalo devolve nil, sem
  -- erro nenhum.
  print(linguagens[1], linguagens[3], linguagens[4]) -- lua rust nil

  -- # devolve o comprimento da lista.
  print(#linguagens, #{}) -- 3 0

  -- table.insert acrescenta no fim; com uma posição no meio, empurra o resto
  -- para a direita.
  table.insert(linguagens, "zig")
  table.insert(linguagens, 1, "c")
  print(#linguagens, linguagens[1], linguagens[5]) -- 5 c zig

  -- table.remove tira da posição indicada (o fim, se não disser qual) e
  -- devolve o que tirou, fechando o buraco.
  local saiu = table.remove(linguagens, 1)
  print(saiu, #linguagens) -- c 4

  -- ipairs percorre as posições 1, 2, 3... em ordem, entregando índice e
  -- valor.
  for i, nome in ipairs(linguagens) do
    io.write(i, "=", nome, " ") -- 1=lua 2=go 3=rust 4=zig
  end
  print()

  -- Escrever direto em nova[i] é equivalente a inserir, quando você já sabe a
  -- posição. É o padrão de "criar uma lista a partir de outra".
  local maiusculas = {}
  for i, nome in ipairs(linguagens) do
    maiusculas[i] = nome:upper()
  end
  print(maiusculas[1], #maiusculas) -- LUA 4

  -- table.concat junta os elementos num texto só, com um separador.
  print(table.concat(maiusculas, ", ")) -- LUA, GO, RUST, ZIG
end

-- SUA VEZ
--
-- Devolva uma lista nova com cada número de t multiplicado por 2.
function M.dobrar(t)
  return {} -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
