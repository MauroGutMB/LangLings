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

local resultado, erro = ex.dividirSeguro(10, 2)
verificar("dividirSeguro(10, 2) -> resultado", 5.0, resultado)
verificar("dividirSeguro(10, 2) -> erro", nil, erro)

local resultado2, erro2 = ex.dividirSeguro(7, 0)
verificar("dividirSeguro(7, 0) -> resultado", nil, resultado2)
verificar("dividirSeguro(7, 0) -> erro", "divisão por zero", erro2)

local resultado3, erro3 = ex.dividirSeguro(0, 5)
verificar("dividirSeguro(0, 5) -> resultado", 0.0, resultado3)
verificar("dividirSeguro(0, 5) -> erro", nil, erro3)

local ok = pcall(ex.dividirSeguro, 1, 0)
verificar("dividirSeguro(1, 0) não lança error", true, ok)

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
