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

verificar("contarChaves({})", 0, ex.contarChaves({}))
verificar("contarChaves({a=1, b=2, c=3})", 3, ex.contarChaves({a = 1, b = 2, c = 3}))
verificar("contarChaves(lista pura)", 4, ex.contarChaves({10, 20, 30, 40}))
verificar("contarChaves(mista)", 3, ex.contarChaves({1, 2, extra = true}))
verificar("contarChaves(um só par)", 1, ex.contarChaves({x = 1}))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
