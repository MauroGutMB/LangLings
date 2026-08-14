local ex = require("exercicio")

local falhas = 0

local function iguais(a, b)
  if type(a) ~= "table" or type(b) ~= "table" then return a == b end
  if #a ~= #b then return false end
  for i = 1, #a do
    if not iguais(a[i], b[i]) then return false end
  end
  return true
end

local function mostrar(v)
  if type(v) ~= "table" then return tostring(v) end
  local partes = {}
  for i = 1, #v do partes[i] = tostring(v[i]) end
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

verificar("dobrar({1, 2, 3})", {2, 4, 6}, ex.dobrar({1, 2, 3}))
verificar("dobrar({-2, 0, 0.5})", {-4, 0, 1.0}, ex.dobrar({-2, 0, 0.5}))
verificar("dobrar({7})", {14}, ex.dobrar({7}))
verificar("dobrar({})", {}, ex.dobrar({}))

local original = {1, 2, 3}
ex.dobrar(original)
verificar("a lista recebida continua intacta", {1, 2, 3}, original)

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
