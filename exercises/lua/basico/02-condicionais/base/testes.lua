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

verificar("classificar(-3)", "negativo", ex.classificar(-3))
verificar("classificar(-0.5)", "negativo", ex.classificar(-0.5))
verificar("classificar(0)", "zero", ex.classificar(0))
verificar("classificar(7)", "positivo", ex.classificar(7))
verificar("classificar(0.1)", "positivo", ex.classificar(0.1))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
