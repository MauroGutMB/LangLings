local ex = require("exercicio")

local falhas = 0

local function iguais(a, b)
  if #a ~= #b then return false end
  for i = 1, #a do
    if a[i] ~= b[i] then return false end
  end
  return true
end

local function mostrar(t)
  local partes = {}
  for i = 1, #t do partes[i] = tostring(t[i]) end
  return "{" .. table.concat(partes, ", ") .. "}"
end

local function verificar(o_que, esperado, obtido)
  if iguais(esperado, obtido) then
    print("ok    " .. o_que)
    return
  end
  print(("FALHA %s\n      esperado: %s\n      obtido:   %s")
    :format(o_que, mostrar(esperado), mostrar(obtido)))
  falhas = falhas + 1
end

-- Esta é a asserção que separa a versão ingênua da correta: dividir(17, 5)
-- devolve 3 e 2, mas colocada no meio do construtor da lista, só o 3
-- sobrevive — o 2 desaparece e "extra" ocupa a posição errada.
verificar("empacotar(17, 5)", {3, 2, "extra"}, ex.empacotar(17, 5))
verificar("empacotar(10, 3)", {3, 1, "extra"}, ex.empacotar(10, 3))
verificar("empacotar(9, 3)", {3, 0, "extra"}, ex.empacotar(9, 3))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
