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

verificar("quantidade 5", "em estoque", ex.descricaoDeEstoque(5))
verificar("quantidade 1", "em estoque", ex.descricaoDeEstoque(1))
-- Esta é a asserção que separa a versão ingênua da correta: 0 é um valor
-- verdadeiro em Lua, então "not 0" é false e a versão ingênua nunca entra no
-- ramo de "sem estoque" para ele.
verificar("quantidade 0", "sem estoque", ex.descricaoDeEstoque(0))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
