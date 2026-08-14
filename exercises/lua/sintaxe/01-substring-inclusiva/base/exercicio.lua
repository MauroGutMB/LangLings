local M = {}

-- cortarAntesDe devolve tudo o que vem antes da primeira ocorrência de
-- separador em texto, sem incluir o separador. Se separador não aparece,
-- devolve texto inteiro.
--
-- TODO: o texto devolvido sempre inclui o separador no fim. Por quê?
function M.cortarAntesDe(texto, separador)
  local pos = texto:find(separador, 1, true)
  if not pos then
    return texto
  end
  return texto:sub(1, pos)
end

return M
