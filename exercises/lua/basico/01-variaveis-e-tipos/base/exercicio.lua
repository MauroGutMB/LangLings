-- Em Lua a variável não tem tipo; o VALOR tem. `local` é como se declara
-- qualquer coisa, e limita o nome ao bloco em que ele aparece.
local M = {}

function M.exemplos()
  local nome = "Ana"    -- string
  local idade = 30      -- number, subtipo integer
  local altura = 1.72   -- number, subtipo float
  local ativo = true    -- boolean
  local apelido         -- declarada sem valor: vale nil
  print(nome, idade, altura, ativo, apelido) -- Ana 30 1.72 true nil

  -- type() devolve o tipo do valor, sempre como texto.
  print(type(nome), type(idade), type(ativo), type(apelido))
  -- string  number  boolean  nil

  -- number é um tipo só, mas por dentro tem dois subtipos. math.type separa os
  -- dois — e devolve nil para o que não é número.
  print(math.type(idade), math.type(altura), math.type(nome))
  -- integer  float  nil

  -- Divisão com / produz SEMPRE float, mesmo quando o resultado é exato. Quem
  -- quer o inteiro usa //, a divisão inteira.
  print(10 / 2, 10 // 2)   -- 5.0  5
  print(math.type(10 / 2)) -- float

  -- Conversão é explícita nos dois sentidos, e tonumber avisa quando não deu.
  print(tonumber("42") + 1)  -- 43
  print(tostring(42) .. "!") -- 42!
  print(tonumber("abc"))     -- nil

  -- Reatribuir troca o valor e pode trocar o tipo junto: a variável é só um
  -- nome, não um contrato.
  local x = 1
  x = "um"
  print(type(x)) -- string
end

-- SUA VEZ
--
-- Devolva true quando n for um número inteiro, e false em qualquer outro caso:
-- 3.0 é float, "3" é texto, e nil e true não são números.
function M.ehInteiro(n)
  return false -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
