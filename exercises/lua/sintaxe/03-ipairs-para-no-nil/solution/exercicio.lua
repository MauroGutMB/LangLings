local M = {}

-- somarConhecendoTamanho soma as tamanho primeiras posições de t, tratando
-- um buraco (nil) como 0.
--
-- Como tamanho já é conhecido, um for numérico visita toda posição de 1 a
-- tamanho, buraco ou não — ao contrário de ipairs, que pararia no primeiro
-- nil.
function M.somarConhecendoTamanho(t, tamanho)
  local total = 0
  for i = 1, tamanho do
    total = total + (t[i] or 0)
  end
  return total
end

return M
