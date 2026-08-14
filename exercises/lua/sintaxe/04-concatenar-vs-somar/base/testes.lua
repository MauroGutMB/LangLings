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

-- Esta é a asserção que separa a versão ingênua da correta: + com um texto
-- não numérico de um dos lados não tem como virar soma, e quebra o programa
-- antes mesmo de comparar um resultado.
verificar("chaveDoItem('item', 42)", "item-42", ex.chaveDoItem("item", 42))
verificar("chaveDoItem('produto', 7)", "produto-7", ex.chaveDoItem("produto", 7))
verificar("chaveDoItem('x', 0)", "x-0", ex.chaveDoItem("x", 0))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
