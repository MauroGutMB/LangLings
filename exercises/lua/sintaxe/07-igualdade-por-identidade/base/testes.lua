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

local mesma = {1, 2, 3}
verificar("a mesma tabela é igual a si mesma", true, ex.mesmoConteudo(mesma, mesma))
-- Esta é a asserção que separa a versão ingênua da correta: {1,2,3} e {1,2,3}
-- são duas tabelas distintas na memória, mesmo tendo o mesmo conteúdo — ==
-- as considera diferentes.
verificar("conteúdo igual, tabelas diferentes", true, ex.mesmoConteudo({1, 2, 3}, {1, 2, 3}))
verificar("conteúdo diferente", false, ex.mesmoConteudo({1, 2, 3}, {1, 2, 4}))
verificar("tamanhos diferentes", false, ex.mesmoConteudo({1, 2}, {1, 2, 3}))
verificar("duas listas vazias", true, ex.mesmoConteudo({}, {}))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
