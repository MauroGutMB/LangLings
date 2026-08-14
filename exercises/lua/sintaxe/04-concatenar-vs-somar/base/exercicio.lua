local M = {}

-- chaveDoItem monta uma chave como "item-42" a partir de um prefixo e um id.
--
-- TODO: chamar isto quebra na hora. Por quê?
function M.chaveDoItem(prefixo, id)
  return prefixo + "-" + id
end

return M
