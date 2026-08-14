local M = {}

-- todosPreenchidos diz se t tem as tamanho primeiras posições ocupadas, sem
-- nenhuma delas valendo nil.
--
-- Só dá pra saber isso olhando posição por posição — # não garante ausência
-- de buracos.
function M.todosPreenchidos(t, tamanho)
  for i = 1, tamanho do
    if t[i] == nil then
      return false
    end
  end
  return true
end

return M
