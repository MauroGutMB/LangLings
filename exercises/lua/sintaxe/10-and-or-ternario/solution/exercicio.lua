local M = {}

-- escolher devolve valorSeVerdadeiro quando condicao é true, e valorSeFalso
-- quando é false.
--
-- Um if comum não tem a armadilha do idioma and/or quando o valor
-- "verdadeiro" é, ele mesmo, false.
function M.escolher(condicao, valorSeVerdadeiro, valorSeFalso)
  if condicao then
    return valorSeVerdadeiro
  end
  return valorSeFalso
end

return M
