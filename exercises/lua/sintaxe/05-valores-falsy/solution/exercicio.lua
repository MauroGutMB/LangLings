local M = {}

-- descricaoDeEstoque devolve "sem estoque" quando quantidade é 0, e
-- "em estoque" quando é maior que 0.
--
-- 0 é verdadeiro em Lua, então a checagem precisa comparar com 0 de forma
-- explícita em vez de usar "not quantidade".
function M.descricaoDeEstoque(quantidade)
  if quantidade == 0 then
    return "sem estoque"
  end
  return "em estoque"
end

return M
