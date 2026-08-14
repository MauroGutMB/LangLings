local ex = require("exercicio")

local falhas = 0

local function verificar(o_que, esperado, obtido)
  if esperado == obtido then
    print("ok    " .. o_que)
    return
  end
  print(("FALHA %s\n      esperado: %s\n      obtido:   %s")
    :format(o_que, tostring(esperado), tostring(obtido)))
  falhas = falhas + 1
end

-- Os dois retornos são checados separadamente para que a mensagem diga qual
-- dos dois está errado.
local function conferirPar(o_que, somaEsperada, qtdEsperada, ...)
  local soma, qtd = ex.somaEContagem(...)
  verificar(o_que .. " -> soma", somaEsperada, soma)
  verificar(o_que .. " -> quantidade", qtdEsperada, qtd)
end

conferirPar("somaEContagem(1, 2, 3)", 6, 3, 1, 2, 3)
conferirPar("somaEContagem(10)", 10, 1, 10)
conferirPar("somaEContagem()", 0, 0)
conferirPar("somaEContagem(-1, 1)", 0, 2, -1, 1)
conferirPar("somaEContagem(0.5, 0.25)", 0.75, 2, 0.5, 0.25)

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
