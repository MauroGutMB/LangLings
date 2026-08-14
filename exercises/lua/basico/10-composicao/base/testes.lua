local ex = require("exercicio")

local falhas = 0

local function iguaisDicionario(a, b)
  for k, v in pairs(a) do
    if b[k] ~= v then return false end
  end
  for k, v in pairs(b) do
    if a[k] ~= v then return false end
  end
  return true
end

local function mostrar(t)
  local partes = {}
  for k, v in pairs(t) do
    table.insert(partes, tostring(k) .. "=" .. tostring(v))
  end
  table.sort(partes)
  return "{" .. table.concat(partes, ", ") .. "}"
end

local function verificar(o_que, esperado, obtido)
  if iguaisDicionario(esperado, obtido) then
    print("ok    " .. o_que)
    return
  end
  print(("FALHA %s\n      esperado: %s\n      obtido:   %s")
    :format(o_que, mostrar(esperado), mostrar(obtido)))
  falhas = falhas + 1
end

verificar("lista vazia", {}, ex.resumoDeEstoque({}))

verificar("um produto, uma entrada", {maca = 5},
  ex.resumoDeEstoque({{produto = "maca", quantidade = 5}}))

verificar("mesmo produto repetido soma", {cafe = 5, cha = 1},
  ex.resumoDeEstoque({
    {produto = "cafe", quantidade = 2},
    {produto = "cha", quantidade = 1},
    {produto = "cafe", quantidade = 3},
  }))

verificar("produtos distintos não se misturam", {a = 1, b = 2, c = 3},
  ex.resumoDeEstoque({
    {produto = "a", quantidade = 1},
    {produto = "b", quantidade = 2},
    {produto = "c", quantidade = 3},
  }))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
