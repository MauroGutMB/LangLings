local M = {}

-- clonar devolve uma tabela nova e independente, com os mesmos elementos
-- de t.
--
-- "return t" devolveria a mesma tabela, só com outro nome — uma cópia de
-- verdade precisa de uma tabela nova.
function M.clonar(t)
  local nova = {}
  for i, v in ipairs(t) do
    nova[i] = v
  end
  return nova
end

return M
