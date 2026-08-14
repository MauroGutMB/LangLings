local M = {}

-- todosPreenchidos diz se t tem as tamanho primeiras posições ocupadas, sem
-- nenhuma delas valendo nil.
--
-- TODO: uma lista com buraco no meio passa nesta checagem. Por quê?
function M.todosPreenchidos(t, tamanho)
  return #t == tamanho
end

return M
