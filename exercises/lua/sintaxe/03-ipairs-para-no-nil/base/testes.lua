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

verificar("sem buraco", 6, ex.somarConhecendoTamanho({1, 2, 3}, 3))
verificar("lista vazia", 0, ex.somarConhecendoTamanho({}, 0))
-- Esta é a asserção que separa a versão ingênua da correta: ipairs para na
-- posição 2 (a 3 é nil) e nunca visita as posições 4 e 5, que somariam 9.
verificar("buraco no meio", 12, ex.somarConhecendoTamanho({1, 2, nil, 4, 5}, 5))
verificar("buraco logo na primeira posição", 5, ex.somarConhecendoTamanho({nil, 5}, 2))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
