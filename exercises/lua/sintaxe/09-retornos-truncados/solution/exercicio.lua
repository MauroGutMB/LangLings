local M = {}

-- dividir devolve o quociente e o resto da divisão inteira de a por b.
local function dividir(a, b)
  return a // b, a % b
end

-- empacotar devolve {quociente, resto, "extra"}, com quociente e resto
-- vindos de dividir(a, b).
--
-- Guardar os dois retornos em variáveis antes evita a posição errada no
-- meio do construtor, onde um deles seria descartado.
function M.empacotar(a, b)
  local quociente, resto = dividir(a, b)
  return {quociente, resto, "extra"}
end

return M
