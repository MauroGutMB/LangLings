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

verificar("lista cheia de 3", true, ex.todosPreenchidos({1, 2, 3}, 3))
verificar("lista vazia, tamanho 0", true, ex.todosPreenchidos({}, 0))
-- Esta é a asserção que separa a versão ingênua da correta: {1, 2, nil, 4, 5}
-- tem um buraco na posição 3, mas o interpretador ainda reporta #t == 5.
verificar("lista de 5 com buraco no meio", false, ex.todosPreenchidos({1, 2, nil, 4, 5}, 5))
verificar("lista de 3 com buraco no fim", false, ex.todosPreenchidos({1, 2, nil}, 3))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
