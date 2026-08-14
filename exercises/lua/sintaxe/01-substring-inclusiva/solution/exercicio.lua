local M = {}

-- cortarAntesDe devolve tudo o que vem antes da primeira ocorrência de
-- separador em texto, sem incluir o separador. Se separador não aparece,
-- devolve texto inteiro.
--
-- pos é a posição do separador, 1-based e inclusiva; o último caractere
-- que interessa é o de antes dela.
function M.cortarAntesDe(texto, separador)
  local pos = texto:find(separador, 1, true)
  if not pos then
    return texto
  end
  return texto:sub(1, pos - 1)
end

return M
