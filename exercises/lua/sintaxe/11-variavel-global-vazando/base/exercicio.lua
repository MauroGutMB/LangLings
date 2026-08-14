local M = {}

-- soma devolve a soma dos números de lista.
--
-- TODO: chamar isto duas vezes seguidas dá um resultado errado na segunda
-- vez. Por quê?
function M.soma(lista)
  total = (total or 0)
  for _, n in ipairs(lista) do
    total = total + n
  end
  return total
end

return M
