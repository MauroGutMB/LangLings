local ex = require("exercicio")

local falhas = 0

local function verificar(o_que, esperado, obtido)
  if esperado == obtido then
    print("ok    " .. o_que)
    return
  end
  print(("FALHA %s\n      esperado: %q\n      obtido:   %q")
    :format(o_que, tostring(esperado), tostring(obtido)))
  falhas = falhas + 1
end

verificar('etiqueta("cafe", 12)', "CAFE custa R$ 12", ex.etiqueta("cafe", 12))
verificar('etiqueta("Pao", 3)', "PAO custa R$ 3", ex.etiqueta("Pao", 3))
verificar('etiqueta("LEITE", 0)', "LEITE custa R$ 0", ex.etiqueta("LEITE", 0))
verificar('etiqueta("", 5)', " custa R$ 5", ex.etiqueta("", 5))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
