local M = {}

-- escolher devolve valorSeVerdadeiro quando condicao é true, e valorSeFalso
-- quando é false.
--
-- TODO: existe um caso em que condicao é true e mesmo assim o devolvido é
-- valorSeFalso. Por quê?
function M.escolher(condicao, valorSeVerdadeiro, valorSeFalso)
  return condicao and valorSeVerdadeiro or valorSeFalso
end

return M
