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

verificar("somaPares(10)", 30, ex.somaPares(10))
verificar("somaPares(9)", 20, ex.somaPares(9))
verificar("somaPares(2)", 2, ex.somaPares(2))
verificar("somaPares(1)", 0, ex.somaPares(1))
verificar("somaPares(0)", 0, ex.somaPares(0))
verificar("somaPares(-4)", 0, ex.somaPares(-4))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
