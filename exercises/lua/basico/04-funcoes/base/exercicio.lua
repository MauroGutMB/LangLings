-- Função em Lua é um valor como qualquer outro: dá para guardar numa variável,
-- passar como argumento e devolver de outra função.
local M = {}

function M.exemplos()
  -- As duas linhas abaixo declaram a mesma coisa; a de cima é açúcar para a
  -- de baixo.
  local function dobro(n) return n * 2 end
  local triplo = function(n) return n * 3 end
  print(dobro(4), triplo(4)) -- 8 12

  -- Lua não confere a quantidade de argumentos. O que sobra é descartado; o
  -- que falta chega como nil. Daí o idioma do valor padrão com `or`.
  local function saudar(nome, tratamento)
    tratamento = tratamento or "olá"
    return tratamento .. ", " .. nome
  end
  print(saudar("Ana"))         -- olá, Ana
  print(saudar("Ana", "bom dia")) -- bom dia, Ana

  -- Vários retornos são nativos: basta separar por vírgula.
  local function dividir(a, b)
    return a // b, a % b
  end
  local quociente, resto = dividir(17, 5)
  print(quociente, resto) -- 3 2

  -- Recebendo menos variáveis do que a função devolve, o excedente se perde.
  local so_o_primeiro = dividir(17, 5)
  print(so_o_primeiro) -- 3

  -- ... recolhe todos os argumentos a mais. {...} os empacota numa lista.
  local function primeiro(...)
    local args = {...}
    return args[1]
  end
  print(primeiro("a", "b", "c")) -- a

  -- select("#", ...) diz QUANTOS argumentos vieram, e é a única forma que
  -- continua certa quando algum deles é nil — {...} perderia a conta.
  local function quantos(...)
    return select("#", ...)
  end
  print(quantos(1, 2, 3), quantos()) -- 3 0

  -- Função como argumento: aplicar recebe o que fazer, não o que é feito.
  local function aplicar(f, valor) return f(valor) end
  print(aplicar(dobro, 10)) -- 20
end

-- SUA VEZ
--
-- Devolva dois valores: a soma dos números recebidos e quantos eram.
function M.somaEContagem(...)
  return -1, -1 -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
