local M = {}

-- descricaoDeEstoque devolve "sem estoque" quando quantidade é 0, e
-- "em estoque" quando é maior que 0.
--
-- TODO: um estoque de 0 unidades é relatado como "em estoque". Por quê?
function M.descricaoDeEstoque(quantidade)
  if not quantidade then
    return "sem estoque"
  end
  return "em estoque"
end

return M
