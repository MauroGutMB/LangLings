-- Lua não tem exceções no sentido de outras linguagens. error() interrompe a
-- função na hora, e pcall() é quem decide se isso derruba o programa ou vira
-- só um retorno de erro.
local M = {}

function M.exemplos()
  local function raiz(n)
    if n < 0 then
      error("não existe raiz de número negativo")
    end
    return n ^ 0.5
  end

  -- pcall chama a função protegida. Se ela não der error, o primeiro retorno
  -- é true e o resto é o que a função devolveu.
  local ok, resultado = pcall(raiz, 16)
  print(ok, resultado) -- true 4.0

  -- Se a função der error, pcall NÃO deixa o erro subir: devolve false e a
  -- mensagem no lugar do resultado. O programa continua rodando.
  local ok2, mensagem = pcall(raiz, -4)
  print(ok2, mensagem) -- false, seguido de algo terminando em "...negativo"

  -- Chamar raiz(-4) direto, sem pcall, teria interrompido o programa —
  -- é por isso que pcall existe: para decidir "isso pode falhar, e eu quero
  -- lidar com a falha aqui, não morrer".

  -- Nem toda função sinaliza erro com error(). Muitas, por convenção,
  -- devolvem nil mais uma mensagem — sem precisar de pcall nenhum para
  -- serem chamadas.
  local function buscar(tabela, chave)
    local valor = tabela[chave]
    if valor == nil then
      return nil, "chave não encontrada: " .. tostring(chave)
    end
    return valor, nil
  end
  local v, erro = buscar({x = 10}, "y")
  print(v, erro) -- nil chave não encontrada: y
end

-- SUA VEZ
--
-- Devolva a / b e nil quando b não for 0; devolva nil e a mensagem
-- "divisão por zero" quando b for 0. Nunca chame error.
function M.dividirSeguro(a, b)
  return -1, nil -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
