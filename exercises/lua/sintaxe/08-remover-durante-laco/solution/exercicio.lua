local M = {}

-- removerNegativos devolve uma lista nova só com os elementos de t que não
-- são negativos, na mesma ordem.
--
-- Construir uma lista nova evita mexer na lista que ipairs ainda está
-- percorrendo.
function M.removerNegativos(t)
  local resultado = {}
  for _, v in ipairs(t) do
    if v >= 0 then
      table.insert(resultado, v)
    end
  end
  return resultado
end

return M
