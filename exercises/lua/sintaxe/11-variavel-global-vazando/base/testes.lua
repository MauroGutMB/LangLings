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

verificar("primeira chamada", 6, ex.soma({1, 2, 3}))
-- Esta é a asserção que separa a versão ingênua da correta: sem local, total
-- sobrevive entre chamadas, então a segunda soma parte do resultado da
-- primeira em vez de partir de 0.
verificar("segunda chamada, lista diferente", 10, ex.soma({10}))
verificar("terceira chamada, lista vazia", 0, ex.soma({}))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
