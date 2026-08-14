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

verificar("clonar tem o mesmo conteúdo", {1, 2, 3}, ex.clonar({1, 2, 3}))
verificar("clonar lista vazia", {}, ex.clonar({}))

-- Esta é a asserção que separa a versão ingênua da correta: mexer na cópia
-- não pode vazar de volta pro original — "return t" devolve a mesma
-- tabela, então mexer numa mexe nas duas.
local original = {1, 2, 3}
local copia = ex.clonar(original)
table.insert(copia, 4)
if iguais({1, 2, 3}, original) then
  print("ok    original não é afetado ao alterar a cópia")
else
  print(("FALHA original não é afetado ao alterar a cópia\n      esperado: %s\n      obtido:   %s")
    :format(mostrar({1, 2, 3}), mostrar(original)))
  falhas = falhas + 1
end

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
