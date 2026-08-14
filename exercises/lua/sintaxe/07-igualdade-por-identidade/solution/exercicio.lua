local M = {}

-- mesmoConteudo diz se as listas a e b têm os mesmos elementos, na mesma
-- ordem.
--
-- == entre tabelas é por identidade, não por conteúdo — comparar conteúdo
-- exige olhar posição por posição.
function M.mesmoConteudo(a, b)
  if #a ~= #b then
    return false
  end
  for i = 1, #a do
    if a[i] ~= b[i] then
      return false
    end
  end
  return true
end

return M
