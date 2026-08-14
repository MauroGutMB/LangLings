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

verificar("ehInteiro(3)", true, ex.ehInteiro(3))
verificar("ehInteiro(-7)", true, ex.ehInteiro(-7))
verificar("ehInteiro(0)", true, ex.ehInteiro(0))
verificar("ehInteiro(3.0)", false, ex.ehInteiro(3.0))
verificar("ehInteiro(3.5)", false, ex.ehInteiro(3.5))
verificar('ehInteiro("3")', false, ex.ehInteiro("3"))
verificar("ehInteiro(nil)", false, ex.ehInteiro(nil))
verificar("ehInteiro(true)", false, ex.ehInteiro(true))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
