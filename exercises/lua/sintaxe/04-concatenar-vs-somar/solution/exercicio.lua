local M = {}

-- chaveDoItem monta uma chave como "item-42" a partir de um prefixo e um id.
--
-- .. concatena texto e converte o número sozinho; + é para aritmética.
function M.chaveDoItem(prefixo, id)
  return prefixo .. "-" .. id
end

return M
