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

verificar("sem negativos", {1, 2, 3}, ex.removerNegativos({1, 2, 3}))
verificar("um negativo isolado", {1, 3}, ex.removerNegativos({1, -2, 3}))
-- Esta é a asserção que separa a versão ingênua da correta: com -2 e -3
-- seguidos, remover -2 empurra -3 para a posição que ipairs já ia pular,
-- deixando -3 escapar da checagem.
verificar("dois negativos seguidos", {1, 4}, ex.removerNegativos({1, -2, -3, 4}))
verificar("lista vazia", {}, ex.removerNegativos({}))

local original = {1, -2, 3}
ex.removerNegativos(original)
verificar("a lista recebida continua intacta", {1, -2, 3}, original)

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
