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

local c1 = ex.criarContador()
verificar("primeiro contador, 1a chamada", 1, c1())
verificar("primeiro contador, 2a chamada", 2, c1())
verificar("primeiro contador, 3a chamada", 3, c1())

local c2 = ex.criarContador()
verificar("segundo contador começa do zero", 1, c2())
verificar("primeiro contador não foi afetado pelo segundo", 4, c1())
verificar("segundo contador, 2a chamada", 2, c2())

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
