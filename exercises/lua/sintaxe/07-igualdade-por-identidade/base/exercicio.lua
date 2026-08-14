local M = {}

-- mesmoConteudo diz se as listas a e b têm os mesmos elementos, na mesma
-- ordem.
--
-- TODO: duas listas com o mesmo conteúdo, mas criadas em separado, são
-- relatadas como diferentes. Por quê?
function M.mesmoConteudo(a, b)
  return a == b
end

return M
