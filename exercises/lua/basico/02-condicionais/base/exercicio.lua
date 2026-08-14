-- Um if em Lua vai de `then` até `end`. Não há parênteses obrigatórios em
-- volta da condição, e não há chaves.
local M = {}

function M.exemplos()
  local idade = 20

  if idade >= 18 then
    print("maior de idade") -- maior de idade
  else
    print("menor de idade")
  end

  -- elseif é uma palavra só. Escrever "else if" abre um segundo if aninhado e
  -- cobra um `end` a mais no fim.
  local nota = 7
  if nota >= 9 then
    print("ótimo")
  elseif nota >= 6 then
    print("aprovado") -- aprovado
  else
    print("reprovado")
  end

  -- Comparações. O "diferente" é ~=, não !=.
  print(1 < 2, 2 <= 2, 3 > 4, "a" == "a", "a" ~= "b")
  -- true  true  false  true  true

  -- and, or e not no lugar de &&, || e !.
  local chovendo, guardaChuva = true, false
  print(chovendo and not guardaChuva) -- true
  print(chovendo or guardaChuva)      -- true

  -- A regra de verdade de Lua é curta e vale a pena decorar: só nil e false
  -- são falsos. Todo o resto é verdadeiro — inclusive o zero e o texto vazio,
  -- que em outras linguagens seriam falsos.
  if 0 then print("zero é verdadeiro") end   -- zero é verdadeiro
  if "" then print("vazio é verdadeiro") end -- vazio é verdadeiro
  if nil then print("nunca imprime") end

  -- and e or devolvem um dos operandos, não um boolean. É o que torna
  -- `x or padrao` o jeito usual de dar valor padrão a um argumento ausente.
  local apelido = nil
  print(apelido or "sem apelido") -- sem apelido
end

-- SUA VEZ
--
-- Devolva "negativo", "zero" ou "positivo" conforme o valor de n.
function M.classificar(n)
  return "" -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
