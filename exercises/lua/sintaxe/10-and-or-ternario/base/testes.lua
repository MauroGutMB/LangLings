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

verificar("condicao true, valores comuns", "sim", ex.escolher(true, "sim", "nao"))
verificar("condicao false, valores comuns", "nao", ex.escolher(false, "sim", "nao"))
-- Esta é a asserção que separa a versão ingênua da correta: condicao é true,
-- mas valorSeVerdadeiro é false — o idioma and/or engole esse false e cai
-- para valorSeFalso, embora a condição tenha sido atendida.
verificar("condicao true, valorSeVerdadeiro é false", false, ex.escolher(true, false, "nao"))
verificar("condicao false, valorSeVerdadeiro é false", "nao", ex.escolher(false, false, "nao"))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
