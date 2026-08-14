local M = {}

-- removerNegativos devolve uma lista nova só com os elementos de t que não
-- são negativos, na mesma ordem.
--
-- TODO: dois negativos seguidos e um deles sobrevive. Por quê?
function M.removerNegativos(t)
  for i, v in ipairs(t) do
    if v < 0 then
      table.remove(t, i)
    end
  end
  return t
end

return M
