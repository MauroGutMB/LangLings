-- Toda variável declarada com local existe só dentro do bloco onde nasce.
-- Uma função criada dentro desse bloco, porém, continua enxergando essa
-- variável mesmo depois que o bloco terminou — isso é uma closure.
local M = {}

function M.exemplos()
  do
    local segredo = 42
    print(segredo) -- 42
  end
  -- segredo não existe mais aqui fora; só existia dentro do do..end.

  -- A função interna não copia o valor de x no momento em que foi criada:
  -- ela guarda uma referência à própria variável. Mudar x depois muda o que
  -- a função enxerga.
  local function fabrica()
    local x = 10
    local function ler() return x end
    x = 99
    return ler
  end
  local ler = fabrica()
  print(ler()) -- 99, não 10

  -- Duas chamadas à mesma fábrica criam duas variáveis x independentes —
  -- cada função interna fecha sobre a sua própria, não uma compartilhada.
  local function fabricaSoma()
    local total = 0
    return function(n)
      total = total + n
      return total
    end
  end
  local somaA = fabricaSoma()
  local somaB = fabricaSoma()
  print(somaA(5), somaA(5), somaB(100)) -- 5 10 100
  -- somaB começou do zero dela mesma, sem ver os 10 acumulados em somaA.
end

-- criarContador guarda o total numa variável local que só a função devolvida
-- enxerga — cada chamada a criarContador cria uma variável nova, então os
-- contadores não se pisam.
function M.criarContador()
  local n = 0
  return function()
    n = n + 1
    return n
  end
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
