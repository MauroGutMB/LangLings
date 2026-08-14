local M = {}

-- soma devolve a soma dos números de lista.
--
-- local garante que total nasce do zero a cada chamada, em vez de
-- sobreviver como global entre uma chamada e a próxima.
function M.soma(lista)
  local total = 0
  for _, n in ipairs(lista) do
    total = total + n
  end
  return total
end

return M
