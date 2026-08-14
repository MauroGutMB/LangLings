local M = {}

-- somarConhecendoTamanho soma as tamanho primeiras posições de t, tratando
-- um buraco (nil) como 0.
--
-- TODO: com um buraco no meio, a soma vem menor do que devia. Por quê?
function M.somarConhecendoTamanho(t, tamanho)
  local total = 0
  for _, v in ipairs(t) do
    total = total + v
  end
  return total
end

return M
